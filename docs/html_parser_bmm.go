package main

//https://zetcode.com/golang/net-html/
import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fileName := "bmm.html"
	text, err := readHtmlFromFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	model := NewModel()
	parseBMM(text, model)
	fmt.Println(len(model.Classes))

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

func parseBMM(text string, model *Model) (data string) {
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
	classDescriptionInTD1 := false
	classInheritInTD1 := false
	classNameInTD1 := false
	constants := false
	attributes := false
	functions := false
	szClassName := ""
	szClassDescription := ""
	szClassInherits := ""
	classNamePassed := false
	classInheritPassed := false
	classDescriptionPassed := false
	var class *Class
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
					constant,_ := NewConstantToProcess(constantName, constantDescription)
					class.AddConstant(constant)
				case attributes:
					//#TODO Process required
					attribute,_ := NewAttributeToProcess(attributeName, attributeDescription,false)
					class.AddAttribute(attribute)
				case functions:
					parameters := AnalyzeParameters(functionName)
					function,_ := NewFunction(functionName,functionDescription)
					function.AddParameters(parameters)
				}
				isTD1 = false
				isTD2 = false
				isTD3 = false
				isTR = false
			}
			if t.Data == "td" && isTR && firstClassPassed {
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
					case isTD1 && strings.Compare(td1, "Class") == 0:
						classNameInTD1 = true
						classNamePassed = false
					case isTD1 && strings.Compare(td1, "Description") == 0:
						classDescriptionInTD1 = true
						classDescriptionPassed = false
					case isTD1 && strings.Compare(td1, "Inherit") == 0:
						classInheritInTD1 = true
						classInheritPassed = false
					case isTD1 && strings.Compare(td1, "Constants") == 0:
						constants = true
					case isTD1 && strings.Compare(td1, "Attributes") == 0:
						attributes = true
					case isTD1 && strings.Compare(td1, "Functions") == 0:
						functions = true
					case isTD2 && classNameInTD1:
						//fmt.Println("Class: " + td2)
						szClassName = td2
						classNameInTD1 = false
						classNamePassed = true
					case isTD2 && classDescriptionInTD1:
						//fmt.Println("ClassDescription: " + td2)
						szClassDescription = td2
						classDescriptionInTD1 = false
						classDescriptionPassed = true
					case isTD2 && classInheritInTD1:
						//fmt.Println("classInheritInTD1: " + td2)
						szClassInherits = td2
						classInheritInTD1 = false
						classInheritPassed = true
					case isTD1 && constants:
						//fmt.Println("constants occurence: " + td1)
					case isTD2 && constants:
						constantName = constantName + td2
					case isTD3 && constants:
						constantDescription = constantDescription + td3
					case isTD1 && attributes:
						//fmt.Println("attribute occurence: " + td1)
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
					if classNamePassed && classInheritPassed && classDescriptionPassed {
						class,_ = NewClass(szClassDescription, szClassInherits, szClassName)
						model.AddClass(class)
						classNamePassed = false
						classInheritPassed = false
						classDescriptionPassed = false
					}
					switch {
					case isTD1 && (strings.Compare(td1, "Attributes") == 0 || strings.Compare(td1, "Functions") == 0 || strings.Compare(td1, "Class") == 0):
						constants = false
					case isTD1 && (strings.Compare(td1, "Constants") == 0 || strings.Compare(td1, "Functions") == 0 || strings.Compare(td1, "Class") == 0):
						attributes = false
					case isTD1 && (strings.Compare(td1, "Constants") == 0 || strings.Compare(td1, "Attributes") == 0 || strings.Compare(td1, "Class") == 0):
						functions = false

					}
				}
			}
		case tt == html.StartTagToken:
			t := tkn.Token()
			// table
			if t.Data == "table" && firstClassPassed {
			}
			//TR
			if t.Data == "tr" && !isTR && firstClassPassed {
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
								// 	classDescriptionInTD1 = true
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
