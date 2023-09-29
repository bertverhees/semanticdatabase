package main

type Class struct {
	Comment    string
	Inherits   []Class
	Attributes []Attribute
	Functions  []Function
}

type Constant struct {
	Name    string
	Type    string
	Comment string
}

type Attribute struct {
	Name    string
	Type    string
	Comment string
}

type Function struct {
	Name    string
	Out     []Parameter
	In      []Parameter
	Comment string
}

type Parameter struct {
	in_out  string
	Comment string
}
