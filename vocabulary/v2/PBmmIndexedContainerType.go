package v2

import (
	"vocabulary"
)


type IPBmmIndexedContainerType interface {
}

type PBmmIndexedContainerType struct {
	IndexType	string	`yaml:"indextype" json:"indextype" xml:"indextype"`
	// Result of create_bmm_type() call.
	BmmType	BMM_INDEXED_CONTAINER_TYPE	`yaml:"bmmtype" json:"bmmtype" xml:"bmmtype"`
}

