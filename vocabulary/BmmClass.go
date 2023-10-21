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
	Type():BmmModelType (  )  BMM_MODEL_TYPE
	AllAncestors():List<string> (  )  List <String>
	AllDescendants():List<string> (  )  List <String>
	Suppliers():List<string> (  )  List <String>
	SuppliersNonPrimitive():List<string> (  )  List <String>
	SupplierClosure():List<string> (  )  List <String>
	PackagePath():String (  )  string
	ClassPath():String (  )  string
	IsPrimitive():Boolean (  )  Boolean
	IsAbstract():Boolean (  )  Boolean
	Features() (  ) 
	FlatFeatures() (  ) 
	FlatProperties():List<BmmProperty> (  )  List < BMM_PROPERTY >
}

type BmmClass struct {
}
