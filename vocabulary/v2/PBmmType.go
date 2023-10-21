package v2

// Persistent form of BMM_TYPE .

type IPBmmType interface {
	CreateBmmType(ASchema:BmmModel[1],AClassDef:BmmClass[1]) ( a_schema BMM_MODEL [1], a_class_def BMM_CLASS [1] ) 
	AsTypeString():String (  )  string
}

type PBmmType struct {
}
