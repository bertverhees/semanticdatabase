package main

//https://zetcode.com/golang/net-html/
import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fileName := "bmm.html"
	text, err := readHtmlFromFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	model := NewModel()
	preProcess(text,"BMM_DEFINITIONS Class")
	text, err = readHtmlFromFile("tmp.html")
	parseBMM(text, model)
	fmt.Println(len(model.Classes))
	for _,c := range model.Classes {
		c.Print(model)
	}
	for _,e := range model.Enumerations {
		e.Print()
	}
	fmt.Println("----------------------------------------------------")
	//#TODO rewrite table
	//#TODO Missing inheritance BMM_MODULE at BMM_CLASS
	//#TODO Missing inheritance EL_PREDICATE at EL_DEFINED
	//#TODO Missing inheritance EL_AGENT_CALL at BMM_PROCEDURE_CALL
	//#TODO Missing inheritance EL_FEATURE_REF at EL_FUNCTION_CALL
	//#TODO Check number of
	//#TODO classes
	//#TODO inheritence
	//#TODO constants
	//#TODO attributes
	//#TODO functions
	for _,c := range model.Classes {
		fmt.Println(c.Name)
		for _,inf := range c.Inherits {
			fmt.Println("Inheriting from:",inf)
		}
		for i,cm := range model.Classes {
			if contains(cm.Inherits,c.Name){
				fmt.Println("Inheriting to:", model.Classes[i].Name)
			}
		}
	}
}

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := os.ReadFile(fileName)
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
	r = strings.ReplaceAll(r, "<", "&lt;")
	r = strings.ReplaceAll(r, ">", "&gt;")
	r = strings.TrimSpace(r)
	return r
}

const regex = `<.*?>`

// This method uses a regular expresion to remove HTML tags.
func stripHtmlRegex(s string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(s, "")
}

func returnTextContentsOfTableCell(in string)string{
	return ""
}

func preProcess(text, firstClass string)string{
	f, _ := os.Create("tmp.html")
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("<html><body>")
	tkn := html.NewTokenizer(strings.NewReader(text))
	h4 := 0
	depth := 0
	//classString := ""
	firstClassPassed := false
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return ""
		case tt == html.TextToken:
			if depth>2{
				tTD := tkn.Token()
				t := formatString(tTD.Data)

				w.WriteString(" " + strings.TrimSpace(t))
			}
		case tt == html.EndTagToken:
			t := tkn.Token()
			if t.Data == "table" && firstClassPassed {
				depth--
				w.WriteString("</table>")
			}
			if t.Data == "tr"  && firstClassPassed {
				depth--
				if depth<2 {
					w.WriteString("</tr>")
				}
			}
			if t.Data == "td"  && firstClassPassed {
				depth--
				if depth<3 {
					w.WriteString("</td>")
				}
			}
		case tt == html.StartTagToken:
			t := tkn.Token()
			if t.Data == "table" && firstClassPassed {
				depth++
				w.WriteString("<table>")
				}
			if t.Data == "tr"  && firstClassPassed {
				depth++
				if depth == 2 {
					w.WriteString("<tr>")
				}
			}
			if t.Data == "td"  && firstClassPassed {
				depth++
				if depth==3 {
					w.WriteString("<td>")
				}
			}
			//FirstClass
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
								if strings.Contains(t3.Data, firstClass) {
									firstClassPassed = true
								}
							}
						}
					}
				}
				h4++
			}
		}
	}
	w.WriteString("</body></html>")
	w.Flush()
	return ""
}



