package vocabulary

// Type reference to a single type i.e. not generic or container type.

type IBmmSimpleType interface {
	TypeName():String (  )  string
	IsAbstract():Boolean (  )  Boolean
	FlattenedTypeList():List<string> (  )  List <String>
	EffectiveBaseClass():BmmSimpleClass (  )  BMM_SIMPLE_CLASS
}

type BmmSimpleType struct {
}
