package vocabulary

// A formal element having a name, type and a type-based signature.

// Interface definition
type IBmmFormalElement interface {
	Signature() IBmmSignature
	IsBoolean() bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
}

// Struct definition
type BmmFormalElement struct {
	// embedded for Inheritance
	BmmModelElement
	// Constants
	// Attributes
	// Declared or inferred static type of the entity.
	Type IBmmType `yaml:"type" json:"type" xml:"type"`
	/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
	*/
	IsNullable bool `yaml:"isnullable" json:"isnullable" xml:"isnullable"`
}

// CONSTRUCTOR
func NewBmmFormalElement() *BmmFormalElement {
	bmmformalelement := new(BmmFormalElement)
	// Constants
	return bmmformalelement
}

// BUILDER
type BmmFormalElementBuilder struct {
	bmmformalelement *BmmFormalElement
}

func NewBmmFormalElementBuilder() *BmmFormalElementBuilder {
	return &BmmFormalElementBuilder{
		bmmformalelement: NewBmmFormalElement(),
	}
}

// BUILDER ATTRIBUTES
// Declared or inferred static type of the entity.
func (i *BmmFormalElementBuilder) SetType(v IBmmType) *BmmFormalElementBuilder {
	i.bmmformalelement.Type = v
	return i
}

/*
*
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmFormalElementBuilder) SetIsNullable(v bool) *BmmFormalElementBuilder {
	i.bmmformalelement.IsNullable = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmFormalElementBuilder) SetName(v string) *BmmFormalElementBuilder {
	i.bmmformalelement.Name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmFormalElementBuilder) SetDocumentation(v map[string]any) *BmmFormalElementBuilder {
	i.bmmformalelement.Documentation = v
	return i
}

// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmFormalElementBuilder) SetScope(v IBmmModelElement) *BmmFormalElementBuilder {
	i.bmmformalelement.Scope = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmFormalElementBuilder) SetExtensions(v map[string]any) *BmmFormalElementBuilder {
	i.bmmformalelement.Extensions = v
	return i
}

func (i *BmmFormalElementBuilder) Build() *BmmFormalElement {
	return i.bmmformalelement
}

//FUNCTIONS
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmFormalElement) Signature() IBmmSignature {
	return nil
}

/*
*
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmFormalElement) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmFormalElement) IsRootScope() bool {
	return false
}
