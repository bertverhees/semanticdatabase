package v2

// Persistent form of BMM_CONTAINER_PROPERTY .

type IPBmmContainerProperty interface {
	CreateBmmProperty ( a_bmm_schema BMM_MODEL [1], a_class_def BMM_CLASS [1] ) 
}

type PBmmContainerProperty struct {
}
