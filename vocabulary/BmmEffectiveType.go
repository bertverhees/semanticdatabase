package vocabulary

/**
	Meta-type for a concrete, unitary type that can be used as an actual parameter
	type in a generic type declaration.
*/


type IBmmEffectiveType interface {
	EffectiveType():BmmEffectiveType (  )  BMM_EFFECTIVE_TYPE
	TypeBaseName():String (  )  string
}

type BmmEffectiveType struct {
}
