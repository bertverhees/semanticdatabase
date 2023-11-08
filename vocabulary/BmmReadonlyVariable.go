package vocabulary

/**
Meta-type for writable variables, including routine parameters and the special
variable Self .
*/

// Interface definition
type IBmmReadonlyVariable interface {
	// From: BMM_MODEL_ELEMENT
	IBmmModelElement

	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	//BMM_VARIABLE
	//BMM_READONLY_VARIABLE
}

// Struct definition
type BmmReadonlyVariable struct {
	// embedded for Inheritance
	BmmModelElement
	BmmFormalElement
	BmmVariable
	// Constants
	// Attributes
}

// CONSTRUCTOR
// abstract, no constructor, no builder
//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmReadonlyVariable) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmReadonlyVariable) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmReadonlyVariable) IsRootScope() bool {
	return false
}
