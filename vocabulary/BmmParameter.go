package vocabulary

// A routine parameter variable (read-only).

// Interface definition
type IBmmParameter interface {
	IBmmReadonlyVariable
	//BMM_PARAMETER
	Direction() BmmParameterDirection
	SetDirection(direction BmmParameterDirection) error
}

// Struct definition
type BmmParameter struct {
	// embedded for Inheritance
	BmmReadonlyVariable
	// Constants
	// Attributes
	/**
	Optional read/write direction of the parameter. If none-supplied, the parameter
	is treated as in , i.e. readable.
	*/
	direction BmmParameterDirection `yaml:"direction" json:"direction" xml:"direction"`
}

func (b *BmmParameter) Direction() BmmParameterDirection {
	return b.direction
}

func (b *BmmParameter) SetDirection(direction BmmParameterDirection) error {
	b.direction = direction
	return nil
}

// CONSTRUCTOR
func NewBmmParameter() *BmmParameter {
	bmmparameter := new(BmmParameter)
	//BmmModelElement
	bmmparameter.documentation = make(map[string]any)
	bmmparameter.extensions = make(map[string]any)
	//BmmFormalElement
	//default, no constant
	bmmparameter.isNullable = false
	return bmmparameter
}

// BUILDER
type BmmParameterBuilder struct {
	bmmparameter *BmmParameter
	errors       []error
}

func NewBmmParameterBuilder() *BmmParameterBuilder {
	return &BmmParameterBuilder{
		bmmparameter: NewBmmParameter(),
		errors:       make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
Optional read/write direction of the parameter. If none-supplied, the parameter
is treated as in , i.e. readable.
*/
func (i *BmmParameterBuilder) SetDirection(v BmmParameterDirection) *BmmParameterBuilder {
	i.AddError(i.bmmparameter.SetDirection(v))
	return i
}

// From: BmmVariable
// Routine within which variable is defined.
func (i *BmmParameterBuilder) SetScope(v IBmmRoutine) *BmmParameterBuilder {
	i.AddError(i.bmmparameter.SetScope(v))
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmParameterBuilder) SetIsNullable(v bool) *BmmParameterBuilder {
	i.AddError(i.bmmparameter.SetIsNullable(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmParameterBuilder) SetName(v string) *BmmParameterBuilder {
	i.AddError(i.bmmparameter.SetName(v))
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
	i.AddError(i.bmmparameter.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmParameterBuilder) SetExtensions(v map[string]any) *BmmParameterBuilder {
	i.AddError(i.bmmparameter.SetExtensions(v))
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmParameterBuilder) SetType(v IBmmType) *BmmParameterBuilder {
	i.AddError(i.bmmparameter.SetType(v))
	return i
}

func (i *BmmParameterBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
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
Post_result : result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmParameter) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmParameter) IsRootScope() bool {
	return false
}
