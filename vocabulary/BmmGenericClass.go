package vocabulary

// Definition of a generic class in an object model.

type IBmmGenericClass interface {
	Suppliers():List<string> (  )  List <String>
	Type():BmmGenericType (  )  BMM_GENERIC_TYPE
	GenericParameterConformanceType(AName:String[1]):String ( a_name String[1] )  string
}

type BmmGenericClass struct {
}
