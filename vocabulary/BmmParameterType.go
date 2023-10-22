package vocabulary

// Definition of a generic parameter in a class definition of a generic type.

type IBmmParameterType interface {
	FlattenedConformsToType (  )  IBmmEffectiveType
	TypeSignature (  )  string
	IsPrimitive (  )  Boolean
	IsAbstract (  )  Boolean
	TypeName (  )  string
	FlattenedTypeList (  )  List <String>
	EffectiveType (  )  IBmmEffectiveType
}

type BmmParameterType struct {
}
