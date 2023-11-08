package vocabulary

/**
Meta-type for an automatically created variable referencing the current
instance. Typically called 'self' or 'this' in programming languages. Read-only.
*/

// Interface definition
type IBmmSelf interface {
	// From: BMM_MODEL_ELEMENT
	IBmmModelElement

	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	//BMM_VARIABLE
	//BMM_READONLY_VARIABLE
	//BMM_SELF
}

// Struct definition
type BmmSelf struct {
	// embedded for Inheritance
	BmmModelElement
	BmmFormalElement
	BmmVariable
	BmmReadonlyVariable
	// Constants
	// Attributes
	// Name of this model element.
	Name string `yaml:"name" json:"name" xml:"name"`
}

// CONSTRUCTOR
func NewBmmSelf() *BmmSelf {
	bmmself := new(BmmSelf)
	//BmmModelElement
	bmmself.documentation = make(map[string]any)
	bmmself.extensions = make(map[string]any)
	//BmmFormalElement
	//default, no constant
	bmmself.IsNullable = false
	//default, no constant
	bmmself.Name = "Self"
	return bmmself
}

// BUILDER
type BmmSelfBuilder struct {
	bmmself *BmmSelf
}

func NewBmmSelfBuilder() *BmmSelfBuilder {
	return &BmmSelfBuilder{
		bmmself: NewBmmSelf(),
	}
}

// BUILDER ATTRIBUTES
// name of this model element.
func (i *BmmSelfBuilder) SetName(v string) *BmmSelfBuilder {
	i.bmmself.Name = v
	return i
}

// From: BmmVariable
// Routine within which variable is defined.
func (i *BmmSelfBuilder) SetScope(v IBmmRoutine) *BmmSelfBuilder {
	i.bmmself.BmmModelElement.scope = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmSelfBuilder) SetType(v IBmmType) *BmmSelfBuilder {
	i.bmmself.Type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmSelfBuilder) SetIsNullable(v bool) *BmmSelfBuilder {
	i.bmmself.IsNullable = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmSelfBuilder) SetDocumentation(v map[string]any) *BmmSelfBuilder {
	i.bmmself.documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmSelfBuilder) SetExtensions(v map[string]any) *BmmSelfBuilder {
	i.bmmself.extensions = v
	return i
}

func (i *BmmSelfBuilder) Build() *BmmSelf {
	return i.bmmself
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmSelf) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmSelf) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmSelf) IsRootScope() bool {
	return false
}
