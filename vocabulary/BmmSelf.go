package vocabulary

/**
Meta-type for an automatically created variable referencing the current
instance. Typically called 'self' or 'this' in programming languages. Read-only.
*/

// Interface definition
type IBmmSelf interface {
	//BMM_READONLY_VARIABLE
	IBmmReadonlyVariable
	//BMM_SELF
}

// Struct definition
type BmmSelf struct {
	BmmReadonlyVariable
	// Constants
	// Attributes
	// name of this model element.
}

// CONSTRUCTOR
func NewBmmSelf() *BmmSelf {
	bmmself := new(BmmSelf)
	//BmmModelElement
	bmmself.documentation = make(map[string]any)
	bmmself.extensions = make(map[string]any)
	//BmmFormalElement
	//default, no constant
	bmmself.isNullable = false
	//default, no constant
	bmmself.name = "Self"
	return bmmself
}

// BUILDER
type BmmSelfBuilder struct {
	bmmself *BmmSelf
	errors  []error
}

func NewBmmSelfBuilder() *BmmSelfBuilder {
	return &BmmSelfBuilder{
		bmmself: NewBmmSelf(),
		errors:  make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// From: BmmVariable
// Routine within which variable is defined.
func (i *BmmSelfBuilder) SetScope(v IBmmRoutine) *BmmSelfBuilder {
	i.AddError(i.bmmself.SetScope(v))
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmSelfBuilder) SetIsNullable(v bool) *BmmSelfBuilder {
	i.AddError(i.bmmself.SetIsNullable(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmSelfBuilder) SetName(v string) *BmmSelfBuilder {
	i.AddError(i.bmmself.SetName(v))
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
	i.AddError(i.bmmself.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmSelfBuilder) SetExtensions(v map[string]any) *BmmSelfBuilder {
	i.AddError(i.bmmself.SetExtensions(v))
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmSelfBuilder) SetType(v IBmmType) *BmmSelfBuilder {
	i.AddError(i.bmmself.SetType(v))
	return i
}

func (i *BmmSelfBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
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
Post_result : result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmSelf) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmSelf) IsRootScope() bool {
	return false
}
