package main

import (
	"fmt"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	fileName := "bmm.html"
	text, err := readHtmlFromFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	model := NewModel()
	parseBMM(text, model)
	fmt.Println(len(model.Classes))
	for i,c := range model.Classes {
		fmt.Println(i,c)
	}
}

func TestAnalyzeParameters(t *testing.T) {
	f := "fun(instring1,instring2 string):out"
	p1,_ := NewParameter("","out","out",false)
	p2,_ := NewParameter("instring1","in","string",false)
	p3,_ := NewParameter("instring2","in","string",false)
	p := AnalyzeParameters(f)
	if p1 !=p[0]{
		t.Error("structs should be equal")
	}
	if p2 !=p[1]{
		t.Error("structs should be equal")
	}
	if p3 !=p[2]{
		t.Error("structs should be equal")
	}
	f = "fun(instring1 string,insint1 int):out"
	p1,_ = NewParameter("","out","out",false)
	p2,_ = NewParameter("instring1","in","string",false)
	p3,_ = NewParameter("insint1","in","int",false)
	p = AnalyzeParameters(f)
	if p1 !=p[0]{
		t.Error("structs should be equal")
	}
	if p2 !=p[1]{
		t.Error("structs should be equal")
	}
	if p3 !=p[2]{
		t.Error("structs should be equal")
	}
	f = "fun(instring1, instring2 string,insint1 int):out"
	p1,_ = NewParameter("","out","out",false)
	p2,_ = NewParameter("instring1","in","string",false)
	p3,_ = NewParameter("instring2","in","string",false)
	p4,_ := NewParameter("insint1","in","int",false)
	p = AnalyzeParameters(f)
	if p1 !=p[0]{
		t.Error("structs should be equal")
	}
	if p2 !=p[1]{
		t.Error("structs should be equal")
	}
	if p3 !=p[2]{
		t.Error("structs should be equal")
	}
	if p4 !=p[3]{
		t.Error("structs should be equal")
	}
	f = "fun(insint0 int, instring1, instring2 string,insint1 int):out"
	p1,_ = NewParameter("","out","out",false)
	p2,_ = NewParameter("insint0","in","int",false)
	p3,_ = NewParameter("instring1","in","string",false)
	p4,_ = NewParameter("instring2","in","string",false)
	p5,_ := NewParameter("insint1","in","int",false)
	p = AnalyzeParameters(f)
	if p1 !=p[0]{
		t.Error("structs should be equal")
	}
	if p2 !=p[1]{
		t.Error("structs should be equal")
	}
	if p3 !=p[2]{
		t.Error("structs should be equal")
	}
	if p4 !=p[3]{
		t.Error("structs should be equal")
	}
	if p5 !=p[4]{
		t.Error("structs should be equal")
	}
	f = "fun(insint0 int, instring1, instring2 string,insint1, insint2 int):out"
	p1,_ = NewParameter("","out","out",false)
	p2,_ = NewParameter("insint0","in","int",false)
	p3,_ = NewParameter("instring1","in","string",false)
	p4,_ = NewParameter("instring2","in","string",false)
	p5,_ = NewParameter("insint1","in","int",false)
	p6,_ := NewParameter("insint2","in","int",false)
	p = AnalyzeParameters(f)
	if p1 !=p[0]{
		t.Error("structs should be equal")
	}
	if p2 !=p[1]{
		t.Error("structs should be equal")
	}
	if p3 !=p[2]{
		t.Error("structs should be equal")
	}
	if p4 !=p[3]{
		t.Error("structs should be equal")
	}
	if p5 !=p[4]{
		t.Error("structs should be equal")
	}
	if p6 !=p[5]{
		t.Error("structs should be equal")
	}
	f = "fun(instring1,instring2: string):out"
	p1,_ = NewParameter("","out","out",false)
	p2,_ = NewParameter("instring1","in","string",false)
	p3,_ = NewParameter("instring2","in","string",false)
	p = AnalyzeParameters(f)
	if p1 !=p[0]{
		t.Error("structs should be equal")
	}
	if p2 !=p[1]{
		t.Error("structs should be equal")
	}
	if p3 !=p[2]{
		t.Error("structs should be equal")
	}
	f = "fun(instring1: string,insint1: int):out"
	p1,_ = NewParameter("","out","out",false)
	p2,_ = NewParameter("instring1","in","string",false)
	p3,_ = NewParameter("insint1","in","int",false)
	p = AnalyzeParameters(f)
	if p1 !=p[0]{
		t.Error("structs should be equal")
	}
	if p2 !=p[1]{
		t.Error("structs should be equal")
	}
	if p3 !=p[2]{
		t.Error("structs should be equal")
	}
	f = "fun(instring1, instring2 :string,insint1 :int):out"
	p1,_ = NewParameter("","out","out",false)
	p2,_ = NewParameter("instring1","in","string",false)
	p3,_ = NewParameter("instring2","in","string",false)
	p4,_ = NewParameter("insint1","in","int",false)
	p = AnalyzeParameters(f)
	if p1 !=p[0]{
		t.Error("structs should be equal")
	}
	if p2 !=p[1]{
		t.Error("structs should be equal")
	}
	if p3 !=p[2]{
		t.Error("structs should be equal")
	}
	if p4 !=p[3]{
		t.Error("structs should be equal")
	}
	f = "fun(insint0: int, instring1, instring2: string,insint1: int):out"
	p1,_ = NewParameter("","out","out",false)
	p2,_ = NewParameter("insint0","in","int",false)
	p3,_ = NewParameter("instring1","in","string",false)
	p4,_ = NewParameter("instring2","in","string",false)
	p5,_ = NewParameter("insint1","in","int",false)
	p = AnalyzeParameters(f)
	if p1 !=p[0]{
		t.Error("structs should be equal")
	}
	if p2 !=p[1]{
		t.Error("structs should be equal")
	}
	if p3 !=p[2]{
		t.Error("structs should be equal")
	}
	if p4 !=p[3]{
		t.Error("structs should be equal")
	}
	if p5 !=p[4]{
		t.Error("structs should be equal")
	}
	f = "fun(insint0: int, instring1, instring2: string,insint1, insint2: int):out"
	p1,_ = NewParameter("","out","out",false)
	p2,_ = NewParameter("insint0","in","int",false)
	p3,_ = NewParameter("instring1","in","string",false)
	p4,_ = NewParameter("instring2","in","string",false)
	p5,_ = NewParameter("insint1","in","int",false)
	p6,_ = NewParameter("insint2","in","int",false)
	p = AnalyzeParameters(f)
	if p1 !=p[0]{
		t.Error("structs should be equal")
	}
	if p2 !=p[1]{
		t.Error("structs should be equal")
	}
	if p3 !=p[2]{
		t.Error("structs should be equal")
	}
	if p4 !=p[3]{
		t.Error("structs should be equal")
	}
	if p5 !=p[4]{
		t.Error("structs should be equal")
	}
	if p6 !=p[5]{
		t.Error("structs should be equal")
	}
}

