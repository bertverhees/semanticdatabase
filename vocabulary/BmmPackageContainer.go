package vocabulary

import (
	"vocabulary"
)

// A BMM model component that contains packages and classes.

type IBmmPackageContainer interface {
	PackageAtPath ( a_path string )  IBmmPackage
	DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] ) 
	HasPackagePath ( a_path string )  bool
}

type BmmPackageContainer struct {
	// Child packages; keys all in upper case for guaranteed matching.
	Packages	Hash <String, BMM_PACKAGE >	`yaml:"packages" json:"packages" xml:"packages"`
	// Model element within which a referenceable element is known.
	Scope	IBmmPackageContainer	`yaml:"scope" json:"scope" xml:"scope"`
}

// Package at the path a_path .
func (b *BmmPackageContainer) PackageAtPath ( a_path string )  IBmmPackage {
	return nil
}
/**
	Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
	on all members of packages.
*/
func (b *BmmPackageContainer) DoRecursivePackages ( action EL_PROCEDURE_AGENT [1] )  {
	return
}
/**
	True if there is a package at the path a_path ; paths are delimited with
	delimiter {BMM_DEFINITIONS} Package_name_delimiter .
*/
func (b *BmmPackageContainer) HasPackagePath ( a_path string )  bool {
	return nil
}
