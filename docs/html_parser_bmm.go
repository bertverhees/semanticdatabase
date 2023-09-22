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
	data := parse(text)
	fmt.Println(data)
}

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func parse(text string) (data string) {
	tkn := html.NewTokenizer(strings.NewReader(text))

	var h4 bool

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return ""

		case tt == html.StartTagToken:

			t := tkn.Token()
			h4 = t.Data == "h4"
			if h4 {
				fmt.Println("->" + t.Data)
				tt = tkn.Next()
				t := tkn.Token()
				a := t.Data == "a"
				if a {
					fmt.Println("---->" + t.Data)
					tt = tkn.Next()
					tt = tkn.Next()
					t3 := tkn.Token()
					span := t3.Data == "span"
					if span {
						fmt.Println("---------->" + t3.Data)
						tt = tkn.Next()
						t4 := tkn.Token()
						if tt == html.TextToken {
							fmt.Println("---------->" + t4.Data)
						}
					}
				}
				h4 = false
			}
		}
	}
}
