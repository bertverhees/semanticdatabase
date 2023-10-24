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
	PackageAtPath ( a_path string )  IBmmPackage
	DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] ) 
	HasPackagePath ( a_path string )  bool
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
	Path (  )  string
	PackageAtPath ( a_path string )  IBmmPackage
	DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] ) 
	HasPackagePath ( a_path string )  bool
	IsRootScope (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmPackage struct {
	BmmPackageContainer
	BmmModelElement
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
// Package at the path a_path .
func (b *BmmPackage) PackageAtPath ( a_path string )  IBmmPackage {
	return nil
}
/**
	Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
	on all members of packages.
*/
func (b *BmmPackage) DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] )  {
	return
}
/**
	True if there is a package at the path a_path ; paths are delimited with
	delimiter {BMM_DEFINITIONS} Package_name_delimiter .
*/
func (b *BmmPackage) HasPackagePath ( a_path string )  bool {
	return nil
}
// True if this model element is the root of a model structure hierarchy.
func (b *BmmPackage) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
// Full path of this package back to root package.
func (b *BmmPackage) Path (  )  string {
	return nil
}
// Package at the path a_path .
func (b *BmmPackage) PackageAtPath ( a_path string )  IBmmPackage {
	return nil
}
/**
	Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
	on all members of packages.
*/
func (b *BmmPackage) DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] )  {
	return
}
/**
	True if there is a package at the path a_path ; paths are delimited with
	delimiter {BMM_DEFINITIONS} Package_name_delimiter .
*/
func (b *BmmPackage) HasPackagePath ( a_path string )  bool {
	return nil
}
// True if this model element is the root of a model structure hierarchy.
func (b *BmmPackage) IsRootScope (  )  Boolean  Post_result : Result = (scope = self) {
	return nil
}
