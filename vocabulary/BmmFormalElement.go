package vocabulary

// A formal element having a name, type and a type-based signature.

type IBmmFormalElement interface {
	Signature (  )  IBmmSignature
	IsBoolean (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
}

type BmmFormalElement struct {
}
