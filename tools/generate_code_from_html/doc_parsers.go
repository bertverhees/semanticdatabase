package main

import (
	"doc-parser/bmm"
	"doc-parser/p_bmm"
	"os"
)

func main(){
	bmm.ParseBMM_HTML()
	p_bmm.ParseP_BMM_HTML()
	os.Remove("tmp.html")
}
