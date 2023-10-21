package v2

// Persistent form of BMM_PROPERTY .

type IPBmmProperty interface {
	CreateBmmProperty(ABmmSchema:BmmModel[1],AClassDef:BmmClass[1]) ( a_bmm_schema BMM_MODEL [1], a_class_def BMM_CLASS [1] ) 
}

type PBmmProperty struct {
}
