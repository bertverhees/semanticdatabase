package vocabulary

import (
	"vocabulary"
)

// Persisted form of a model component that contains packages.

type IPBmmPackageContainer interface {
}

type PBmmPackageContainer struct {
	/**
		Package structure as a hierarchy of packages each potentially containing names
		of classes in that package in the original model.
	*/
	Packages	Hash< P_BMM_PACKAGE , String >	`yaml:"packages" json:"packages" xml:"packages"`
}

