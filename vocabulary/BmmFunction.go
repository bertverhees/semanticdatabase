package vocabulary

/**
A formal element with signature of the form: name ({arg:TArg}*):TResult . A
function is a computed (rather than data) element, generally assumed to be
non-state-changing.
*/

// Interface definition
type IBmmFunction interface {
	IBmmRoutine
	OperatorDefinition() IBmmOperator
	SetOperatorDefinition(operatorDefinition IBmmOperator) error
	Result() IBmmResult
	SetResult(result IBmmResult) error
}

// Struct definition
type BmmFunction struct {
	BmmRoutine
	// Attributes
	/**
	Optional details enabling a function to be represented as an operator in a
	syntactic representation.
	*/
	operatorDefinition IBmmOperator `yaml:"operatordefinition" json:"operatordefinition" xml:"operatordefinition"`
	// Automatically created result variable, usable in body and post-condition.
	result IBmmResult `yaml:"result" json:"result" xml:"result"`
}

func (b *BmmFunction) OperatorDefinition() IBmmOperator {
	return b.operatorDefinition
}

func (b *BmmFunction) SetOperatorDefinition(operatorDefinition IBmmOperator) error {
	b.operatorDefinition = operatorDefinition
	return nil
}

func (b *BmmFunction) Result() IBmmResult {
	return b.result
}

func (b *BmmFunction) SetResult(result IBmmResult) error {
	b.result = result
	return nil
}

// CONSTRUCTOR
func NewBmmFunction() *BmmFunction {
	bmmfunction := new(BmmFunction)
	//BmmFormalElement
	//default, no constant
	bmmfunction.isNullable = false
	//BmmModelElement
	bmmfunction.documentation = make(map[string]any)
	bmmfunction.extensions = make(map[string]any)
	//BmmFeature
	bmmfunction.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmRoutine
	bmmfunction.parameters = make([]IBmmParameter, 0)
	bmmfunction.preConditions = make([]IBmmAssertion, 0)
	bmmfunction.postConditions = make([]IBmmAssertion, 0)

	return bmmfunction
}

// BUILDER
type BmmFunctionBuilder struct {
	bmmfunction *BmmFunction
	errors      []error
}

func NewBmmFunctionBuilder() *BmmFunctionBuilder {
	return &BmmFunctionBuilder{
		bmmfunction: NewBmmFunction(),
		errors:      make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
Optional details enabling a function to be represented as an operator in a
syntactic representation.
*/
func (i *BmmFunctionBuilder) SetOperatorDefinition(v IBmmOperator) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetOperatorDefinition(v))
	return i
}

// Automatically created result variable, usable in body and post-condition.
func (i *BmmFunctionBuilder) SetResult(v IBmmResult) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetResult(v))
	return i
}

// From: BmmRoutine
// Formal parameters of the routine.
func (i *BmmFunctionBuilder) SetParameters(v []IBmmParameter) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetParameters(v))
	return i
}

// From: BmmRoutine
/**
Boolean conditions that must evaluate to True for the routine to execute
correctly, May be used to generate exceptions if included in run-time build. A
False pre-condition implies an error in the passed parameters.
*/
func (i *BmmFunctionBuilder) SetPreConditions(v []IBmmAssertion) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetPreConditions(v))
	return i
}

// From: BmmRoutine
/**
Boolean conditions that will evaluate to True if the routine executed correctly,
May be used to generate exceptions if included in run-time build. A False
post-condition implies an error (i.e. bug) in routine code.
*/
func (i *BmmFunctionBuilder) SetPostConditions(v []IBmmAssertion) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetPostConditions(v))
	return i
}

// From: BmmRoutine
// body of a routine, i.e. executable program.
func (i *BmmFunctionBuilder) SetDefinition(v IBmmRoutineDefinition) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetDefinition(v))
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmFunctionBuilder) SetIsSynthesisedGeneric(v bool) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetIsSynthesisedGeneric(v))
	return i
}

// From: BmmFeature
// extensions to feature-level meta-types.
func (i *BmmFunctionBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetFeatureExtensions(v))
	return i
}

// From: BmmFeature
// group containing this feature.
func (i *BmmFunctionBuilder) SetGroup(v IBmmFeatureGroup) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetGroup(v))
	return i
}

// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmFunctionBuilder) SetScope(v IBmmClass) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetScope(v))
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmFunctionBuilder) SetIsNullable(v bool) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetIsNullable(v))
	return i
}

// Declared or inferred static type of the entity.
func (i *BmmFunctionBuilder) SetType(v IBmmType) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetType(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmFunctionBuilder) SetName(v string) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmFunctionBuilder) SetDocumentation(v map[string]any) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmFunctionBuilder) SetExtensions(v map[string]any) *BmmFunctionBuilder {
	i.AddError(i.bmmfunction.SetExtensions(v))
	return i
}

func (i *BmmFunctionBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmFunctionBuilder) Build() *BmmFunction {
	return i.bmmfunction
}

// FUNCTIONS
// From: BMM_ROUTINE
// Return number of arguments of this routine.
func (b *BmmFunction) Arity() int {
	return 0
}

// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmFunction) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmFunction) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmFunction) IsRootScope() bool {
	return false
}
