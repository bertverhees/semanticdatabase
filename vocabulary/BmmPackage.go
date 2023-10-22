package vocabulary

import (
	"vocabulary"
)

/**
	Abstraction of a package as a tree structure whose nodes can contain other
	packages and classes. The name may be qualified if it is a top-level package.
*/

type IBmmPackage interface {
	RootClasses (  )  List < BMM_CLASS >
	Path (  )  string
}

type BmmPackage struct {
	// Member modules in this package.
	Members	List < BMM_MODULE >	`yaml:"members" json:"members" xml:"members"`
}

/**
	Obtain the set of top-level classes in this package, either from this package
	itself or by recursing into the structure until classes are obtained from child
	packages. Recurse into each child only far enough to find the first level of
	classes.
*/
func (b *BmmPackage) RootClasses (  )  List < BMM_CLASS > {
	return nil
}
// Full path of this package back to root package.
func (b *BmmPackage) Path (  )  string {
	return nil
}
