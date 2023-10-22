package vocabulary

// Type reference to a single type i.e. not generic or container type.

type IBmmSimpleType interface {
	TypeName (  )  string
	IsAbstract (  )  Boolean
	FlattenedTypeList (  )  List <String>
	EffectiveBaseClass (  )  IBmmSimpleClass
}

type BmmSimpleType struct {
}
