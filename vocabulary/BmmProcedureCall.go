package vocabulary

/**
A call made on a closed procedure agent. The method in BMM via which external
actions are achieved from within a program.
*/

// Interface definition
type IBmmProcedureCall interface {
	IElAgentCall
	IBmmSimpleStatement
	IBmmStatement
	IBmmStatementItem
}

// Struct definition
type BmmProcedureCall struct {
	// embedded for Inheritance
	ElAgentCall
	BmmSimpleStatement
	BmmStatement
	BmmStatementItem
	// Constants
	// Attributes
	// The procedure agent being called.
	Agent IElProcedureAgent `yaml:"agent" json:"agent" xml:"agent"`
}

// CONSTRUCTOR
func NewBmmProcedureCall() *BmmProcedureCall {
	bmmprocedurecall := new(BmmProcedureCall)
	// Constants
	return bmmprocedurecall
}

// BUILDER
type BmmProcedureCallBuilder struct {
	bmmprocedurecall *BmmProcedureCall
}

func NewBmmProcedureCallBuilder() *BmmProcedureCallBuilder {
	return &BmmProcedureCallBuilder{
		bmmprocedurecall: NewBmmProcedureCall(),
	}
}

// BUILDER ATTRIBUTES
// The procedure agent being called.
func (i *BmmProcedureCallBuilder) SetAgent(v IElProcedureAgent) *BmmProcedureCallBuilder {
	i.bmmprocedurecall.Agent = v
	return i
}

func (i *BmmProcedureCallBuilder) Build() *BmmProcedureCall {
	return i.bmmprocedurecall
}

//FUNCTIONS
