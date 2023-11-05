package vocabulary

// A routine parameter variable (read-only).

// Interface definition
type IBmmParameter interface {
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	//BMM_VARIABLE
	//BMM_READONLY_VARIABLE
	//BMM_PARAMETER
}

// Struct definition
type BmmParameter struct {
	// embedded for Inheritance
	BmmModelElement
	BmmFormalElement
	BmmVariable
	BmmReadonlyVariable
	// Constants
	// Attributes
	/**
	Optional read/write direction of the parameter. If none-supplied, the parameter
	is treated as in , i.e. readable.
	*/
	Direction BmmParameterDirection `yaml:"direction" json:"direction" xml:"direction"`
}

// CONSTRUCTOR
func NewBmmParameter() *BmmParameter {
	bmmparameter := new(BmmParameter)
	//BmmModelElement
	bmmparameter.Documentation = make(map[string]any)
	bmmparameter.Extensions = make(map[string]any)
	//BmmFormalElement
	//default, no constant
	bmmparameter.IsNullable = false
	return bmmparameter
}

// BUILDER
type BmmParameterBuilder struct {
	bmmparameter *BmmParameter
}

func NewBmmParameterBuilder() *BmmParameterBuilder {
	return &BmmParameterBuilder{
		bmmparameter: NewBmmParameter(),
	}
}

//BUILDER ATTRIBUTES
/**
Optional read/write direction of the parameter. If none-supplied, the parameter
is treated as in , i.e. readable.
*/
func (i *BmmParameterBuilder) SetDirection(v BmmParameterDirection) *BmmParameterBuilder {
	i.bmmparameter.Direction = v
	return i
}

// From: BmmVariable
// Routine within which variable is defined.
func (i *BmmParameterBuilder) SetScope(v IBmmRoutine) *BmmParameterBuilder {
	i.bmmparameter.BmmModelElement.Scope = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmParameterBuilder) SetType(v IBmmType) *BmmParameterBuilder {
	i.bmmparameter.Type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmParameterBuilder) SetIsNullable(v bool) *BmmParameterBuilder {
	i.bmmparameter.IsNullable = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmParameterBuilder) SetName(v string) *BmmParameterBuilder {
	i.bmmparameter.Name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmParameterBuilder) SetDocumentation(v map[string]any) *BmmParameterBuilder {
	i.bmmparameter.Documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmParameterBuilder) SetExtensions(v map[string]any) *BmmParameterBuilder {
	i.bmmparameter.Extensions = v
	return i
}

func (i *BmmParameterBuilder) Build() *BmmParameter {
	return i.bmmparameter
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmParameter) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmParameter) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmParameter) IsRootScope() bool {
	return false
}
