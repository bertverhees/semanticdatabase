package vocabulary

// Type reference to a single type i.e. not generic or container type.

type IBmmSimpleType interface {
// Result is base_class.name .
	type_name (): String (  )  String
// Result is base_class.is_abstract .
	is_abstract (): Boolean (  )  Boolean
// Result is base_class.name .
	flattened_type_list (): List <String> (  )  List <String>
// Main design class for this type, from which properties etc can be extracted.
	effective_base_class (): BMM_SIMPLE_CLASS (  )  BMM_SIMPLE_CLASS
}

type BmmSimpleType struct {
}
