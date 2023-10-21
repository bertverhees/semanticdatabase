package vocabulary

/**
	Built-in meta-type representing the type of a tuple, i.e. an array of any number
	of other types. This includes both container and unitary types, since tuple
	instances represent concrete objects. Note that both open and closed generic
	parameters are allowed, as with any generic type, but open generic parameters
	are only valid within the scope of a generic class.
*/


type IBmmTupleType interface {
	FlattenedTypeList():List<string> (  )  List <String>
}

type BmmTupleType struct {
	// Base name (built-in).
	BaseName	string	`yaml:"base_name" json:"base_name" xml:"base_name"`
}
