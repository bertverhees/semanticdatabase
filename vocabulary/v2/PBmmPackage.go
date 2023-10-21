package v2

/**
	Persisted form of a package as a tree structure whose nodes can contain more
	packages and/or classes.
*/

type IPBmmPackage interface {
	/**
		Merge packages and classes from other (from an included P_BMM_SCHEMA ) into this
		package.
	*/
	merge ( other: P_BMM_PACKAGE [1] ) ( other P_BMM_PACKAGE [1] ) 
	/**
		Generate a BMM_PACKAGE_DEFINITION object and write it to bmm_package_definition
		.
	*/
	create_bmm_package_definition (  )  create_bmm_package_definition
}

type PBmmPackage struct {
}
