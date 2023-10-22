package vocabulary

// Meta-type based on a non-container generic class, e.g. Packet<Header> .

type IBmmGenericType interface {
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  Boolean
	FlattenedTypeList (  )  List <String>
	IsPartiallyClosed (  )  Boolean
	EffectiveBaseClass (  )  IBmmGenericClass
	IsOpen (  )  Boolean
}

type BmmGenericType struct {
}
