package vocabulary

import (
	"vocabulary"
)

// A routine parameter variable (read-only).

// Interface definition
type IBmmParameter interface {
	// From: BMM_READONLY_VARIABLE
	// From: BMM_VARIABLE
	// From: BMM_FORMAL_ELEMENT
	Signature (  )  IBmmSignature
	IsBoolean (  )  bool
	// From: BMM_MODEL_ELEMENT
	IsRootScope (  )  bool
}

// Struct definition
type BmmParameter struct {
	// embedded for Inheritance
	BmmReadonlyVariable
	BmmVariable
	BmmFormalElement
	BmmModelElement
	// Constants
	// Attributes
	/**
		Optional read/write direction of the parameter. If none-supplied, the parameter
		is treated as in , i.e. readable.
	*/
	Direction	IBmmParameterDirection	`yaml:"direction" json:"direction" xml:"direction"`
}

//CONSTRUCTOR
func NewBmmParameter() *BmmParameter {
	bmmparameter := new(BmmParameter)
	// Constants
	// From: BmmReadonlyVariable
	// From: BmmVariable
	// From: BmmFormalElement
	// From: BmmModelElement
	return bmmparameter
}
//BUILDER
type BmmParameterBuilder struct {
	bmmparameter *BmmParameter
}

func NewBmmParameterBuilder() *BmmParameterBuilder {
	 return &BmmParameterBuilder {
		bmmparameter : NewBmmParameter(),
	}
}

//BUILDER ATTRIBUTES
	/**
		Optional read/write direction of the parameter. If none-supplied, the parameter
		is treated as in , i.e. readable.
	*/
func (i *BmmParameterBuilder) SetDirection ( v IBmmParameterDirection ) *BmmParameterBuilder{
	i.bmmparameter.Direction = v
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
func (b *BmmParameter) Signature (  )  IBmmSignature {
	return nil
}
// From: BMM_FORMAL_ELEMENT
/**
	Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
	True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
	'Boolean' ).
*/
func (b *BmmParameter) IsBoolean (  )  bool {
	return nil
}
// From: BMM_MODEL_ELEMENT
/**
	Post_result : Result = (scope = self). True if this model element is the root of
	a model structure hierarchy.
*/
func (b *BmmParameter) IsRootScope (  )  bool {
	return nil
}
