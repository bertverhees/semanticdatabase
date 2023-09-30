package main

import (
	"testing"
)

func TestParameters(t *testing.T) {
	f := "fun(instring1,instring2 string):out"
	AnalyzeParameters(f)
	f = "fun(instring1 string,insint1 int):out"
	AnalyzeParameters(f)
	f = "fun(instring1, instring2 string,insint1 int):out"
	AnalyzeParameters(f)
	f = "fun(insint0 int, instring1, instring2 string,insint1 int):out"
	AnalyzeParameters(f)
	f = "fun(insint0 int, instring1, instring2 string,insint1, insint2 int):out"
	AnalyzeParameters(f)
	f = "fun(instring1,instring2 :string):out"
	AnalyzeParameters(f)
	f = "fun(instring1: string,insint1: int):out"
	AnalyzeParameters(f)
	f = "fun(instring1, instring2 :string,insint1 :int):out"
	AnalyzeParameters(f)
	f = "fun(insint0: int, instring1, instring2: string,insint1: int):out"
	AnalyzeParameters(f)
	f = "fun(insint0: int, instring1, instring2: string,insint1, insint2: int):out"
	AnalyzeParameters(f)
}
