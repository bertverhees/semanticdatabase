package main

import (
	"errors"
	"strings"
)

type Class struct {
	Comment    string
	Inherits   []string
	Attributes []Attribute
	Functions  []Function
	Constants  []Constant
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
	class.Constants = make([]Constant, 0)
	return class, nil
}

func (c *Class) AddAttribute(attribute Attribute) error {
	c.Attributes = append(c.Attributes, attribute)
	return nil
}

func (c *Class) AddFunction(function Function) error {
	c.Functions = append(c.Functions, function)
	return nil
}

func (c *Class) AddConstant(constant Constant) error {
	c.Constants = append(c.Constants, constant)
	return nil
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

func (f *Function) AddParameter(parameter Parameter) error {
	switch {
	case strings.ToLower(parameter.InOut) == "out":
		f.Out = append(f.Out, parameter)
	case strings.ToLower(parameter.InOut) == "in":
		f.In = append(f.In, parameter)
	default:
		return errors.New("Inout of Parameter must be \"in\" or \"out\" (without quotes, case insentive)")
	}
	return nil
}

type Parameter struct {
	Name     string
	InOut    string
	Type     string
	Required bool
}

func NewParameter(name, inOut, _type string, required bool) (*Parameter, error) {
	return &Parameter{name,inOut,_type,required}, nil
}

func AnalyzeParameters(functionName string) []Parameter {
	parameters := make([]Parameter, 0)
	outType := functionName
	if strings.Contains(outType, "(") && strings.Contains(outType, ")") {
		outType = outType[strings.Index(outType, ")"):]
	}
	outType = strings.TrimSpace(outType[strings.Index(outType, ":")+1:])
	parameter, _ := NewParameter("", "out", outType, false)
	parameters = append(parameters, *parameter)
	if strings.Contains(functionName, "(") && strings.Contains(functionName, ")") {
		functionParameters := functionName[strings.Index(functionName, "(")+1 : strings.Index(functionName, ")")]
		if functionParameters != "" {
			typeSeparator := " "
			if strings.Contains(functionParameters, ":") {
				typeSeparator = ":"
			}
			ps := strings.Split(functionParameters, ",")
			if len(ps) > 0 {
				for i,psl := range ps {
					parameterName := ""
					parameterType := ""
					nameType := strings.Split(strings.TrimSpace(psl), typeSeparator)
					parameterName = strings.TrimSpace(nameType[0])
					if len(nameType) == 1 {
						if i<len(ps) {
							for j := i+1; j < len(ps); j++ {
								nameType = strings.Split(strings.TrimSpace(ps[j]), typeSeparator)
								if len(nameType)>1 {
									parameterType = strings.TrimSpace(nameType[1])
									break
								}
							}
						}
					} else {
						parameterType = strings.TrimSpace(nameType[1])
					}
					parameter, _ = NewParameter(parameterName, "in", parameterType, false)
					parameters = append(parameters, *parameter)
				}
			}
		}
	}
	return parameters
}
