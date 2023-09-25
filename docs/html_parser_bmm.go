package main

//https://zetcode.com/golang/net-html/
import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fileName := "bmm.html"
	text, err := readHtmlFromFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	data := parseBMM(text)
	fmt.Println(data)

}

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func parseBMM(text string) (data string) {
	tkn := html.NewTokenizer(strings.NewReader(text))

	h2 := 0
	h3 := 0
	h4 := 0
	packageString := ""
	classString := ""
	isTD1 := false
	isTD2 := false
	isTD3 := false
	isTR := false
	firstClassPassed := false

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return ""

		case tt == html.EndTagToken:
			t := tkn.Token()
			//TR
			if t.Data == "tr" && isTR && firstClassPassed {
				isTD1 = false
				isTD2 = false
				isTD3 = false
				isTR = false
			}

		case tt == html.TextToken:
			if isTR {
				//GET Text
				for tt != html.TextToken {
					tt = tkn.Next()
				}
				tTD := tkn.Token()
				fmt.Println("\t\t\t\t\t" + tTD.Data)
			}
		case tt == html.StartTagToken:
			t := tkn.Token()
			// table
			if t.Data == "table" && firstClassPassed {
			}
			//TR
			if t.Data == "tr" && !isTR && firstClassPassed {
				tTR := tkn.Token()
				fmt.Println("\t\t\t" + tTR.Data)
				isTR = true
				isTD1 = true
				isTD2 = false
				isTD3 = false
			}
			//TD
			if t.Data == "td" && firstClassPassed {
				switch {
				//TD1
				case isTR && isTD1 && !isTD2 && !isTD3:
					tTD := tkn.Token()
					fmt.Println("\t\t\t\t td1 " + tTD.Data)
					isTD2 = true
					isTD1 = false
					//TD2
				case isTR && !isTD1 && isTD2 && !isTD3:
					tTD := tkn.Token()
					fmt.Println("\t\t\t\t td2" + tTD.Data)
					isTD3 = true
					isTD2 = false
					//TD3
				case isTR && !isTD1 && !isTD2 && isTD3:
					tTD := tkn.Token()
					fmt.Println("\t\t\t\t td3" + tTD.Data)
					isTD3 = false
				}
			}
			//Package
			if t.Data == "h2" {
				tkn.Next()
				t1 := tkn.Token()
				if t1.Data == "a" {
					tkn.Next()
					tkn.Next()
					t2 := tkn.Token()
					if t2.Data == "span" {
						tt = tkn.Next()
						t3 := tkn.Token()
						if tt == html.TextToken {
							packageString = "org.openehr.lang.bmm"
							switch {
							case strings.HasPrefix(t3.Data, "3."):
								packageString = "org.openehr.lang.bmm"
							case strings.HasPrefix(t3.Data, "4."):
								packageString = "org.openehr.lang.bmm.model_access"
							case strings.HasPrefix(t3.Data, "5."):
								packageString = "base.bmm.model_access"
							case strings.HasPrefix(t3.Data, "6."):
								packageString = "base.bmm.core.entity"
							case strings.HasPrefix(t3.Data, "7."):
								packageString = "base.bmm.core"
							case strings.HasPrefix(t3.Data, "8."):
								packageString = "base.bmm.core.feature"
							case strings.HasPrefix(t3.Data, "9."):
								packageString = "base.bmm.core.literal_value"
							case strings.HasPrefix(t3.Data, "10."):
								packageString = "base.bmm.core.expression"
							case strings.HasPrefix(t3.Data, "11."):
								packageString = ""
							case strings.HasPrefix(t3.Data, "12."):
								packageString = "base.bmm.core.statement"
							case strings.HasPrefix(t3.Data, "13."):
								packageString = ""
							case strings.HasPrefix(t3.Data, "14."):
								packageString = ""
							}
							fmt.Println(packageString)
						}
					}
				}
				h2++
			}
			if t.Data == "h3" {
				h3++
			}
			//Class
			if t.Data == "h4" {
				if h4 > 2 {
					tkn.Next()
					t1 := tkn.Token()
					for t1.Data != "a" {
						tkn.Next()
						t1 = tkn.Token()
					}
					if t1.Data == "a" {
						tkn.Next()
						t2 := tkn.Token()
						for t2.Data != "span" {
							tkn.Next()
							t2 = tkn.Token()
						}
						if t2.Data == "span" {
							tt = tkn.Next()
							for tt != html.TextToken {
								tt = tkn.Next()
							}
							t3 := tkn.Token()
							if tt == html.TextToken {
								classString = t3.Data
								if strings.Contains(classString, "BMM_DEFINITIONS Class") {
									firstClassPassed = true
								}
								if firstClassPassed {
									fmt.Println("\t" + classString)
								}
							}
						}
					}
				}
				h4++
			}
		}
	}
}
