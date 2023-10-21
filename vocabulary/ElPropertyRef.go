package vocabulary

// Reference to a writable property.

type IElPropertyRef interface {
	/**
		Type definition (i.e. BMM meta-type definition object) of the constant, property
		or variable, inferred by inspection of the current scoping instance. Return
		definition.type .
	*/
	eval_type (): BMM_TYPE (  )  BMM_TYPE
}

type ElPropertyRef struct {
}
