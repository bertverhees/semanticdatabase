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

func formatString(in string) string {
	r := strings.ReplaceAll(in, "\n", " ")
	r = strings.ReplaceAll(r, "\t", "")
	r = strings.ReplaceAll(r, "Â", "")
	r = strings.ReplaceAll(r, "  ", " ")
	r = strings.TrimSpace(r)
	return r
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
	td1 := ""
	td2 := ""
	td3 := ""
	constantName := ""
	constantDescription := ""
	attributeName := ""
	attributeDescription := ""
	functionName := ""
	functionDescription := ""
	classDescription := false
	classInherit := false
	className := false
	constants := false
	attributes := false
	functions := false

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return ""

		case tt == html.EndTagToken:
			t := tkn.Token()
			//TABLE
			if t.Data == "table" && firstClassPassed {
				constants = false
				attributes = false
				functions = false
			}
			//TR
			if t.Data == "tr" && isTR && firstClassPassed {
				switch {
				case constants:
					fmt.Println("constants name: " + constantName)
					fmt.Println("constants description: " + constantDescription)
				case attributes:
					fmt.Println("attribute name: " + attributeName)
					fmt.Println("attribute description: " + attributeDescription)
				case functions:
					fmt.Println("function name: " + functionName)
					fmt.Println("function description: " + functionDescription)
				}
				isTD1 = false
				isTD2 = false
				isTD3 = false
				isTR = false
			}
			if t.Data == "td" && isTR && firstClassPassed {
				// fmt.Println(isTD1, isTD2, isTD3)
				switch {
				case isTD1:
					isTD2 = true
					isTD1 = false
				case isTD2:
					isTD3 = true
					isTD2 = false
				case isTD3:
					isTD3 = false
				}
				// fmt.Println(isTD1, isTD2, isTD3)
			}

		case tt == html.TextToken:
			if isTR {
				//GET Text
				for tt != html.TextToken {
					tt = tkn.Next()
				}
				tTD := tkn.Token()
				t := formatString(tTD.Data)
				if t != "" {
					switch {
					case isTD1:
						td1 = t
					case isTD2:
						td2 = t
					case isTD3:
						td3 = t
					}
					switch {
					case isTD1 && strings.Compare(td1, "Class") == 0:
						className = true
					case isTD1 && strings.Compare(td1, "Description") == 0:
						classDescription = true
					case isTD1 && strings.Compare(td1, "Inherit") == 0:
						classInherit = true
					case isTD1 && strings.Compare(td1, "Constants") == 0:
						constants = true
					case isTD1 && strings.Compare(td1, "Attributes") == 0:
						attributes = true
					case isTD1 && strings.Compare(td1, "Functions") == 0:
						functions = true
					case isTD2 && className:
						fmt.Println("Class: " + td2)
						className = false
					case isTD2 && classDescription:
						fmt.Println("ClassDescription: " + td2)
						classDescription = false
					case isTD2 && classInherit:
						fmt.Println("classInherit: " + td2)
						classInherit = false
					case isTD1 && constants:
						fmt.Println("constants occurence: " + td1)
					case isTD2 && constants:
						constantName = constantName + td2
					case isTD3 && constants:
						constantDescription = constantDescription + td3
					case isTD1 && attributes:
						fmt.Println("attribute occurence: " + td1)
					case isTD2 && attributes:
						attributeName = attributeName + td2
					case isTD3 && attributes:
						attributeDescription = attributeDescription + td3
					case isTD1 && functions:
						fmt.Println("functions occurence: " + td1)
					case isTD2 && functions:
						functionName = functionName + td2
					case isTD3 && functions:
						functionDescription = functionDescription + td3
					}
					switch {
					case isTD1 && (strings.Compare(td1, "Attributes") == 0 || strings.Compare(td1, "Functions") == 0 || strings.Compare(td1, "Class") == 0):
						constants = false
					case isTD1 && (strings.Compare(td1, "Constants") == 0 || strings.Compare(td1, "Functions") == 0 || strings.Compare(td1, "Class") == 0):
						attributes = false
					case isTD1 && (strings.Compare(td1, "Constants") == 0 || strings.Compare(td1, "Attributes") == 0 || strings.Compare(td1, "Class") == 0):
						functions = false

					}
					// fmt.Println("*" + t + "*")
				}
				// fmt.Println("\n")
				// for i := 0; i < len(t); i++ {
				// 	fmt.Printf("%x ", t[i])
				// }
				// fmt.Println("\n")
				// for i := 0; i < len(t); i++ {
				// 	fmt.Printf("%q", t[i])
				// }
			}
		case tt == html.StartTagToken:
			t := tkn.Token()
			// table
			if t.Data == "table" && firstClassPassed {
			}
			//TR
			if t.Data == "tr" && !isTR && firstClassPassed {
				// tTR := tkn.Token()
				// fmt.Println("\t\t\t" + tTR.Data)
				isTR = true
				isTD1 = true
				isTD2 = false
				isTD3 = false
				constantName = ""
				constantDescription = ""
				attributeName = ""
				attributeDescription = ""
				functionName = ""
				functionDescription = ""
				switch {
				case constants:
					fmt.Println("\t\t\t" + "CONSTANT")
				case attributes:
					fmt.Println("\t\t\t" + "ATTRIBUTE")
				case functions:
					fmt.Println("\t\t\t" + "FUNCTION")
				}
			}
			//TD
			if t.Data == "td" && firstClassPassed {
				switch {
				//TD1
				case isTR && isTD1 && !isTD2 && !isTD3:
					// tkn.Token()
					// fmt.Print("td1: ")
					//TD2
				case isTR && !isTD1 && isTD2 && !isTD3:
					// tkn.Token()
					// fmt.Print("td2: ")
					//TD3
				case isTR && !isTD1 && !isTD2 && isTD3:
					// tkn.Token()
					// fmt.Print("td3: ")
				}
			}
			//Package
			if t.Data == "h2" {
				tkn.Next()
				t1 := tkn.Token()
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
								// if firstClassPassed {
								// 	fmt.Println("\t" + classString)
								// 	classDescription = true
								// }
							}
						}
					}
				}
				h4++
			}
		}
	}
}
