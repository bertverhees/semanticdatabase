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
	for _,c := range model.Classes {
		c.Print(model)
	}
	for _,e := range model.Enumerations {
		e.Print()
	}
	fmt.Println("----------------------------------------------------")
	//#TODO Enumeration BMM_SCHEMA_STATE and BMM_OPERATOR_POSITION, BMM_PARAMETER_DIRECTION
	//#TODO Missing inheritance BMM_PACKAGE_CONTAINER at BMM_MODEL
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
			if t.Data == "table" && firstClassPassed {
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
			if t.Data == "tr" && isTR && firstClassPassed {
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
						szDescription = szDescription + " " + td2
					case isTD2 && strings.TrimSpace(td1)=="Inherit":
						szClassInherits = td2
						// This section passes more times because of not having plain text in scanned html
						//Like this
						//matching_bmm_models:
						//matching_bmm_models:Hash
						//matching_bmm_models:Hash<String,
						//	matching_bmm_models:Hash<String,BMM_MODEL
						//matching_bmm_models:Hash<String,BMM_MODEL>
					case isTD2 && constants:
						constantName = constantName + td2
					case isTD3 && constants:
						constantDescription = constantDescription + td3
					case isTD2 && attributes && szClassName != "":
						attributeName = attributeName + td2
					case isTD2 && attributes && szEnumerationName != "":
						attributeName = attributeName + td1
					case isTD3 && attributes && szClassName != "":
						attributeDescription = attributeDescription + td3
					case isTD3 && attributes && szEnumerationName != "":
						attributeDescription = attributeDescription + td2
					case isTD2 && functions:
						functionName = functionName + td2
					case isTD3 && functions:
						functionDescription = functionDescription + td3
						//Here we have the complete strings in our variables, finish of this weird section
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
