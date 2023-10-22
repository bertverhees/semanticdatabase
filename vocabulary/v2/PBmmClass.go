package v2

/**
	Definition of persistent form of BMM_CLASS for serialisation to ODIN, JSON, XML
	etc.
*/

type IPBmmClass interface {
	IsGeneric (  )  Boolean  Post : Result := generic_parameter_defs /= Void
	CreateBmmClass (  )  create_bmm_class
	PopulateBmmClass ( a_bmm_schema BMM_MODEL [1] ) 
}

type PBmmClass struct {
}
