package vocabulary

// A BMM model component that contains packages and classes.

type IBmmPackageContainer interface {
	PackageAtPath(APath:String[1]):BmmPackage ( a_path String[1] )  BMM_PACKAGE
	DoRecursivePackages(Action:ElProcedureAgent[1]) ( action EL_PROCEDURE_AGENT [1] ) 
	HasPackagePath(APath:String[1]):Boolean ( a_path String[1] )  Boolean
}

type BmmPackageContainer struct {
}
