package vocabulary

// Meta-type based on a non-container generic class, e.g. Packet<Header> .

type IBmmGenericType interface {
	TypeName():String (  )  string
	TypeSignature():String (  )  string
	IsAbstract():Boolean (  )  Boolean
	FlattenedTypeList():List<string> (  )  List <String>
	IsPartiallyClosed():Boolean (  )  Boolean
	EffectiveBaseClass():BmmGenericClass (  )  BMM_GENERIC_CLASS
	IsOpen():Boolean (  )  Boolean
}

type BmmGenericType struct {
}
