package vocabulary

// A formal element having a name, type and a type-based signature.

type IBmmFormalElement interface {
	Signature():BmmSignature (  )  BMM_SIGNATURE
	IsBoolean():BooleanPostResult:Result=Type().equal({bmmModel}.booleanTypeDefinition()) (  )  Boolean  Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition())
}

type BmmFormalElement struct {
}
