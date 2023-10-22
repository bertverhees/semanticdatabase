package vocabulary

import (
	"vocabulary"
)

// Meta-type based on a non-container generic class, e.g. Packet<Header> .

type IBmmGenericType interface {
	TypeName (  )  string
	TypeSignature (  )  string
	IsAbstract (  )  bool
	FlattenedTypeList (  )  []string
	IsPartiallyClosed (  )  bool
	EffectiveBaseClass (  )  IBmmGenericClass
	IsOpen (  )  bool
}

type BmmGenericType struct {
	/**
		Generic parameters of the root_type in this type specifier. The order must match
		the order of the owning class’s formal generic parameter declarations, and the
		types may be defined types or formal parameter types.
	*/
	GenericParameters	List < BMM_UNITARY_TYPE >	`yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
	// Defining generic class of this type.
	BaseClass	IBmmGenericClass	`yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

/**
	Return the full name of the type including generic parameters, e.g.
	DV_INTERVAL<T> , TABLE<List<THING>,String> .
*/
func (b *BmmGenericType) TypeName (  )  string {
	return nil
}
/**
	Signature form of the type, which for generics includes generic parameter
	constrainer types E.g. Interval<T:Ordered> .
*/
func (b *BmmGenericType) TypeSignature (  )  string {
	return nil
}
// True if base_class.is_abstract or if any (non-open) parameter type is abstract.
func (b *BmmGenericType) IsAbstract (  )  bool {
	return nil
}
/**
	Result is base_class.name followed by names of all generic parameter type names,
	which may be open or closed.
*/
func (b *BmmGenericType) FlattenedTypeList (  )  []string {
	return nil
}
// Returns True if there is any substituted generic parameter.
func (b *BmmGenericType) IsPartiallyClosed (  )  bool {
	return nil
}
// Effective underlying class for this type, abstracting away any container type.
func (b *BmmGenericType) EffectiveBaseClass (  )  IBmmGenericClass {
	return nil
}
/**
	True if all generic parameters from ancestor generic types have been substituted
	in this type.
*/
func (b *BmmGenericType) IsOpen (  )  bool {
	return nil
}
