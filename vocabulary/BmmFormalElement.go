package vocabulary

// A formal element having a name, type and a type-based signature.

// Interface definition
type IBmmFormalElement interface {
	IBmmModelElement
	// BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
}

// Struct definition
type BmmFormalElement struct {
	BmmModelElement
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
