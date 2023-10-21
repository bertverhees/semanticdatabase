package vocabulary

// Definition of a generic parameter in a class definition of a generic type.

type IBmmParameterType interface {
	FlattenedConformsToType():BmmEffectiveType (  )  BMM_EFFECTIVE_TYPE
	TypeSignature():String (  )  string
	IsPrimitive():Boolean (  )  Boolean
	IsAbstract():Boolean (  )  Boolean
	TypeName():String (  )  string
	FlattenedTypeList():List<string> (  )  List <String>
	EffectiveType():BmmEffectiveType (  )  BMM_EFFECTIVE_TYPE
}

type BmmParameterType struct {
}
