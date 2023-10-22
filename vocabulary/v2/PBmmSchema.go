package v2

import (
	"vocabulary"
)

// Persisted form of BMM_SCHEMA .

type IPBmmSchema interface {
	ValidateCreatedPreState (  ) 
	LoadFinalisePreState (  ) 
	Merge ( other P_BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id)
	Validate (  ) 
	CreateBmmModelPreState (  ) 
	CanonicalPackages (  )  IPBmmPackage
}

type PBmmSchema struct {
	// Primitive type definitions. Persisted attribute.
	PrimitiveTypes	List < P_BMM_CLASS >	`yaml:"primitivetypes" json:"primitivetypes" xml:"primitivetypes"`
	// Class definitions. Persisted attribute.
	ClassDefinitions	List < P_BMM_CLASS >	`yaml:"classdefinitions" json:"classdefinitions" xml:"classdefinitions"`
}

// Implementation of validate_created()
func (p *PBmmSchema) ValidateCreatedPreState (  )  {
	return
}
// Implementation of load_finalise()
func (p *PBmmSchema) LoadFinalisePreState (  )  {
	return
}
// Implementation of merge()
func (p *PBmmSchema) Merge ( other P_BMM_SCHEMA [1] )  state = State_includes_pending Pre_other_valid : includes_to_process.has (included_schema.schema_id) {
	return nil
}
// Implementation of validate()
func (p *PBmmSchema) Validate (  )  {
	return
}
// Implementation of create_bmm_model()
func (p *PBmmSchema) CreateBmmModelPreState (  )  {
	return
}
/**
	Package structure in which all top-level qualified package names like xx.yy.zz
	have been expanded out to a hierarchy of BMM_PACKAGE objects.
*/
func (p *PBmmSchema) CanonicalPackages (  )  IPBmmPackage {
	return nil
}
