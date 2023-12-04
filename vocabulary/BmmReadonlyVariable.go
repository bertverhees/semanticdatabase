package vocabulary

/**
Meta-type for writable variables, including routine parameters and the special
variable Self .
*/

// Interface definition
type IBmmReadonlyVariable interface {
	IBmmVariable
	//BMM_READONLY_VARIABLE
}

// Struct definition
type BmmReadonlyVariable struct {
	BmmVariable
	// Attributes
}

// CONSTRUCTOR
// abstract, no constructor, no builder
//FUNCTIONS
