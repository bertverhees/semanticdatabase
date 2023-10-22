package vocabulary

import (
	"vocabulary"
)

/**
	Parent of meta-types that may be used as the type of any instantiated object
	that is not a container object.
*/

type IBmmUnitaryType interface {
	UnitaryType (  )  IBmmUnitaryType
}

type BmmUnitaryType struct {
}

// Result = self.
func (b *BmmUnitaryType) UnitaryType (  )  IBmmUnitaryType {
	return nil
}
