package vocabulary

import "errors"

// A routine-scoped formal element.

// Interface definition
type IBmmVariable interface {
	IBmmFormalElement
	//BMM_VARIABLE
}

// Struct definition
type BmmVariable struct {
	BmmFormalElement
	// Attributes
	// Routine within which variable is defined.
	scope IBmmRoutine `yaml:"scope" json:"scope" xml:"scope"`
}

func (b *BmmVariable) SetScope(scope IBmmModelElement) error {
	s, ok := scope.(IBmmRoutine)
	if !ok {
		return errors.New("_type-assertion in BmmVariable->SetScope went wrong")
	} else {
		b.scope = s
		return nil
	}
}

// CONSTRUCTOR
// abstract, no constructor, no builder

//FUNCTIONS
