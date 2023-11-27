package vocabulary

/**
A reference that is scoped by a containing entity and requires a context
qualifier if it is not the currently scoping entity.
*/

// Interface definition
type IElFeatureRef interface {
	IElValueGenerator
	//EL_FEATURE_REF
	Reference() string //redefined
	Scoper() IElValueGenerator
	SetScoper(scoper IElValueGenerator) error
}

// Struct definition
type ElFeatureRef struct {
	ElValueGenerator
	// Attributes
	// Scoping expression, which must be a EL_VALUE_GENERATOR .
	scoper IElValueGenerator `yaml:"scoper" json:"scoper" xml:"scoper"`
}

func (e *ElFeatureRef) Scoper() IElValueGenerator {
	return e.scoper
}

func (e *ElFeatureRef) SetScoper(scoper IElValueGenerator) error {
	e.scoper = scoper
	return nil
}

// CONSTRUCTOR
//abstract, no constructor, no builder
//FUNCTIONS
/**
Generated full reference name, consisting of scoping elements and name
concatenated using dot notation.
*/
func (e *ElFeatureRef) Reference() string {
	return ""
}
