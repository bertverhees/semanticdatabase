package vocabulary

import "errors"

// Persistent form of BMM_ENUMERATION_STRING .

// Interface definition
type IPBmmEnumerationString interface {
	IPBmmEnumeration
}

// Struct definition
type PBmmEnumerationString struct {
	// embedded for Inheritance
	PBmmEnumeration
	// Attributes
	/**
	BMM_CLASS object build by create_bmm_class_definition and
	populate_bmm_class_definition .
	*/
	bmmClass IBmmEnumerationString `yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

func (b *PBmmEnumerationString) SetBmmClass(bmmClass IBmmClass) error {
	s, ok := bmmClass.(IBmmEnumerationString)
	if !ok {
		return errors.New("type-assertion  for IBmmEnumerationString in PBmmEnumerationString->SetBmmClass failed")
	} else {
		b.bmmClass = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmEnumerationString() *PBmmEnumerationString {
	pbmmenumerationstring := new(PBmmEnumerationString)
	// Constants
	// From: PBmmEnumeration
	// From: PBmmClass
	// From: PBmmModelElement
	return pbmmenumerationstring
}

// BUILDER
type PBmmEnumerationStringBuilder struct {
	pbmmenumerationstring *PBmmEnumerationString
	errors                []error
}

func NewPBmmEnumerationStringBuilder() *PBmmEnumerationStringBuilder {
	return &PBmmEnumerationStringBuilder{
		pbmmenumerationstring: NewPBmmEnumerationString(),
		errors:                make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
BMM_CLASS object build by create_bmm_class_definition and
populate_bmm_class_definition .
*/
func (i *PBmmEnumerationStringBuilder) SetBmmClass(v IBmmEnumerationString) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.bmmClass = v
	return i
}

// //From: PBmmEnumeration
func (i *PBmmEnumerationStringBuilder) SetItemNames(v []string) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.itemNames = v
	return i
}
func (i *PBmmEnumerationStringBuilder) SetItemValues(v []any) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.itemValues = v
	return i
}

// //From: PBmmClass
// name of the class. Persisted attribute.
func (i *PBmmEnumerationStringBuilder) SetName(v string) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.name = v
	return i
}

/*
*
List of immediate inheritance parents. If there are generic ancestors, use
ancestor_defs instead. Persisted attribute.
*/
func (i *PBmmEnumerationStringBuilder) SetAncestors(v []string) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.ancestors = v
	return i
}

// List of attributes defined in this class. Persistent attribute.
func (i *PBmmEnumerationStringBuilder) SetProperties(v map[string]IPBmmProperty) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.properties = v
	return i
}

// True if this is an abstract type. Persisted attribute.
func (i *PBmmEnumerationStringBuilder) SetIsAbstract(v bool) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.isAbstract = v
	return i
}

// True if this class definition overrides one found in an included schema.
func (i *PBmmEnumerationStringBuilder) SetIsOverride(v bool) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.isOverride = v
	return i
}

// List of generic parameter definitions. Persisted attribute.
func (i *PBmmEnumerationStringBuilder) SetGenericParameterDefs(v map[string]IPBmmGenericParameter) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.genericParameterDefs = v
	return i
}

/*
*
Reference to original source schema defining this class. Set during BMM_SCHEMA
materialise. Useful for GUI tools to enable user to edit the schema file
containing a given class (i.e. taking into account that a class may be in any of
the schemas in a schema inclusion hierarchy).
*/
func (i *PBmmEnumerationStringBuilder) SetSourceSchemaId(v string) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.sourceSchemaId = v
	return i
}

/*
*
Unique id generated for later comparison during merging, in order to detect if
two classes are the same. Assigned in post-load processing.
*/
func (i *PBmmEnumerationStringBuilder) SetUid(v int) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.uid = v
	return i
}

/*
*
List of structured inheritance ancestors, used only in the case of generic
inheritance. Persisted attribute.
*/
//func (i *PBmmEnumerationIntegerBuilder) SetAncestorDefs(v []IPBmmGenericType) *PBmmEnumerationIntegerBuilder {
//	i.pbmmenumerationinteger.AncestorDefs = v
//	return i
//}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmEnumerationStringBuilder) SetDocumentation(v string) *PBmmEnumerationStringBuilder {
	i.pbmmenumerationstring.documentation = v
	return i
}

func (i *PBmmEnumerationStringBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmEnumerationStringBuilder) Build() *PBmmEnumerationString {
	return i.pbmmenumerationstring
}

//FUNCTIONS
