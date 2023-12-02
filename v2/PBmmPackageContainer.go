package v2

// Persisted form of a model component that contains packages.

// Interface definition
type IPBmmPackageContainer interface {
	Packages() map[string]IPBmmPackage
	SetPackages(packages map[string]IPBmmPackage) error
}

// Struct definition
type PBmmPackageContainer struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	/**
	Package structure as a hierarchy of packages each potentially containing names
	of classes in that package in the original model.
	*/
	packages map[string]IPBmmPackage `yaml:"packages" json:"packages" xml:"packages"`
}

func (P *PBmmPackageContainer) Packages() map[string]IPBmmPackage {
	return P.packages
}

func (P *PBmmPackageContainer) SetPackages(packages map[string]IPBmmPackage) error {
	P.packages = packages
	return nil
}

// CONSTRUCTOR
func NewPBmmPackageContainer() *PBmmPackageContainer {
	pbmmpackagecontainer := new(PBmmPackageContainer)
	pbmmpackagecontainer.packages = make(map[string]IPBmmPackage)
	return pbmmpackagecontainer
}

// BUILDER
type PBmmPackageContainerBuilder struct {
	pbmmpackagecontainer *PBmmPackageContainer
	errors               []error
}

func NewPBmmPackageContainerBuilder() *PBmmPackageContainerBuilder {
	return &PBmmPackageContainerBuilder{
		pbmmpackagecontainer: NewPBmmPackageContainer(),
		errors:               make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
Package structure as a hierarchy of packages each potentially containing names
of classes in that package in the original model.
*/
func (i *PBmmPackageContainerBuilder) SetPackages(v map[string]IPBmmPackage) *PBmmPackageContainerBuilder {
	i.AddError(i.pbmmpackagecontainer.SetPackages(v))
	return i
}

func (i *PBmmPackageContainerBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *PBmmPackageContainerBuilder) Build() *PBmmPackageContainer {
	return i.pbmmpackagecontainer
}

//FUNCTIONS
