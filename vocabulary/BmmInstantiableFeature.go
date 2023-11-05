package vocabulary

/**
Meta-type representing instantiable features, i.e. features that are created as
value objects.
*/

// Interface definition
type IBmmInstantiableFeature interface {
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
	// BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	//BMM_FEATURE
	//BMM_INSTANTIABLE_FEATURE
}

// Struct definition
type BmmInstantiableFeature struct {
	// embedded for Inheritance
	BmmModelElement
	BmmFormalElement
	BmmFeature
}

// CONSTRUCTOR
// abstract, no constructor, no builder

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmInstantiableFeature) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmInstantiableFeature) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmInstantiableFeature) IsRootScope() bool {
	return false
}
