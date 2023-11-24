package vocabulary

/**
Automatically declared variable representing result of a Function call
(writable).
*/

// Interface definition
type IBmmResult interface {
	IBmmWritableVariable
	//BMM_RESULT
}

// Struct definition
type BmmResult struct {
	BmmWritableVariable
	// Attributes
	// name of this model element.
	name string `yaml:"name" json:"name" xml:"name"`
}

// CONSTRUCTOR
func NewBmmResult() *BmmResult {
	bmmresult := new(BmmResult)
	//BmmModelElement
	bmmresult.documentation = make(map[string]any)
	bmmresult.extensions = make(map[string]any)
	//BmmFormalElement
	//default, no constant
	bmmresult.isNullable = false

	return bmmresult
}

// BUILDER
type BmmResultBuilder struct {
	bmmresult *BmmResult
}

func NewBmmResultBuilder() *BmmResultBuilder {
	return &BmmResultBuilder{
		bmmresult: NewBmmResult(),
	}
}

// BUILDER ATTRIBUTES
// name of this model element.
func (i *BmmResultBuilder) SetName(v string) *BmmResultBuilder {
	i.bmmresult.name = v
	return i
}

// From: BmmVariable
// Routine within which variable is defined.
func (i *BmmResultBuilder) SetScope(v IBmmRoutine) *BmmResultBuilder {
	i.bmmresult.BmmModelElement.scope = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmResultBuilder) SetType(v IBmmType) *BmmResultBuilder {
	i.bmmresult._type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmResultBuilder) SetIsNullable(v bool) *BmmResultBuilder {
	i.bmmresult.isNullable = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmResultBuilder) SetDocumentation(v map[string]any) *BmmResultBuilder {
	i.bmmresult.documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmResultBuilder) SetExtensions(v map[string]any) *BmmResultBuilder {
	i.bmmresult.extensions = v
	return i
}

func (i *BmmResultBuilder) Build() *BmmResult {
	return i.bmmresult
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmResult) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmResult) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmResult) IsRootScope() bool {
	return false
}
