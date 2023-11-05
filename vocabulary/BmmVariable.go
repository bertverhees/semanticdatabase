package vocabulary

// A routine-scoped formal element.

// Interface definition
type IBmmVariable interface {
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	//BMM_VARIABLE
}

// Struct definition
type BmmVariable struct {
	// embedded for Inheritance
	BmmModelElement
	BmmFormalElement
	// Constants
	// Attributes
	// Routine within which variable is defined.
	Scope IBmmRoutine `yaml:"scope" json:"scope" xml:"scope"`
}

// CONSTRUCTOR
// abstract, no constructor, no builder

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmVariable) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmVariable) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmVariable) IsRootScope() bool {
	return false
}
