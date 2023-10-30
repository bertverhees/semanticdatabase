package v2

// Persisted form of a model component that contains packages.

// Interface definition
type IPBmmPackageContainer interface {
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
	Packages map[IPBmmPackage]string `yaml:"packages" json:"packages" xml:"packages"`
}

// CONSTRUCTOR
func NewPBmmPackageContainer() *PBmmPackageContainer {
	pbmmpackagecontainer := new(PBmmPackageContainer)
	// Constants
	return pbmmpackagecontainer
}

// BUILDER
type PBmmPackageContainerBuilder struct {
	pbmmpackagecontainer *PBmmPackageContainer
}

func NewPBmmPackageContainerBuilder() *PBmmPackageContainerBuilder {
	return &PBmmPackageContainerBuilder{
		pbmmpackagecontainer: NewPBmmPackageContainer(),
	}
}

//BUILDER ATTRIBUTES
/**
Package structure as a hierarchy of packages each potentially containing names
of classes in that package in the original model.
*/
func (i *PBmmPackageContainerBuilder) SetPackages(v map[IPBmmPackage]string) *PBmmPackageContainerBuilder {
	i.pbmmpackagecontainer.Packages = v
	return i
}

func (i *PBmmPackageContainerBuilder) Build() *PBmmPackageContainer {
	return i.pbmmpackagecontainer
}

//FUNCTIONS
