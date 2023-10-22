package vocabulary

import (
	"vocabulary"
)

// Persistent form of BMM_TYPE .

type IPBmmType interface {
	CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	AsTypeString (  )  string
}

type PBmmType struct {
	// Result of create_bmm_type() call.
	BmmType	IBmmType	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmType) CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
// Formal name of the type for display.
func (p *PBmmType) AsTypeString (  )  string {
	return nil
}
