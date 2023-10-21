package vocabulary

/**
	Abstraction of a package as a tree structure whose nodes can contain other
	packages and classes. The name may be qualified if it is a top-level package.
*/

type IBmmPackage interface {
	/**
		Obtain the set of top-level classes in this package, either from this package
		itself or by recursing into the structure until classes are obtained from child
		packages. Recurse into each child only far enough to find the first level of
		classes.
	*/
	root_classes (): List < BMM_CLASS > (  )  List < BMM_CLASS >
// Full path of this package back to root package.
	path (): String (  )  String
}

type BmmPackage struct {
}
