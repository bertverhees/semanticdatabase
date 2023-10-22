package vocabulary

/**
	Abstract idea of specifying a type in some context. This is not the same as
	'defining' a class. A type specification is essentially a reference of some
	kind, that defines the type of an attribute, or function result or argument. It
	may include generic parameters that might or might not be bound. See subtypes.
*/

type IBmmType interface {
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  Boolean
	IsPrimitive (  )  Boolean
	UnitaryType (  )  IBmmUnitaryType
	EffectiveType (  )  IBmmEffectiveType
	FlattenedTypeList (  )  List <String>
}

type BmmType struct {
}
