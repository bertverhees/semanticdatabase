package v2

import (
	"vocabulary"
)

// Persistent form of BMM_CONTAINER_TYPE .

type IPBmmContainerType interface {
	TypeRef (  )  IPBmmBaseType
}

type PBmmContainerType struct {
	/**
		The type of the container. This converts to the root_type in BMM_GENERIC_TYPE .
		Persisted attribute.
	*/
	ContainerType	string	`yaml:"containertype" json:"containertype" xml:"containertype"`
	/**
		Type definition of type , if not a simple String type reference. Persisted
		attribute.
	*/
	TypeDef	IPBmmBaseType	`yaml:"typedef" json:"typedef" xml:"typedef"`
	/**
		The target type; this converts to the first parameter in generic_parameters in
		BMM_GENERIC_TYPE . Persisted attribute.
	*/
	Type	string	`yaml:"type" json:"type" xml:"type"`
	// Result of create_bmm_type() call.
	BmmType	vocabulary.IBmmContainerType	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

/**
	The target type; this converts to the first parameter in generic_parameters in
	BMM_GENERIC_TYPE . Persisted attribute.
*/
func (p *PBmmContainerType) TypeRef (  )  IPBmmBaseType {
	return nil
}
