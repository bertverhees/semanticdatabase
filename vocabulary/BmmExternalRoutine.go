package vocabulary

import (
	"vocabulary"
)

/**
	External routine meta-type, containing sufficient meta-data to enable a routine
	in an external library to be called.
*/

type IBmmExternalRoutine interface {
}

type BmmExternalRoutine struct {
	BmmRoutineDefinition
	/**
		External call general meta-data, including target routine name, type mapping
		etc.
	*/
	MetaData	Hash < String , String >	`yaml:"metadata" json:"metadata" xml:"metadata"`
	// Optional argument-mapping meta-data.
	ArgumentMapping	Hash < String , String >	`yaml:"argumentmapping" json:"argumentmapping" xml:"argumentmapping"`
}

