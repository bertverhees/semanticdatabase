package main

import "strings"

type Class struct {
	Comment    string
	Inherits   []string
	Attributes []Attribute
	Functions  []Function
}

func NewClass(comment, inherits string) (*Class, error) {
	var class *Class
	class = new(Class)
	class.Comment = comment
	//inherits-slice always exists
	if inherits != "" {
		class.Inherits = strings.Split(inherits, ",")
	} else {
		class.Inherits = make([]string, 0)
	}
	class.Attributes = make([]Attribute, 0)
	class.Functions = make([]Function, 0)
	return class, nil
}

type Constant struct {
	Name    string
	Type    string
	Comment string
}

func NewConstant(name, _type, comment string) (*Constant, error) {
	var constant *Constant
	constant.Name = name
	constant.Type = _type
	constant.Comment = comment
	return constant, nil
}

type Attribute struct {
	Name     string
	Type     string
	Comment  string
	Required bool
}

func NewAttribute(name, _type, comment string, required bool) (*Attribute, error) {
	var attribute *Attribute
	attribute.Name = name
	attribute.Type = _type
	attribute.Comment = comment
	attribute.Required = required
	return attribute, nil
}

type Function struct {
	Name    string
	Out     []Parameter
	In      []Parameter
	Comment string
}

func NewFunction(name, comment string) (*Function, error) {
	var function *Function
	function.Name = name
	function.Comment = comment
	function.Out = make([]Parameter, 0)
	function.In = make([]Parameter, 0)
	return function, nil
}

type Parameter struct {
	Name     string
	In_out   string
	Type     string
	Required bool
}

func NewParameter(name, in_out, _type string, required bool) (*Parameter, error) {
	var parameter *Parameter
	parameter.Name = name
	parameter.In_out = name
	parameter.Type = _type
	parameter.Required = required
	return parameter, nil
}
