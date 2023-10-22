package vocabulary

import (
	"vocabulary"
)

/**
	Definition of a range-constrained class whose value range is defined by
	reference to a 'value set' within an external resource, e.g. a reference data
	service.
*/

type IBmmValueSetSpec interface {
}

type BmmValueSetSpec struct {
	/**
		Identifier of a resource (typically available as a service) that contains legal
		values of a specific type. This is typically a URI but need not be.
	*/
	ResourceId	string	`yaml:"resourceid" json:"resourceid" xml:"resourceid"`
	/**
		Identifier of a value set within the resource identified by resource_id , which
		specifies the set of legal values of a type. This might be a URI, but need not
		be.
	*/
	ValueSetId	string	`yaml:"valuesetid" json:"valuesetid" xml:"valuesetid"`
}

