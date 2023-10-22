package vocabulary

import (
	"vocabulary"
)

// Definition of a generic class in an object model.

type IBmmGenericClass interface {
	Suppliers (  )  []string
	Type (  )  IBmmGenericType
	GenericParameterConformanceType ( a_name string )  string
}

type BmmGenericClass struct {
	/**
		List of formal generic parameters, keyed by name. These are defined either
		directly on this class or by the inclusion of an ancestor class which is
		generic.
	*/
	GenericParameters	Hash <String, BMM_PARAMETER_TYPE >	`yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
}

// Add suppliers from generic parameters.
func (b *BmmGenericClass) Suppliers (  )  []string {
	return nil
}
/**
	Generate a fully open BMM_GENERIC_TYPE instance that corresponds to this class
	definition
*/
func (b *BmmGenericClass) Type (  )  IBmmGenericType {
	return nil
}
/**
	For a generic class, type to which generic parameter a_name conforms e.g. if
	this class is Interval <T:Comparable> then the Result will be the single type
	Comparable . For an unconstrained type T , the Result will be Any .
*/
func (b *BmmGenericClass) GenericParameterConformanceType ( a_name string )  string {
	return nil
}
