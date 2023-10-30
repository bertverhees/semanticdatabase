package vocabulary

/**
Meta-type for writable variables, including routine parameters and the special
variable Self .
*/

// Interface definition
type IBmmReadonlyVariable interface {
	// From: BMM_VARIABLE
	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
}

// Struct definition
type BmmReadonlyVariable struct {
	// embedded for Inheritance
	BmmVariable
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmReadonlyVariable() *BmmReadonlyVariable {
	bmmreadonlyvariable := new(BmmReadonlyVariable)
	// Constants
	return bmmreadonlyvariable
}

// BUILDER
type BmmReadonlyVariableBuilder struct {
	bmmreadonlyvariable *BmmReadonlyVariable
}

func NewBmmReadonlyVariableBuilder() *BmmReadonlyVariableBuilder {
	return &BmmReadonlyVariableBuilder{
		bmmreadonlyvariable: NewBmmReadonlyVariable(),
	}
}

// BUILDER ATTRIBUTES
// From: BmmVariable
// Routine within which variable is defined.
func (i *BmmReadonlyVariableBuilder) SetScope(v IBmmRoutine) *BmmReadonlyVariableBuilder {
	i.bmmreadonlyvariable.BmmModelElement.Scope = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmReadonlyVariableBuilder) SetType(v IBmmType) *BmmReadonlyVariableBuilder {
	i.bmmreadonlyvariable.Type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmReadonlyVariableBuilder) SetIsNullable(v bool) *BmmReadonlyVariableBuilder {
	i.bmmreadonlyvariable.IsNullable = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmReadonlyVariableBuilder) SetName(v string) *BmmReadonlyVariableBuilder {
	i.bmmreadonlyvariable.Name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmReadonlyVariableBuilder) SetDocumentation(v map[string]any) *BmmReadonlyVariableBuilder {
	i.bmmreadonlyvariable.Documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmReadonlyVariableBuilder) SetExtensions(v map[string]any) *BmmReadonlyVariableBuilder {
	i.bmmreadonlyvariable.Extensions = v
	return i
}

func (i *BmmReadonlyVariableBuilder) Build() *BmmReadonlyVariable {
	return i.bmmreadonlyvariable
}

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
