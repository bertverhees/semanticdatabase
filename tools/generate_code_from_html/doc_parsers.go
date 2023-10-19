package main

import (
	"doc-parser/bmm"
	"doc-parser/p_bmm"
	"os"
)

func main() {
	model := bmm.ParseBMM_HTML()
	CreateFiles("../../vocabulary", model)
	model = p_bmm.ParseP_BMM_HTML()
	CreateFiles("../../vocabulary/v2", model)
	os.Remove("tmp.html")
}
