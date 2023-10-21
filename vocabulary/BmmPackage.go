package vocabulary

/**
	Abstraction of a package as a tree structure whose nodes can contain other
	packages and classes. The name may be qualified if it is a top-level package.
*/


type IBmmPackage interface {
	RootClasses():List<BmmClass> (  )  List < BMM_CLASS >
	Path():String (  )  string
}

type BmmPackage struct {
}
