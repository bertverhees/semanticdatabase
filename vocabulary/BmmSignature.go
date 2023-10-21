package vocabulary

/**
	Built-in meta-type that expresses the type structure of any referenceable
	element of a model. Consists of potential arguments and result , with
	constraints in descendants determining the exact form.
*/

type IBmmSignature interface {
	/**
		Return the logical set (i.e. unique items) consisting of
		argument_types.flattened_type_list () and result_type.flattened_type_list () .
	*/
	flattened_type_list (): List <String> (  )  List <String>
}

type BmmSignature struct {
}
