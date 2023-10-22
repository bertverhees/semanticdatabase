package vocabulary

// A BMM model component that contains packages and classes.

type IBmmPackageContainer interface {
	PackageAtPath ( a_path string )  IBmmPackage
	DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] ) 
	HasPackagePath ( a_path string )  Boolean
}

type BmmPackageContainer struct {
}
