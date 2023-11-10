package vocabulary

// Multi-branch conditional statement structure

// Interface definition
type IBmmActionTable interface {
	IBmmStatement
	DecisionTable() (IBmmActionDecisionTable, error)
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

func (b *BmmActionTable) DecisionTable() (IBmmActionDecisionTable, error) {
	return b.decisionTable, nil
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
}

func NewBmmActionTableBuilder() *BmmActionTableBuilder {
	return &BmmActionTableBuilder{
		bmmactiontable: NewBmmActionTable(),
	}
}

//BUILDER ATTRIBUTES
/**
A specialised decision table whose outputs can only be procedure agents. In
execution, the matched agent will be invoked.
*/
func (i *BmmActionTableBuilder) SetDecisionTable(v IBmmActionDecisionTable) (*BmmActionTableBuilder, error) {
	e := i.bmmactiontable.SetDecisionTable(v)
	return i, e
}

func (i *BmmActionTableBuilder) Build() *BmmActionTable {
	return i.bmmactiontable
}

//FUNCTIONS
