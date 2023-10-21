package vocabulary

// A type that is defined by a class (or classes) in the model.

type IBmmModelType interface {
// Result = base_class.name .
	type_base_name (): String (  )  String
// Result = base_class.is_primitive .
	is_primitive (): Boolean (  )  Boolean
}

type BmmModelType struct {
}
