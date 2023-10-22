package classes

import (
	"errors"
	"fmt"
	"strings"
)

type Model struct {
	Classes             []*Class
	ClassNames          []string
	NewClassNames       []string
	Enumerations        []*Enumeration
	EnumerationNames    []string
	NewEnumerationNames []string
}

func NewModel() *Model {
	model := new(Model)
	model.Classes = make([]*Class, 0)
	model.ClassNames = make([]string, 0)
	model.NewClassNames = make([]string, 0)
	model.Enumerations = make([]*Enumeration, 0)
	model.EnumerationNames = make([]string, 0)
	model.NewEnumerationNames = make([]string, 0)
	return model
}
func (m *Model) AddEnumeration(enumeration *Enumeration) error {
	for _, c := range m.Enumerations {
		if enumeration.Name == c.Name {
			return errors.New("enumeration:" + c.Name + " already added")
		}
	}
	m.Enumerations = append(m.Enumerations, enumeration)
	m.EnumerationNames = append(m.EnumerationNames, enumeration.Name)
	return nil
}

func (m *Model) FindEnumerationByName(name string) *Enumeration {
	for _, c := range m.Enumerations {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func (m *Model) ExistEnumerationByName(name string) bool {
	for _, c := range m.EnumerationNames {
		if c == name {
			return true
		}
	}
	return false
}

func (m *Model) ExistEnumerationByNewName(name string) bool {
	for _, c := range m.NewEnumerationNames {
		if c == name {
			return true
		}
	}
	return false
}

func (m *Model) ExistClassByName(name string) bool {
	for _, c := range m.ClassNames {
		if c == name {
			return true
		}
	}
	return false
}

func (m *Model) ExistClassByNewName(name string) bool {
	for _, c := range m.NewClassNames {
		if c == name {
			return true
		}
	}
	return false
}

func (m *Model) FindClassByName(name string) *Class {
	for _, c := range m.Classes {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func (m *Model) AddClass(class *Class) error {
	for _, c := range m.ClassNames {
		if class.Name == c {
			return errors.New("class:" + c + " already added")
		}
	}
	m.Classes = append(m.Classes, class)
	m.ClassNames = append(m.ClassNames, class.Name)
	return nil
}

type Enumeration struct {
	Name       string
	Comment    string
	Attributes []*Attribute
}

func NewEnumeration(comment, name string) (*Enumeration, error) {
	enumeration := new(Enumeration)
	enumeration.Comment = comment
	enumeration.Name = strings.TrimSpace(name)
	enumeration.Attributes = make([]*Attribute, 0)
	return enumeration, nil
}

func (e *Enumeration) AddAttribute(attribute *Attribute) error {
	e.Attributes = append(e.Attributes, attribute)
	return nil
}

type Class struct {
	Name                string
	Comment             string
	Abstract            bool
	IsGeneric           bool
	GenericPartTypes    []string
	Inherits            []string
	InheritingFromChain []*Class
	AttributeFunctions  []*Function
	Attributes          []*Attribute
	Functions           []*Function
	Constants           []*Constant
}

func (c *Class) Print(m *Model) {
	fmt.Println(c.Name)
	fmt.Println("IsGeneric:", c.IsGeneric)
	if c.IsGeneric {
		for i, p := range c.GenericPartTypes {
			fmt.Println("\t\tGenericPartTypes:", p[i])
		}
	}
	fmt.Println("Abstract:", c.Abstract)
	fmt.Println(c.Comment)
	for _, inf := range c.Inherits {
		fmt.Println("Inheriting from:", inf)
	}
	for i, cm := range m.Classes {
		if Contains(cm.Inherits, c.Name) {
			fmt.Println("Inheriting to:", m.Classes[i].Name)
		}
	}
	for _, co := range c.Constants {
		co.Print()
	}
	for _, a := range c.Attributes {
		a.Print()
	}
	for _, f := range c.Functions {
		f.Print()
	}
}

func AnalyzeGenerics(name string) (IsGeneric bool, GenericPartTypes []string, newName string) {
	//get generic part
	if strings.Contains(name, "<") {
		IsGeneric = true
		genericParts := strings.Split(name, "<")
		newName = strings.TrimSpace(genericParts[0])
		genericPart := strings.TrimSpace(genericParts[1])
		GenericPartType := strings.Split(genericPart, ">")[0]
		if GenericPartType == "" {
			GenericPartType = "any"
		}
		GenericPartTypes = strings.Split(GenericPartType, ",")
	}
	return IsGeneric, GenericPartTypes, newName
}

func CreateInheritanceChain(i1 *Class, m *Model, inheritingChain []*Class) []*Class {
	if len(i1.Inherits) > 0 {
		for _, i_name := range i1.Inherits {
			i2 := m.FindClassByName(i_name)
			if i2 != nil {
				inheritingChain = append(inheritingChain, i2)
				inheritingChain = CreateInheritanceChain(i2, m, inheritingChain)
			}
		}
	}
	return inheritingChain
}

func (e *Enumeration) Print() {
	fmt.Println(e.Name)
	fmt.Println(e.Comment)
	for _, a := range e.Attributes {
		a.Print()
	}
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func NewClass(comment, inherits string, name string, abstract bool) (*Class, error) {
	class := new(Class)
	class.Comment = comment
	class.Name = strings.TrimSpace(name)
	isGeneric, GenericPartTypes, newName := AnalyzeGenerics(class.Name)
	if isGeneric {
		class.IsGeneric = isGeneric
		class.GenericPartTypes = GenericPartTypes
		class.Name = newName
	}
	//inherits-slice always exists
	if inherits != "" {
		class.Inherits = strings.Split(inherits, ",")
	} else {
		class.Inherits = make([]string, 0)
	}
	for i, _ := range class.Inherits {
		class.Inherits[i] = strings.TrimSpace(class.Inherits[i])
	}
	class.Attributes = make([]*Attribute, 0)
	class.Functions = make([]*Function, 0)
	class.AttributeFunctions = make([]*Function, 0)
	class.Constants = make([]*Constant, 0)
	class.Abstract = abstract
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
	Value   string
	Comment string
}

func (c *Constant) Print() {
	fmt.Println("\tConstant:", c.Name)
	fmt.Println("\t\t", c.Type)
	fmt.Println("\t\t", c.Value)
	fmt.Println("\t\t", c.Comment)
}

func NewConstant(name, _type, value, comment string) (*Constant, error) {
	constant := new(Constant)
	constant.Name = name
	constant.Type = _type
	constant.Value = value
	constant.Comment = comment
	return constant, nil
}

// #TODO process
func NewConstantToProcess(name_type_value, comment string) (*Constant, error) {
	var e error
	name, _type, value, e := splitConstantNameType(name_type_value)
	if e != nil {
		return nil, e
	}
	comment = comment
	return NewConstant(name, _type, value, comment)
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func splitConstantNameType(name_type_value string) (string, string, string, error) {
	nameTypeValue := strings.Split(strings.TrimSpace(name_type_value), " : ")
	if len(nameTypeValue) != 2 {
		return "", "", "", errors.New("Not a valid Constant, it should look like: \"name: type=value\"")
	}
	typeValue := strings.Split(strings.TrimSpace(nameTypeValue[1]), "=")
	if len(typeValue) > 1 {
		value := trimQuotes(strings.TrimSpace(typeValue[1]))
		return strings.TrimSpace(nameTypeValue[0]), strings.TrimSpace(typeValue[0]), value, nil
	}
	if len(typeValue) == 1 {
		return strings.TrimSpace(nameTypeValue[0]), strings.TrimSpace(typeValue[0]), "", nil
	}
	return "", "", "", errors.New("Not a valid Constant, it should look like: \"name: type=value\"")
}

type Attribute struct {
	Name             string
	Type             string
	Comment          string
	IsGeneric        bool
	GenericPartTypes []string
	defaultValue     string
	Required         bool
}

func (a *Attribute) Print() {
	fmt.Println("\tAttribute:", a.Name)
	fmt.Println("\t\tIsGeneric:", a.IsGeneric)
	if a.IsGeneric {
		for i, p := range a.GenericPartTypes {
			fmt.Println("\t\tGenericPartTypes:", p[i])
		}
	}
	fmt.Println("\t\t", a.Type)
	fmt.Println("\t\t", a.defaultValue)
	fmt.Println("\t\t", a.Required)
	fmt.Println("\t\t", a.Comment)
}

func NewAttribute(name, _type, comment string, defaultValue string, required string) (*Attribute, error) {
	if name == "" {
		return nil, errors.New("Attribute with empty name is not allowed")
	}
	attribute := new(Attribute)
	attribute.Name = name
	attribute.Type = _type
	attribute.Comment = comment
	attribute.defaultValue = defaultValue
	IsGeneric, GenericPartTypes, _ := AnalyzeGenerics(_type)
	//if !strings.Contains(strings.ToLower(_type), "hash") {
	if IsGeneric {
		attribute.IsGeneric = IsGeneric
		attribute.GenericPartTypes = GenericPartTypes
		//attribute.Name = newName
	}
	//}
	if strings.TrimSpace(required) == "1..1" {
		attribute.Required = true
	} else {
		attribute.Required = false
	}
	return attribute, nil
}

func NewAttributeToProcess(name_type, comment string, defaultValue string, required string) (*Attribute, error) {
	var e error
	var name, _type string
	name, _type, defaultValue, e = splitAttributeNameType(name_type)
	if e != nil {
		return nil, e
	}
	return NewAttribute(name, _type, comment, defaultValue, required)
}

func splitAttributeNameType(name_type string) (name string, _type string, defaultValue string, e error) {
	nameType := strings.Split(strings.TrimSpace(name_type), ":")
	if len(nameType) != 2 {
		return "", "", "", errors.New("Not a valid attribute, it should look like: \"name: type\"")
	}
	name = strings.TrimSpace(strings.ReplaceAll(nameType[0], "()", ""))
	typeDefault := strings.Split(strings.TrimSpace(nameType[1]), "{")
	_type = strings.TrimSpace(typeDefault[0])
	if len(typeDefault) > 1 {
		defaultValues := strings.Split(typeDefault[1], "=")
		if len(defaultValues) > 1 {
			defaultValue = trimQuotes(strings.TrimSpace(strings.Trim(strings.TrimSpace(defaultValues[1]), "}")))
			return name, _type, defaultValue, nil
		} else {
			return "", "", "", errors.New("Not a valid attribute, error in defaultValue")
		}
	} else {
		return name, _type, "", nil
	}
}

type Function struct {
	Name      string
	Out       []*Parameter
	In        []*Parameter
	Comment   string
	Redefined bool
}

func (f *Function) Print() {
	fmt.Println("\tFunction:", f.Name)
	fmt.Println("\tRedefined: ", f.Redefined)
	fmt.Println("\tIn:")
	for _, p := range f.In {
		fmt.Println("\t\t", p.Name, p.Type, p.InOut)
		fmt.Println("\t\tIsGeneric:", p.IsGeneric)
		if p.IsGeneric {
			for i, q := range p.GenericPartTypes {
				fmt.Println("\t\tGenericPartTypes:", q[i])
			}
		}
	}
	switch {
	case len(f.Out) == 1:
		fmt.Println("\tOut:")
		if f.Out[0].Type != "" {
			fmt.Println("\t\t", f.Out[0].Type, f.Out[0].InOut)
			fmt.Println("IsGeneric:", f.Out[0].IsGeneric)
			if f.Out[0].IsGeneric {
				for i, q := range f.Out[0].GenericPartTypes {
					fmt.Println("\t\tGenericPartTypes:", q[i])
				}
			}
		}
	case len(f.Out) > 1:
		fmt.Println("\tOut:")
		for i, p := range f.Out {
			if i == 0 {
				fmt.Println("\tOut:")
				fmt.Println("\tIsGeneric:", p.IsGeneric)
				if p.IsGeneric {
					for i, q := range p.GenericPartTypes {
						fmt.Println("\t\tGenericPartTypes:", q[i])
					}
				}

			}
			if p.Type != "" {
				fmt.Println("\t\t", p.Type, p.InOut)
				fmt.Println("\t\tIsGeneric:", p.IsGeneric)
				if p.IsGeneric {
					for i, q := range p.GenericPartTypes {
						fmt.Println("\t\tGenericPartTypes:", q[i])
					}
				}
			}
		}
	}
	fmt.Println("\t\t", f.Comment)
}

func NewFunction(name, comment string) (*Function, error) {
	functionName := name
	if strings.Contains(name, "(") {
		functionName = strings.TrimSpace(strings.Split(name, "(")[0])
	}
	if strings.Contains(functionName, ":") {
		functionName = strings.TrimSpace(strings.Split(name, ":")[0])
	}
	function := new(Function)
	function.Name = functionName
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
	for _, p := range parameters {
		e := f.AddParameter(p)
		if e != nil {
			return e
		}
	}
	return nil
}

type Parameter struct {
	Name             string
	InOut            string
	Type             string
	Required         bool
	IsGeneric        bool
	GenericPartTypes []string
}

func NewParameter(name, inOut, _type string, required bool) (*Parameter, error) {
	isGeneric, genericPartTypes, newName := AnalyzeGenerics(name)
	IsGeneric := false
	GenericPartTypes := make([]string, 0)
	if isGeneric {
		IsGeneric = isGeneric
		GenericPartTypes = genericPartTypes
		name = newName
	}
	return &Parameter{name, inOut, _type, required, IsGeneric, GenericPartTypes}, nil
}

func AnalyzeParameters(functionName string) []*Parameter {
	parameters := make([]*Parameter, 0)
	outType := functionName
	if strings.Contains(outType, "(") && strings.Contains(outType, ")") {
		outType = outType[strings.Index(outType, ")")+1:]
	} else {
		outType = ""
	}
	outType = strings.TrimSpace(outType[strings.Index(outType, ":")+1:])
	var parameter *Parameter
	if outType != "" {
		parameter, _ = NewParameter("", "out", outType, false)
		parameters = append(parameters, parameter)
	}
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