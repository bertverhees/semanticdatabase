package vocabulary

import (
	"log"
)

// A BMM model component that contains packages and classes.

const PACKAGE_PATH_DELIMITER = "."

// Interface definition
type IBmmPackageContainer interface {
	// From: BMM_MODEL_ELEMENT
	IBmmModelElement
	//BMM_PACKAGE_CONTAINER
	PackageAtPath(a_path string) IBmmPackage
	DoRecursivePackages(action IElProcedureAgent)
	HasPackagePath(a_path string) bool
	Packages() map[string]IBmmPackage
	SetPackages(packages map[string]IBmmPackage)
	Scope() IBmmPackageContainer
	SetScope(scope IBmmPackageContainer)
}

// Struct definition
type BmmPackageContainer struct {
	// embedded for Inheritance
	BmmModelElement
	// Constants
	// Attributes
	// Child packages; keys all in upper case for guaranteed matching.
	packages map[string]IBmmPackage `yaml:"packages" json:"packages" xml:"packages"`
	// Model element within which a referenceable element is known.
	scope IBmmPackageContainer `yaml:"scope" json:"scope" xml:"scope"`
}

func (b *BmmPackageContainer) Packages() map[string]IBmmPackage {
	return b.packages
}

func (b *BmmPackageContainer) SetPackages(packages map[string]IBmmPackage) {
	b.packages = packages
}

func (b *BmmPackageContainer) Scope() IBmmPackageContainer {
	return b.scope
}

func (b *BmmPackageContainer) SetScope(scope IBmmPackageContainer) {
	b.scope = scope
}

// CONSTRUCTOR
// No constructor or Builder, class is abstract

// FUNCTIONS
// Package at the path a_path .
func (b *BmmPackageContainer) PackageAtPath(a_path string) IBmmPackage {
	log.Fatal("The class BmmPackageContainer is not yet supported")
	return nil
}

/*
*
Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
on all members of packages.
*/
func (b *BmmPackageContainer) DoRecursivePackages(action IElProcedureAgent) {
	log.Fatal("The class BmmPackageContainer is not yet supported")
	//for _, p := range b.packages {
	//	//action.IsCallable()
	//	p.DoRecursivePackages(action)
	//}
	return
}

/*
*
True if there is a package at the path a_path ; paths are delimited with
delimiter {BMM_DEFINITIONS} Package_name_delimiter .
*/
func (b *BmmPackageContainer) HasPackagePath(a_path string) bool {
	log.Fatal("The class BmmPackageContainer is not yet supported")
	//components := strings.Split(a_path, PackagePathDelimiter)
	//packagePath := make([]string, 0)
	//exists := false
	//_, ok := b.(IBmmPackage)
	//if ok {
	//}
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmPackageContainer) IsRootScope() bool {
	log.Fatal("The class BmmPackageContainer is not yet supported")
	return false
}
