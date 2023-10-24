package vocabulary

import (
	"vocabulary"
)

/**
	A formal element with signature of the form: name ({arg:TArg}*):TResult . A
	function is a computed (rather than data) element, generally assumed to be
	non-state-changing.
*/

// Interface definition
type IBmmFunction interface {
	// From: BMM_ROUTINE
	Arity (  )  int
	// From: BMM_FEATURE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
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
	OperatorDefinition	IBmmOperator	`yaml:"operatordefinition" json:"operatordefinition" xml:"operatordefinition"`
	// Automatically created Result variable, usable in body and post-condition.
	Result	IBmmResult	`yaml:"result" json:"result" xml:"result"`
}

//CONSTRUCTOR
func NewBmmFunction() *BmmFunction {
	bmmfunction := new(BmmFunction)
	// Constants
	// From: BmmRoutine
	// From: BmmFeature
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmfunction
}
//BUILDER
type BmmFunctionBuilder struct {
	bmmfunction *BmmFunction
}

func NewBmmFunctionBuilder() *BmmFunctionBuilder {
	 return &BmmFunctionBuilder {
		bmmfunction : NewBmmFunction(),
	}
}

//BUILDER ATTRIBUTES
	/**
		Optional details enabling a function to be represented as an operator in a
		syntactic representation.
	*/
func (i *BmmFunctionBuilder) SetOperatorDefinition ( v IBmmOperator ) *BmmFunctionBuilder{
	i.bmmfunction.OperatorDefinition = v
	return i
}
	// Automatically created Result variable, usable in body and post-condition.
func (i *BmmFunctionBuilder) SetResult ( v IBmmResult ) *BmmFunctionBuilder{
	i.bmmfunction.Result = v
	return i
}

func (i *BmmFunctionBuilder) Build() *BmmFunction {
	 return i.bmmfunction
}

//FUNCTIONS
// From: BMM_ROUTINE
// Return number of arguments of this routine.
func (b *BmmFunction) Arity (  )  int {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Formal signature of this element, in the form: name [arg1_name: T_arg1,
	…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmFunction) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmFunction) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmFunction) IsRootScope (  )  bool {
	return nil
}
