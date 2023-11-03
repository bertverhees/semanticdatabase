package vocabulary

/**
A formal element with signature of the form: name ({arg:TArg}*):TResult . A
function is a computed (rather than data) element, generally assumed to be
non-state-changing.
*/

// Interface definition
type IBmmFunction interface {
	// From: BMM_ROUTINE
	Arity() int
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
}

// Struct definition
type BmmFunction struct {
	// embedded for Inheritance
	BmmRoutine
	BmmFeature
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	/**
	Optional details enabling a function to be represented as an operator in a
	syntactic representation.
	*/
	OperatorDefinition IBmmOperator `yaml:"operatordefinition" json:"operatordefinition" xml:"operatordefinition"`
	// Automatically created Result variable, usable in body and post-condition.
	Result IBmmResult `yaml:"result" json:"result" xml:"result"`
}

// CONSTRUCTOR
func NewBmmFunction() *BmmFunction {
	bmmfunction := new(BmmFunction)
	return bmmfunction
}

// BUILDER
type BmmFunctionBuilder struct {
	bmmfunction *BmmFunction
}

func NewBmmFunctionBuilder() *BmmFunctionBuilder {
	return &BmmFunctionBuilder{
		bmmfunction: NewBmmFunction(),
	}
}

//BUILDER ATTRIBUTES
/**
Optional details enabling a function to be represented as an operator in a
syntactic representation.
*/
func (i *BmmFunctionBuilder) SetOperatorDefinition(v IBmmOperator) *BmmFunctionBuilder {
	i.bmmfunction.OperatorDefinition = v
	return i
}

// Automatically created Result variable, usable in body and post-condition.
func (i *BmmFunctionBuilder) SetResult(v IBmmResult) *BmmFunctionBuilder {
	i.bmmfunction.Result = v
	return i
}

// From: BmmRoutine
// Formal parameters of the routine.
func (i *BmmFunctionBuilder) SetParameters(v []IBmmParameter) *BmmFunctionBuilder {
	i.bmmfunction.Parameters = v
	return i
}

// From: BmmRoutine
/**
Boolean conditions that must evaluate to True for the routine to execute
correctly, May be used to generate exceptions if included in run-time build. A
False pre-condition implies an error in the passed parameters.
*/
func (i *BmmFunctionBuilder) SetPreConditions(v []IBmmAssertion) *BmmFunctionBuilder {
	i.bmmfunction.BmmRoutine.PreConditions = v
	return i
}

// From: BmmRoutine
/**
Boolean conditions that will evaluate to True if the routine executed correctly,
May be used to generate exceptions if included in run-time build. A False
post-condition implies an error (i.e. bug) in routine code.
*/
func (i *BmmFunctionBuilder) SetPostConditions(v []IBmmAssertion) *BmmFunctionBuilder {
	i.bmmfunction.BmmRoutine.PostConditions = v
	return i
}

// From: BmmRoutine
// Body of a routine, i.e. executable program.
func (i *BmmFunctionBuilder) SetDefinition(v IBmmRoutineDefinition) *BmmFunctionBuilder {
	i.bmmfunction.BmmRoutine.Definition = v
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmFunctionBuilder) SetIsSynthesisedGeneric(v bool) *BmmFunctionBuilder {
	i.bmmfunction.IsSynthesisedGeneric = v
	return i
}

// From: BmmFeature
// Extensions to feature-level meta-types.
func (i *BmmFunctionBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmFunctionBuilder {
	i.bmmfunction.FeatureExtensions = v
	return i
}

// From: BmmFeature
// Group containing this feature.
func (i *BmmFunctionBuilder) SetGroup(v IBmmFeatureGroup) *BmmFunctionBuilder {
	i.bmmfunction.Group = v
	return i
}

// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmFunctionBuilder) SetScope(v IBmmClass) *BmmFunctionBuilder {
	i.bmmfunction.BmmModelElement.Scope = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmFunctionBuilder) SetType(v IBmmType) *BmmFunctionBuilder {
	i.bmmfunction.Type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmFunctionBuilder) SetIsNullable(v bool) *BmmFunctionBuilder {
	i.bmmfunction.IsNullable = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmFunctionBuilder) SetName(v string) *BmmFunctionBuilder {
	i.bmmfunction.Name = v
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
	i.bmmfunction.Documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmFunctionBuilder) SetExtensions(v map[string]any) *BmmFunctionBuilder {
	i.bmmfunction.Extensions = v
	return i
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
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmFunction) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmFunction) IsRootScope() bool {
	return false
}