func parseBMM(text string, model *Model) (data string) {
	tkn := html.NewTokenizer(strings.NewReader(text))

	h2 := 0
	h3 := 0
	packageString := ""
	isTD1 := false
	isTD2 := false
	isTD3 := false
	isTR := false
	td1 := ""
	td2 := ""
	td3 := ""
	constantName := ""
	constantDescription := ""
	attributeName := ""
	attributeDescription := ""
	functionName := ""
	functionDescription := ""
	constants := false
	attributes := false
	functions := false
	szClassName := ""
	szDescription := ""
	szClassInherits := ""
	szEnumerationName := ""
	constantList := make([]*Constant,0)
	attributeList := make([]*Attribute,0)
	functionList := make([]*Function,0)
	var class *Class
	var enumeration *Enumeration
	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return ""

		case tt == html.EndTagToken:
			t := tkn.Token()
			//TABLE
			if t.Data == "table" {
				if szClassName != "" {
					class, _ = NewClass(szDescription, szClassInherits, szClassName)
					for _, c := range constantList {
						class.AddConstant(c)
					}
					for _, a := range attributeList {
						class.AddAttribute(a)
					}
					for _, f := range functionList {
						class.AddFunction(f)
					}
					constantList = nil
					attributeList = nil
					functionList = nil
					szClassInherits = ""
					szClassName = ""
					szDescription = ""
					model.AddClass(class)
				}else if szEnumerationName != "" {
					enumeration, _ = NewEnumeration(szDescription, szEnumerationName)
					for _, a := range attributeList {
						enumeration.AddAttribute(a)
					}
					attributeList = nil
					szEnumerationName = ""
					szDescription = ""
					model.AddEnumeration(enumeration)
				}

				constants = false
				attributes = false
				functions = false
			}
			//TR
			if t.Data == "tr" && isTR {
				switch {
				case constants:
					constant,e := NewConstantToProcess(constantName, constantDescription)
					if e==nil {
						constantList = append(constantList,constant)
					}
				case attributes:
					var e error
					var attribute *Attribute
					if szEnumerationName != "" {
						attribute, e = NewAttribute(attributeName,"",attributeDescription,"","")
					}else {
						attribute, e = NewAttributeToProcess(attributeName, attributeDescription, "", td1)
					}
					if e==nil {
						attributeList = append(attributeList,attribute)
					}
				case functions:
					parameters := AnalyzeParameters(functionName)
					function,e := NewFunction(functionName,functionDescription)
					if e==nil {
						function.AddParameters(parameters)
						functionList = append(functionList,function)
					}
				}
				isTD1 = false
				isTD2 = false
				isTD3 = false
				isTR = false
			}
			if t.Data == "td" && isTR {
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
					case isTD1 && strings.Compare(td1, "Constants") == 0:
						constants = true
						attributes = false
						functions = false
					case isTD1 && strings.Compare(td1, "Attributes") == 0:
						attributes = true
						constants = false
						functions = false
					case isTD1 && strings.Compare(td1, "Functions") == 0:
						functions = true
						constants = false
						attributes = false
					case isTD2 && strings.TrimSpace(td1)=="Class":
						szClassName = td2
					case isTD2 && strings.TrimSpace(td1)=="Enumeration":
						szEnumerationName = td2
					case isTD2 && strings.TrimSpace(td1)=="Description":
						szDescription = td2
						fmt.Println(">>>>>>>>>>>>>>>>>>>",szDescription)
					case isTD2 && strings.TrimSpace(td1)=="Inherit":
						szClassInherits = td2
					case isTD2 && constants:
						constantName = td2
					case isTD3 && constants:
						constantDescription = td3
					case isTD2 && attributes && szClassName != "":
						attributeName = td2
					case isTD3 && attributes && szClassName != "":
						attributeDescription = td3
						//because of empty td1 field in Enumeration (not required), Enumeration shifts to left
					case isTD1 && attributes && szEnumerationName != "":
						attributeName = td1
					case isTD2 && attributes && szEnumerationName != "":
						attributeDescription = td2
					case isTD2 && functions:
						functionName = td2
					case isTD3 && functions:
						functionDescription = td3
					}
				}
			}
		case tt == html.StartTagToken:
			t := tkn.Token()
			// table
			if t.Data == "table" {
			}
			//TR
			if t.Data == "tr" && !isTR {
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
		}
	}
}
