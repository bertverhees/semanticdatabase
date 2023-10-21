package main

import (
	"doc-parser/bmm"
	"doc-parser/p_bmm"
	"doc-parser/string_handling"
	"os"
)

func main() {
	model := bmm.ParseBMM_HTML()
	string_handling.CreateFiles("vocabulary", "../../vocabulary", model)
	model = p_bmm.ParseP_BMM_HTML(model)
	string_handling.CreateFiles("v2", "../../vocabulary/v2", model)
	os.Remove("tmp.html")
}
