package vocabulary

/**
	Definition of a simple class, i.e. a class that has no generic parameters and is
	1:1 with the type it generates.
*/

type IBmmSimpleClass interface {
	Type (  )  IBmmSimpleType
}

type BmmSimpleClass struct {
}
