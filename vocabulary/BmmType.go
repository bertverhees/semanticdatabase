package vocabulary

/**
	Abstract idea of specifying a type in some context. This is not the same as
	'defining' a class. A type specification is essentially a reference of some
	kind, that defines the type of an attribute, or function result or argument. It
	may include generic parameters that might or might not be bound. See subtypes.
*/

type IBmmType interface {
	TypeName():String (  )  string
	TypeSignature():String (  )  string
	IsAbstract():Boolean (  )  Boolean
	IsPrimitive():Boolean (  )  Boolean
	UnitaryType():BmmUnitaryType (  )  BMM_UNITARY_TYPE
	EffectiveType():BmmEffectiveType (  )  BMM_EFFECTIVE_TYPE
	FlattenedTypeList():List<string> (  )  List <String>
}

type BmmType struct {
}
