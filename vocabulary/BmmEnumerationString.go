package vocabulary

import "errors"

// String-based enumeration meta-type.

// Interface definition
type IBmmEnumerationString interface {
	IBmmEnumeration
}

// Struct definition
type BmmEnumerationString struct {
	BmmEnumeration
	// Constants
	// Attributes
	// Optional list of specific values. Must be 1:1 with item_names list.
	itemValues []IBmmStringValue `yaml:"itemvalues" json:"itemvalues" xml:"itemvalues"`
}

func (b *BmmEnumerationString) SetItemValues(itemValues []IBmmPrimitiveValue) error {
	b.itemValues = make([]IBmmStringValue, 0)
	for _, s := range itemValues {
		s, ok := s.(IBmmStringValue)
		if !ok {
			return errors.New("_type-assertion in BmmEnumerationString->SetItemValues went wrong")
		} else {
			b.itemValues = append(b.itemValues, s)
		}
	}
	return nil
}

// CONSTRUCTOR
func NewBmmEnumerationString() *BmmEnumerationString {
	bmmenumerationstring := new(BmmEnumerationString)
	//BmmModelElement
	bmmenumerationstring.documentation = make(map[string]any)
	bmmenumerationstring.extensions = make(map[string]any)
	//BmmModule
	bmmenumerationstring.features = make([]IBmmFeature, 0)
	bmmenumerationstring.featureGroups = make([]IBmmFeatureGroup, 0)
	//BmmClass
	bmmenumerationstring.ancestors = make(map[string]IBmmModelType)
	bmmenumerationstring.properties = make(map[string]IBmmProperty)
	bmmenumerationstring.immediateDescendants = make([]IBmmClass, 0)
	bmmenumerationstring.staticProperties = make(map[string]IBmmStatic)
	bmmenumerationstring.functions = make(map[string]IBmmFunction)
	bmmenumerationstring.procedures = make(map[string]IBmmProcedure)
	bmmenumerationstring.invariants = make([]IBmmAssertion, 0)
	bmmenumerationstring.creators = make(map[string]IBmmProcedure)
	bmmenumerationstring.converters = make(map[string]IBmmProcedure)
	//BmmEnumeration
	bmmenumerationstring.itemValues = make([]IBmmStringValue, 0)

	return bmmenumerationstring
}

// BUILDER
type BmmEnumerationStringBuilder struct {
	bmmenumerationstring *BmmEnumerationString
	errors               []error
}

func NewBmmEnumerationStringBuilder() *BmmEnumerationStringBuilder {
	return &BmmEnumerationStringBuilder{
		bmmenumerationstring: NewBmmEnumerationString(),
		errors:               make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Optional list of specific values. Must be 1:1 with item_names list.
func (i *BmmEnumerationStringBuilder) SetItemValues(v []IBmmPrimitiveValue) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetItemValues(v))
	return i
}

// From: BmmEnumeration
/**
The list of names of the enumeration. If no values are supplied, the integer
values 0, 1, 2, …​ are assumed.
*/
func (i *BmmEnumerationStringBuilder) SetItemNames(v []string) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetItemNames(v))
	return i
}

// From: BmmClass
// List of immediate inheritance parents.
func (i *BmmEnumerationStringBuilder) SetAncestors(v map[string]IBmmModelType) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetAncestors(v))
	return i
}

// From: BmmClass
// Package this class belongs to.
func (i *BmmEnumerationStringBuilder) SetPackage(v IBmmPackage) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetPackage(v))
	return i
}

// From: BmmClass
// Properties defined in this class (subset of features ).
func (i *BmmEnumerationStringBuilder) SetProperties(v map[string]IBmmProperty) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetProperties(v))
	return i
}

// From: BmmClass
/**
Reference to original source schema defining this class. Useful for UI tools to
determine which original schema file to open for a given class for manual
editing.
*/
func (i *BmmEnumerationStringBuilder) SetSourceSchemaId(v string) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetSourceSchemaId(v))
	return i
}

// From: BmmClass
/**
List of computed references to base classes of immediate inheritance
descendants, derived when members of ancestors are attached at creation time.
*/
func (i *BmmEnumerationStringBuilder) SetImmediateDescendants(v []IBmmClass) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetImmediateDescendants(v))
	return i
}

// From: BmmClass
/**
True if this definition overrides a class of the same name in an included
schema.
*/
func (i *BmmEnumerationStringBuilder) SetIsOverride(v bool) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetIsOverride(v))
	return i
}

// From: BmmClass
// Static properties defined in this class (subset of features ).
func (i *BmmEnumerationStringBuilder) SetStaticProperties(v map[string]IBmmStatic) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetStaticProperties(v))
	return i
}

// From: BmmClass
// Functions defined in this class (subset of features ).
func (i *BmmEnumerationStringBuilder) SetFunctions(v map[string]IBmmFunction) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetFunctions(v))
	return i
}

// From: BmmClass
// Procedures defined in this class (subset of features ).
func (i *BmmEnumerationStringBuilder) SetProcedures(v map[string]IBmmProcedure) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetProcedures(v))
	return i
}

// From: BmmClass
/**
True if this class represents a type considered to be primitive in the type
system, i.e. any typically built-in or standard library type such as String ,
Date , Hash<K,V> etc.
*/
func (i *BmmEnumerationStringBuilder) SetIsPrimitive(v bool) *BmmEnumerationStringBuilder {
	i.bmmenumerationstring.SetIsPrimitive(v)
	return i
}

// From: BmmClass
/**
True if this class is marked as abstract, i.e. direct instances cannot be
created from its direct type.
*/
func (i *BmmEnumerationStringBuilder) SetIsAbstract(v bool) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetIsAbstract(v))
	return i
}

// From: BmmClass
func (i *BmmEnumerationStringBuilder) SetInvariants(v []IBmmAssertion) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetInvariants(v))
	return i
}

// From: BmmClass
/**
Subset of procedures that may be used to initialise a new instance of an object,
and whose execution will guarantee that class invariants are satisfied.
*/
func (i *BmmEnumerationStringBuilder) SetCreators(v map[string]IBmmProcedure) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetCreators(v))
	return i
}

// From: BmmClass
/**
Subset of creators that create a new instance from a single argument of another
type.
*/
func (i *BmmEnumerationStringBuilder) SetConverters(v map[string]IBmmProcedure) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetConverters(v))
	return i
}

// From: BmmClass
// features of this module.
func (i *BmmEnumerationStringBuilder) SetFeatures(v []IBmmFormalElement) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetFeatures(v))
	return i
}

// From: BmmModule
// List of feature groups in this class.
func (i *BmmEnumerationStringBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetFeatureGroups(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmEnumerationStringBuilder) SetName(v string) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmEnumerationStringBuilder) SetDocumentation(v map[string]any) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetDocumentation(v))
	return i
}

// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmEnumerationStringBuilder) SetScope(v IBmmModelElement) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetScope(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmEnumerationStringBuilder) SetExtensions(v map[string]any) *BmmEnumerationStringBuilder {
	i.AddError(i.bmmenumerationstring.SetExtensions(v))
	return i
}

func (i *BmmEnumerationStringBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmEnumerationStringBuilder) Build() *BmmEnumerationString {
	return i.bmmenumerationstring
}

// FUNCTIONS
