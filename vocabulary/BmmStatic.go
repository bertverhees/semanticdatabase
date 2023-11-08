package vocabulary

// Meta-type for static (i.e. read-only) properties.

// Interface definition
type IBmmStatic interface {
	// From: BMM_MODEL_ELEMENT
	IBmmModelElement

	// BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	//BMM_FEATURE
	//BMM_INSTANTIABLE_FEATURE
	//BMM_STATIC
}

// Struct definition
type BmmStatic struct {
	// embedded for Inheritance
	BmmModelElement
	BmmFormalElement
	BmmFeature
	BmmInstantiableFeature
	// Attributes
}

// CONSTRUCTOR
//Abstract, no constructor, no builder

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmStatic) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmStatic) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmStatic) IsRootScope() bool {
	return false
}
