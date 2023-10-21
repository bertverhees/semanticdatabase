package vocabulary

/**
	Parent of meta-types that may be used as the type of any instantiated object
	that is not a container object.
*/

type IBmmUnitaryType interface {
	UnitaryType():BmmUnitaryType (  )  BMM_UNITARY_TYPE
}

type BmmUnitaryType struct {
}
