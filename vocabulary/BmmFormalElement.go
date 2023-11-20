package vocabulary

// A formal element having a name, type and a type-based signature.

// Interface definition
type IBmmFormalElement interface {
	IBmmModelElement
	// BMM_FORMAL_ELEMENT
	Type() IBmmType
	SetType(_type IBmmType) error
	IsNullable() bool
	SetIsNullable(isNullable bool) error
	Signature() IBmmSignature
	IsBoolean() bool
}

// Struct definition
type BmmFormalElement struct {
	BmmModelElement
	// Attributes
	// Declared or inferred static type of the entity.
	_type IBmmType `yaml:"type" json:"type" xml:"type"`
	/**
	True if this element can be null (Void) at execution time. May be interpreted as
	optionality in subtypes..
	*/
	isNullable bool `yaml:"isnullable" json:"isnullable" xml:"isnullable"`
}

func (b *BmmFormalElement) Type() IBmmType {
	return b._type
}

func (b *BmmFormalElement) SetType(_type IBmmType) error {
	b._type = _type
	return nil
}

func (b *BmmFormalElement) IsNullable() bool {
	return b.isNullable
}

func (b *BmmFormalElement) SetIsNullable(isNullable bool) error {
	b.isNullable = isNullable
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder

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
