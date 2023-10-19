package vocabulary

// Defines globally used constant values.

type BasicDefinition struct {
		// Carriage return character.
	CR	rune	`yaml:"cr" json:"cr" xml:"cr"`
	// Line feed character.
	LF	rune	`yaml:"lf" json:"lf" xml:"lf"`
	AnyTypeName	string	`yaml:"anynametype" json:"anynametype" xml:"anynametype"`
	RegexAnyPattern	string	`yaml:"regexanypattern" json:"regexanypattern" xml:"regexanypattern"`
	DefaultEncoding	string	`yaml:"defaultencoding" json:"defaultencoding" xml:"defaultencoding"`
	NoneTypeName	string	`yaml:"nonetypename" json:"nonetypename" xml:"nonetypename"`
}
