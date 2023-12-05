package vocabulary

// Persistent form of BMM_PROPER_TYPE .

// Interface definition
type IPBmmBaseType interface {
	IPBmmType
	ValueConstraint() string
	SetValueConstraint(valueConstraint string) error
}

// Struct definition
type PBmmBaseType struct {
	// embedded for Inheritance
	PBmmType
	// Constants
	// Attributes
	valueConstraint string `yaml:"valueconstraint" json:"valueconstraint" xml:"valueconstraint"`
}

func (P *PBmmBaseType) ValueConstraint() string {
	return P.valueConstraint
}

func (P *PBmmBaseType) SetValueConstraint(valueConstraint string) error {
	P.valueConstraint = valueConstraint
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder
// FUNCTIONS
