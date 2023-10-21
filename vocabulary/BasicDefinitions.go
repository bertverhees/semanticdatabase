package vocabulary

// Defines globally used constant values.

type IBasicDefinitions interface {
}

type BasicDefinitions struct {
	// Carriage return character.
	Cr	rune	`yaml:"cr" json:"cr" xml:"cr"`
	// Line feed character.
	Lf	rune	`yaml:"lf" json:"lf" xml:"lf"`
	AnyTypeName	string	`yaml:"any_type_name" json:"any_type_name" xml:"any_type_name"`
	RegexAnyPattern	string	`yaml:"regex_any_pattern" json:"regex_any_pattern" xml:"regex_any_pattern"`
	DefaultEncoding	string	`yaml:"default_encoding" json:"default_encoding" xml:"default_encoding"`
	NoneTypeName	string	`yaml:"none_type_name" json:"none_type_name" xml:"none_type_name"`
}
