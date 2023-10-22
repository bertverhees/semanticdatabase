package vocabulary

// Definition of a generic class in an object model.

type IBmmGenericClass interface {
	Suppliers (  )  List <String>
	Type (  )  IBmmGenericType
	GenericParameterConformanceType ( a_name string )  string
}

type BmmGenericClass struct {
}
