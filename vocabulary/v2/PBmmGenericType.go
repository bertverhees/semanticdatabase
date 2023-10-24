package v2

import (
	"vocabulary"
)

// Persistent form of BMM_GENERIC_TYPE .

type IPBmmGenericType interface {
	GenericParameterRefs (  )  List < P_BMM_TYPE >
	CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass ) 
	AsTypeString (  )  string
}

type PBmmGenericType struct {
	PBmmBaseType
	PBmmType
	// Root type of this generic type, e.g. Interval in Interval<Integer> .
	RootType	string	`yaml:"roottype" json:"roottype" xml:"roottype"`
	/**
		Generic parameters of the root_type in this type specifier if non-simple types.
		The order must match the order of the owning class’s formal generic parameter
		declarations. Persistent attribute.
	*/
	GenericParameterDefs	List < P_BMM_TYPE >	`yaml:"genericparameterdefs" json:"genericparameterdefs" xml:"genericparameterdefs"`
	/**
		Generic parameters of the root_type in this type specifier, if simple types. The
		order must match the order of the owning class’s formal generic parameter
		declarations. Persistent attribute.
	*/
	GenericParameters	[]string	`yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
	// Result of create_bmm_type() call.
	BmmType	BMM_GENERIC_TYPE	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

/**
	Generic parameters of the root_type in this type specifier. The order must match
	the order of the owning class’s formal generic parameter declarations
*/
func (p *PBmmGenericType) GenericParameterRefs (  )  List < P_BMM_TYPE > {
	return nil
}
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmGenericType) CreateBmmType ( a_schema vocabulary.IBmmModel, a_class_def vocabulary.IBmmClass )  {
	return
}
// Formal name of the type for display.
func (p *PBmmGenericType) AsTypeString (  )  string {
	return nil
}
