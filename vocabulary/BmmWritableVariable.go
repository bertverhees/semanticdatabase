package vocabulary

// Meta-type for writable variables, including the special variable Result .

// Interface definition
type IBmmWritableVariable interface {
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
	Name() string
	SetName(name string)
	Documentation() map[string]any
	SetDocumentation(v map[string]any)
	Scope() IBmmModelElement
	SetScope(v IBmmModelElement)
	Extensions() map[string]any
	SetExtensions(v map[string]any)

	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	//BMM_VARIABLE
	//BMM_WRITEABLE_VARIABLE
}

// Struct definition
type BmmWritableVariable struct {
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
func (b *BmmWritableVariable) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmWritableVariable) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmWritableVariable) IsRootScope() bool {
	return false
}
