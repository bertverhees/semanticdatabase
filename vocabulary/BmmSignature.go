package vocabulary

/**
	Built-in meta-type that expresses the type structure of any referenceable
	element of a model. Consists of potential arguments and result , with
	constraints in descendants determining the exact form.
*/

type IBmmSignature interface {
	FlattenedTypeList():List<string> (  )  List <String>
}

type BmmSignature struct {
	// Base name (built-in).
	BaseName	string	`yaml:"base_name" json:"base_name" xml:"base_name"`
}
