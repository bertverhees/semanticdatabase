package main

import (
	"doc-parser/bmm"
	"doc-parser/p_bmm"
	"os"
)

func main() {
	WriteBasicDefinitions("vocabulary", "../../vocabulary")
	model := bmm.ParseBMM_HTML()
	CreateFiles("vocabulary", "../../vocabulary", model)
	model = p_bmm.ParseP_BMM_HTML(model)
	CreateFiles("v2", "../../vocabulary/v2", model)
	os.Remove("tmp.html")
}
