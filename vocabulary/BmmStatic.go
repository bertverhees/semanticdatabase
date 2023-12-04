package vocabulary

// Meta-type for static (i.e. read-only) properties.

// Interface definition
type IBmmStatic interface {
	IBmmInstantiableFeature
	//BMM_STATIC
}

// Struct definition
type BmmStatic struct {
	BmmInstantiableFeature
	// Attributes
}

// CONSTRUCTOR
//Abstract, no constructor, no builder

//FUNCTIONS
