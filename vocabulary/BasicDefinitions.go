package vocabulary

import (
	"vocabulary"
)

// Defines globally used constant values.

type IBasicDefinitions interface {
}

type BasicDefinitions struct {
	// Carriage return character.
	Cr	rune	`yaml:"cr" json:"cr" xml:"cr"`
	// Line feed character.
	Lf	rune	`yaml:"lf" json:"lf" xml:"lf"`
	AnyTypeName	string	`yaml:"anytypename" json:"anytypename" xml:"anytypename"`
	RegexAnyPattern	string	`yaml:"regexanypattern" json:"regexanypattern" xml:"regexanypattern"`
	DefaultEncoding	string	`yaml:"defaultencoding" json:"defaultencoding" xml:"defaultencoding"`
	NoneTypeName	string	`yaml:"nonetypename" json:"nonetypename" xml:"nonetypename"`
}

