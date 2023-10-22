package vocabulary

/**
	Abstraction of a package as a tree structure whose nodes can contain other
	packages and classes. The name may be qualified if it is a top-level package.
*/

type IBmmPackage interface {
	RootClasses (  )  List < BMM_CLASS >
	Path (  )  string
}

type BmmPackage struct {
}
