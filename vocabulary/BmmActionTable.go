package vocabulary

// Multi-branch conditional statement structure

// Interface definition
type IBmmActionTable interface {
	IBmmStatement
	DecisionTable() IBmmActionDecisionTable
	SetDecisionTable(decisionTable IBmmActionDecisionTable) error
}

// Struct definition
type BmmActionTable struct {
	BmmStatement
	// Attributes
	/**
	A specialised decision table whose outputs can only be procedure agents. In
	execution, the matched agent will be invoked.
	*/
	decisionTable IBmmActionDecisionTable `yaml:"decisiontable" json:"decisiontable" xml:"decisiontable"`
}

func (b *BmmActionTable) DecisionTable() IBmmActionDecisionTable {
	return b.decisionTable
}

func (b *BmmActionTable) SetDecisionTable(decisionTable IBmmActionDecisionTable) error {
	b.decisionTable = decisionTable
	return nil
}

// CONSTRUCTOR
func NewBmmActionTable() *BmmActionTable {
	bmmactiontable := new(BmmActionTable)
	// Constants
	return bmmactiontable
}

// BUILDER
type BmmActionTableBuilder struct {
	bmmactiontable *BmmActionTable
	errors         []error
}

func NewBmmActionTableBuilder() *BmmActionTableBuilder {
	return &BmmActionTableBuilder{
		bmmactiontable: NewBmmActionTable(),
		errors:         make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
A specialised decision table whose outputs can only be procedure agents. In
execution, the matched agent will be invoked.
*/
func (i *BmmActionTableBuilder) SetDecisionTable(v IBmmActionDecisionTable) *BmmActionTableBuilder {
	i.AddError(i.bmmactiontable.SetDecisionTable(v))
	return i
}

func (i *BmmActionTableBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmActionTableBuilder) Build() *BmmActionTable {
	return i.bmmactiontable
}

//FUNCTIONS
