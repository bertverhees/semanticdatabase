package vocabulary

/**
	Parent of built-in types, which are treated as being primitive and non-abstract.
*/

type IBmmBuiltinType interface {
	IsAbstract (  )  Boolean
	IsPrimitive (  )  Boolean
	TypeBaseName (  )  string
	TypeName (  )  string
}

type BmmBuiltinType struct {
	// Base name (built-in typename).
	BaseName	string	`yaml:"basename" json:"basename" xml:"basename"`
}
