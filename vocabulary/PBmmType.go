package vocabulary

// Persistent form of BMM_TYPE .

// Interface definition
type IPBmmType interface {
	//P_BMM_TYPE
	CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass)
	AsTypeString() string
	BmmType() IBmmType
	SetBmmType(bmmType IBmmType) error
}

// Struct definition
type PBmmType struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// result of create_bmm_type() call.
	bmmType IBmmType `yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

func (p *PBmmType) BmmType() IBmmType {
	return p.bmmType
}

func (p *PBmmType) SetBmmType(bmmType IBmmType) error {
	p.bmmType = bmmType
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder
// FUNCTIONS
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmType) CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass) {
	return
}

// Formal name of the type for display.
func (p *PBmmType) AsTypeString() string {
	return ""
}
