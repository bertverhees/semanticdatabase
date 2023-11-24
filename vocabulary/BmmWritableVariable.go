package vocabulary

// Meta-type for writable variables, including the special variable result .

// Interface definition
type IBmmWritableVariable interface {
	IBmmVariable
	//BMM_WRITEABLE_VARIABLE
}

// Struct definition
type BmmWritableVariable struct {
	BmmVariable
	// Attributes
}

// CONSTRUCTOR
// abstract, no constructor, no builder

//FUNCTIONS
