package v2

/**
	Persisted form of a package as a tree structure whose nodes can contain more
	packages and/or classes.
*/


type IPBmmPackage interface {
	Merge(Other:PBmmPackage[1]) ( other P_BMM_PACKAGE [1] ) 
	CreateBmmPackageDefinition (  )  create_bmm_package_definition
}

type PBmmPackage struct {
}
