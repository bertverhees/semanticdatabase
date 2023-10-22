package vocabulary

import (
	"vocabulary"
)

/**
	Core properties of BMM_MODEL , may be used in a serial representation as well,
	such as P_BMM_SCHEMA .
*/

type IBmmModelMetadata interface {
}

type BmmModelMetadata struct {
	// Publisher of model expressed in the schema.
	RmPublisher	string	`yaml:"rmpublisher" json:"rmpublisher" xml:"rmpublisher"`
	// Release of model expressed in the schema as a 3-part numeric, e.g. "3.1.0" .
	RmRelease	string	`yaml:"rmrelease" json:"rmrelease" xml:"rmrelease"`
}

