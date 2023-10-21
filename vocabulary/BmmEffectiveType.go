package vocabulary

/**
	Meta-type for a concrete, unitary type that can be used as an actual parameter
	type in a generic type declaration.
*/

type IBmmEffectiveType interface {
// Result = self.
	effective_type (): BMM_EFFECTIVE_TYPE (  )  BMM_EFFECTIVE_TYPE
// Name of base generator type, i.e. excluding any generic parts if present.
	type_base_name (): String (  )  String
}

type BmmEffectiveType struct {
}
