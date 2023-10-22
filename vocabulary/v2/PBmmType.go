package v2

// Persistent form of BMM_TYPE .

type IPBmmType interface {
	CreateBmmType ( a_schema BMM_MODEL [1], a_class_def BMM_CLASS [1] ) 
	AsTypeString (  )  string
}

type PBmmType struct {
}
