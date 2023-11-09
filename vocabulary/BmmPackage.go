package vocabulary

/**
Abstraction of a package as a tree structure whose nodes can contain other
packages and classes. The name may be qualified if it is a top-level package.
*/

// Interface definition
type IBmmPackage interface {
	// From: BMM_MODEL_ELEMENT
	IBmmPackageContainer
	// BMM_PACKAGE
	RootClasses() []IBmmClass
	Path() string
}

// Struct definition
type BmmPackage struct {
	// embedded for Inheritance
	BmmPackageContainer
	// Constants
	// Attributes
	// Member modules in this package.
	Members []IBmmModule `yaml:"members" json:"members" xml:"members"`
}

// CONSTRUCTOR
func NewBmmPackage() *BmmPackage {
	bmmpackage := new(BmmPackage)
	//BmmModelElement
	bmmpackage.documentation = make(map[string]any)
	bmmpackage.extensions = make(map[string]any)
	//BmmPackageContainer
	bmmpackage.packages = make(map[string]IBmmPackage)
	//BmmPackage
	bmmpackage.Members = make([]IBmmModule, 0)

	return bmmpackage
}

// BUILDER
type BmmPackageBuilder struct {
	bmmpackage *BmmPackage
}

func NewBmmPackageBuilder() *BmmPackageBuilder {
	return &BmmPackageBuilder{
		bmmpackage: NewBmmPackage(),
	}
}

// BUILDER ATTRIBUTES
// Member modules in this package.
func (i *BmmPackageBuilder) SetMembers(v []IBmmModule) *BmmPackageBuilder {
	i.bmmpackage.Members = v
	return i
}

// From: BmmPackageContainer
// Child packages; keys all in upper case for guaranteed matching.
func (i *BmmPackageBuilder) SetPackages(v map[string]IBmmPackage) *BmmPackageBuilder {
	i.bmmpackage.packages = v
	return i
}

// From: BmmPackageContainer
// Model element within which a referenceable element is known.
func (i *BmmPackageBuilder) SetScope(v IBmmPackageContainer) *BmmPackageBuilder {
	i.bmmpackage.BmmModelElement.scope = v
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmPackageBuilder) SetName(v string) *BmmPackageBuilder {
	i.bmmpackage.name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmPackageBuilder) SetDocumentation(v map[string]any) *BmmPackageBuilder {
	i.bmmpackage.documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmPackageBuilder) SetExtensions(v map[string]any) *BmmPackageBuilder {
	i.bmmpackage.extensions = v
	return i
}

func (i *BmmPackageBuilder) Build() *BmmPackage {
	return i.bmmpackage
}

//FUNCTIONS
/**
Obtain the set of top-level classes in this package, either from this package
itself or by recursing into the structure until classes are obtained from child
packages. Recurse into each child only far enough to find the first level of
classes.
*/
func (b *BmmPackage) RootClasses() []IBmmClass {
	return nil
}

// Full path of this package back to root package.
func (b *BmmPackage) Path() string {
	return ""
}

// From: BMM_PACKAGE_CONTAINER
// Package at the path a_path .
func (b *BmmPackage) PackageAtPath(a_path string) IBmmPackage {
	return nil
}

// From: BMM_PACKAGE_CONTAINER
/**
Recursively execute action , which is a procedure taking a BMM_PACKAGE argument,
on all members of packages.
*/
func (b *BmmPackage) DoRecursivePackages(action IElProcedureAgent) {
	return
}

// From: BMM_PACKAGE_CONTAINER
/**
True if there is a package at the path a_path ; paths are delimited with
delimiter {BMM_DEFINITIONS} Package_name_delimiter .
*/
func (b *BmmPackage) HasPackagePath(a_path string) bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmPackage) IsRootScope() bool {
	return false
}
