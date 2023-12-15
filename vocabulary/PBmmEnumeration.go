package vocabulary

import "errors"

// Persistent form of BMM_ENUMERATION attributes.

// Interface definition
type IPBmmEnumeration interface {
	IPBmmClass
}

// Struct definition
type PBmmEnumeration struct {
	// embedded for Inheritance
	PBmmClass
	// Constants
	// Attributes
	itemNames  []string `yaml:"itemnames" json:"itemnames" xml:"itemnames"`
	itemValues []any    `yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
	/**
	BMM_CLASS object build by create_bmm_class_definition and
	populate_bmm_class_definition .
	*/
	bmmClass IBmmEnumeration `yaml:"bmmclass" json:"bmmclass" xml:"bmmclass"`
}

func (p *PBmmEnumeration) ItemNames() []string {
	return p.itemNames
}

func (p *PBmmEnumeration) SetItemNames(itemNames []string) error {
	p.itemNames = itemNames
	return nil
}

func (p *PBmmEnumeration) ItemValues() []any {
	return p.itemValues
}

func (p *PBmmEnumeration) SetItemValues(itemValues []any) error {
	p.itemValues = itemValues
	return nil
}

func (b *PBmmEnumeration) SetBmmClass(bmmClass IBmmClass) error {
	s, ok := bmmClass.(IBmmEnumeration)
	if !ok {
		return errors.New("type-assertion  for IBmmEnumeration in PBmmEnumeration->SetBmmClass failed")
	} else {
		b.bmmClass = s
		return nil
	}
}

// CONSTRUCTOR
func NewPBmmEnumeration() *PBmmEnumeration {
	pbmmenumeration := new(PBmmEnumeration)
	// Constants
	// From: PBmmClass
	// From: PBmmModelElement
	return pbmmenumeration
}

// BUILDER
type PBmmEnumerationBuilder struct {
	pbmmenumeration *PBmmEnumeration
	errors          []error
}

func NewPBmmEnumerationBuilder() *PBmmEnumerationBuilder {
	return &PBmmEnumerationBuilder{
		pbmmenumeration: NewPBmmEnumeration(),
		errors:          make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
func (i *PBmmEnumerationBuilder) SetItemNames(v []string) *PBmmEnumerationBuilder {
	i.pbmmenumeration.itemNames = v
	return i
}
func (i *PBmmEnumerationBuilder) SetItemValues(v []any) *PBmmEnumerationBuilder {
	i.pbmmenumeration.itemValues = v
	return i
}

/*
*
List of immediate inheritance parents. If there are generic ancestors, use
ancestor_defs instead. Persisted attribute.
*/
func (i *PBmmEnumerationBuilder) SetAncestors(v []string) *PBmmEnumerationBuilder {
	i.AddError(i.pbmmenumeration.SetAncestors(v))
	return i
}

// List of attributes defined in this class. Persistent attribute.
func (i *PBmmEnumerationBuilder) SetProperties(v map[string]IPBmmProperty) *PBmmEnumerationBuilder {
	i.AddError(i.pbmmenumeration.SetProperties(v))
	return i
}

// True if this is an abstract type. Persisted attribute.
func (i *PBmmEnumerationBuilder) SetIsAbstract(v bool) *PBmmEnumerationBuilder {
	i.AddError(i.pbmmenumeration.SetIsAbstract(v))
	return i
}

// True if this class definition overrides one found in an included schema.
func (i *PBmmEnumerationBuilder) SetIsOverride(v bool) *PBmmEnumerationBuilder {
	i.AddError(i.pbmmenumeration.SetIsOverride(v))
	return i
}

// List of generic parameter definitions. Persisted attribute.
func (i *PBmmEnumerationBuilder) SetGenericParameterDefs(v map[string]IPBmmGenericParameter) *PBmmEnumerationBuilder {
	i.AddError(i.pbmmenumeration.SetGenericParameterDefs(v))
	return i
}

/*
*
Reference to original source schema defining this class. Set during BMM_SCHEMA
materialise. Useful for GUI tools to enable user to edit the schema file
containing a given class (i.e. taking into account that a class may be in any of
the schemas in a schema inclusion hierarchy).
*/
func (i *PBmmEnumerationBuilder) SetSourceSchemaId(v string) *PBmmEnumerationBuilder {
	i.AddError(i.pbmmenumeration.SetSourceSchemaId(v))
	return i
}

/*
*
BMM_CLASS object built by create_bmm_class_definition and
populate_bmm_class_definition .
*/
func (i *PBmmEnumerationBuilder) SetBmmClass(v IBmmClass) *PBmmEnumerationBuilder {
	i.AddError(i.pbmmenumeration.SetBmmClass(v))
	return i
}

/*
*
Unique id generated for later comparison during merging, in order to detect if
two classes are the same. Assigned in post-load processing.
*/
func (i *PBmmEnumerationBuilder) SetUid(v int) *PBmmEnumerationBuilder {
	i.AddError(i.pbmmenumeration.SetUid(v))
	return i
}

/*
*
List of structured inheritance ancestors, used only in the case of generic
inheritance. Persisted attribute.
*/
//func (i *PBmmClassBuilder) SetAncestorDefs(v []IPBmmGenericType) *PBmmClassBuilder {
//	i.pbmmclass.AncestorDefs = v
//	return i
//}

// //From: PBmmModelElement
// Optional documentation of this element.
func (i *PBmmEnumerationBuilder) SetDocumentation(v string) *PBmmEnumerationBuilder {
	i.AddError(i.pbmmenumeration.SetDocumentation(v))
	return i
}

func (i *PBmmEnumerationBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmEnumerationBuilder) Build() *PBmmEnumeration {
	return i.pbmmenumeration
}

//FUNCTIONS
// From: P_BMM_CLASS
/**
Post : result := generic_parameter_defs /= Void. True if this class is a generic
class.
*/
func (p *PBmmEnumeration) IsGeneric() bool {
	return false
}
