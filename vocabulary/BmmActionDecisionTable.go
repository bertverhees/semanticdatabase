package vocabulary

/**
Specialised form of Decision Table that allows only procedure call agents
(lambdas) as the result of branches.
*/

// Interface definition
type IBmmActionDecisionTable interface {
}

// Struct definition
type BmmActionDecisionTable struct {
	// embedded for Inheritance
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmActionDecisionTable() *BmmActionDecisionTable {
	bmmactiondecisiontable := new(BmmActionDecisionTable)
	// Constants
	return bmmactiondecisiontable
}

// BUILDER
type BmmActionDecisionTableBuilder struct {
	bmmactiondecisiontable *BmmActionDecisionTable
	errors                 []error
}

func NewBmmActionDecisionTableBuilder() *BmmActionDecisionTableBuilder {
	return &BmmActionDecisionTableBuilder{
		bmmactiondecisiontable: NewBmmActionDecisionTable(),
		errors:                 make([]error, 0),
	}
}

//BUILDER ATTRIBUTES

func (i *BmmActionDecisionTableBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmActionDecisionTableBuilder) Build() *BmmActionDecisionTable {
	return i.bmmactiondecisiontable
}

//FUNCTIONS
