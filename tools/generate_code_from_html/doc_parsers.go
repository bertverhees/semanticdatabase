package main

import (
	"doc-parser/bmm"
	"doc-parser/p_bmm"
	"doc-parser/string_handling"
	"os"
)

func main() {
	model := bmm.ParseBMM_HTML()
	model = p_bmm.ParseP_BMM_HTML(model)
	for _,c := range model.Classes {
		model.NewClassNames = append(model.NewClassNames, 	string_handling.UpperCamelCase(c.Name))
	}
	for _,e := range model.Enumerations {
		model.NewEnumerationNames = append(model.EnumerationNames, 	string_handling.UpperCamelCase(e.Name))
	}
	string_handling.CreateFiles("vocabulary", "../../vocabulary", model)
	string_handling.CreateFiles("v2", "../../vocabulary/v2", model)
	os.Remove("tmp.html")
}


