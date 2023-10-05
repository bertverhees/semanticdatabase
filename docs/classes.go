package main

import (
	"errors"
	"strings"
)

type Model struct {
	Classes []*Class
}

func NewModel() *Model {
	model := new(Model)
	model.Classes = make([]*Class, 0)
	return model
}

func (m *Model) AddClass(class *Class) error {
	m.Classes = append(m.Classes, class)
	return nil
}

type Class struct {
	Name string
	Comment    string
	Inherits   []string
	Attributes []*Attribute
	Functions  []*Function
	Constants  []*Constant
}

func NewClass(comment, inherits string, name string) (*Class, error) {
	class := new(Class)
	class.Comment = comment
	class.Name = name
	//inherits-slice always exists
	if inherits != "" {
		class.Inherits = strings.Split(inherits, ",")
	} else {
		class.Inherits = make([]string, 0)
	}
	for i,_ := range class.Inherits {
		class.Inherits[i] = strings.TrimSpace(class.Inherits[i])
	}
	class.Attributes = make([]*Attribute, 0)
	class.Functions = make([]*Function, 0)
	class.Constants = make([]*Constant, 0)
	return class, nil
}

func (c *Class) AddAttribute(attribute *Attribute) error {
	c.Attributes = append(c.Attributes, attribute)
	return nil
}

func (c *Class) AddFunction(function *Function) error {
	c.Functions = append(c.Functions, function)
	return nil
}

func (c *Class) AddConstant(constant *Constant) error {
	c.Constants = append(c.Constants, constant)
	return nil
}

type Constant struct {
	Name    string
	Type    string
	Value	string
	Comment string
}

func NewConstant(name, _type, value, comment string) (*Constant, error) {
	constant := new(Constant)
	constant.Name = name
	constant.Type = _type
	constant.Value = value
	constant.Comment = comment
	return constant, nil
}
	//#TODO process
func NewConstantToProcess(name_type_value, comment string) (*Constant, error) {
	constant := new(Constant)
	var e error
	constant.Name, constant.Type, constant.Value, e = splitConstantNameType(name_type_value)
	constant.Comment = comment
	return constant, e
}
func splitConstantNameType(name_type_value string)(string, string, string, error){
	nameTypeValue := strings.Split(strings.TrimSpace(name_type_value), ":")
	if len(nameTypeValue)!=2{
		return "","","", errors.New("Not a valid Constant, it should look like: \"name: type=value\"")
	}
	typeValue := strings.Split(strings.TrimSpace(name_type_value), "=")
	if len(typeValue)>1 {
		return nameTypeValue[0], typeValue[0], typeValue[1], nil
	}
	if len(typeValue)==1 {
		return nameTypeValue[0], typeValue[0], "", nil
	}
	return "", "", "", errors.New("Not a valid Constant, it should look like: \"name: type=value\"")
}

type Attribute struct {
	Name     string
	Type     string
	Comment  string
	Required bool
}

func NewAttribute(name, _type, comment string, required bool) (*Attribute, error) {
	attribute := new(Attribute)
	attribute.Name = name
	attribute.Type = _type
	attribute.Comment = comment
	attribute.Required = required
	return attribute, nil
}

func NewAttributeToProcess(name_type, comment string, required bool) (*Attribute, error) {
	attribute := new(Attribute)
	var e error
	attribute.Name, attribute.Type, e = splitAttributeNameType(name_type)
	attribute.Comment = comment
	attribute.Required = required
	return attribute, e
}

func splitAttributeNameType(name_type string)(string, string, error){
	nameType := strings.Split(strings.TrimSpace(name_type), ":")
	if len(nameType)!=2{
		return "","", errors.New("Not a valid attribute, it should look like: \"name: type\"")
	}
	return nameType[0],nameType[1],nil
}

type Function struct {
	Name    string
	Out     []*Parameter
	In      []*Parameter
	Comment string
}

func NewFunction(name, comment string) (*Function, error) {
	function := new(Function)
	function.Name = name
	function.Comment = comment
	function.Out = make([]*Parameter, 0)
	function.In = make([]*Parameter, 0)
	return function, nil
}

func (f *Function) AddParameter(parameter *Parameter) error {
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

func (f *Function) AddParameters(parameters []*Parameter) error {
	for _,p := range parameters {
		e := f.AddParameter(p)
		if e !=nil{
			return e
		}
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
	return &Parameter{name, inOut, _type, required}, nil
}

func AnalyzeParameters(functionName string) []*Parameter {
	parameters := make([]*Parameter, 0)
	outType := functionName
	if strings.Contains(outType, "(") && strings.Contains(outType, ")") {
		outType = outType[strings.Index(outType, ")"):]
	}
	outType = strings.TrimSpace(outType[strings.Index(outType, ":")+1:])
	parameter, _ := NewParameter("", "out", outType, false)
	parameters = append(parameters, parameter)
	if strings.Contains(functionName, "(") && strings.Contains(functionName, ")") {
		functionParameters := functionName[strings.Index(functionName, "(")+1 : strings.Index(functionName, ")")]
		if functionParameters != "" {
			typeSeparator := " "
			if strings.Contains(functionParameters, ":") {
				typeSeparator = ":"
			}
			ps := strings.Split(functionParameters, ",")
			if len(ps) > 0 {
				for i, psl := range ps {
					parameterName := ""
					parameterType := ""
					nameType := strings.Split(strings.TrimSpace(psl), typeSeparator)
					parameterName = strings.TrimSpace(nameType[0])
					if len(nameType) == 1 {
						if i < len(ps) {
							for j := i + 1; j < len(ps); j++ {
								nameType = strings.Split(strings.TrimSpace(ps[j]), typeSeparator)
								if len(nameType) > 1 {
									parameterType = strings.TrimSpace(nameType[1])
									break
								}
							}
						}
					} else {
						parameterType = strings.TrimSpace(nameType[1])
					}
					parameter, _ = NewParameter(parameterName, "in", parameterType, false)
					parameters = append(parameters, parameter)
				}
			}
		}
	}
	return parameters
}
