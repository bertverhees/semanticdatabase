package vocabulary

// A BMM model component that contains packages and classes.

type IBmmPackageContainer interface {
// Package at the path a_path .
	package_at_path ( a_path: String[1] ): BMM_PACKAGE ( a_path String[1] )  BMM_PACKAGE
	/**
		Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
		on all members of packages.
	*/
	do_recursive_packages ( action: EL_PROCEDURE_AGENT [1] ) ( action EL_PROCEDURE_AGENT [1] ) 
	/**
		True if there is a package at the path a_path ; paths are delimited with
		delimiter {BMM_DEFINITIONS} Package_name_delimiter .
	*/
	has_package_path ( a_path: String[1] ): Boolean ( a_path String[1] )  Boolean
}

type BmmPackageContainer struct {
}
