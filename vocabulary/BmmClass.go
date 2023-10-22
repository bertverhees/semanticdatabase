package vocabulary

/**
	Meta-type corresponding a class definition in an object model. Inheritance is
	specified by the ancestors attribute, which contains a list of types rather than
	classes. Inheritance is thus understood in BMM as a stated relationship between
	classes. The equivalent relationship between types is conformance. Note unlike
	UML, the name is just the root name, even if the class is generic. Use
	type_name() to obtain the qualified type name.
*/

type IBmmClass interface {
	Type (  )  IBmmModelType
	AllAncestors (  )  List <String>
	AllDescendants (  )  List <String>
	Suppliers (  )  List <String>
	SuppliersNonPrimitive (  )  List <String>
	SupplierClosure (  )  List <String>
	PackagePath (  )  string
	ClassPath (  )  string
	IsPrimitive (  )  Boolean
	IsAbstract (  )  Boolean
	Features (  ) 
	FlatFeatures (  ) 
	FlatProperties (  )  List < BMM_PROPERTY >
}

type BmmClass struct {
}
