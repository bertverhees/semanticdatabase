package vocabulary

/**
	Parent of built-in types, which are treated as being primitive and non-abstract.
*/


type IBmmBuiltinType interface {
	IsAbstract():Boolean (  )  Boolean
	IsPrimitive():Boolean (  )  Boolean
	TypeBaseName():String (  )  string
	TypeName():String (  )  string
}

type BmmBuiltinType struct {
	// Base name (built-in typename).
	BaseName	string	`yaml:"base_name" json:"base_name" xml:"base_name"`
}
