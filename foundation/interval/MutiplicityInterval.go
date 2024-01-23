package interval

import "semanticdatabase/generics"

type MultiplicityInterval struct {
	ProperInterval[int]
	multiplicityRangeMarker     string `yaml:"multiplicityrangemarker" json:"multiplicityrangemarker" xml:"multiplicityrangemarker"`
	multiplicityUnboundedMarker string `yaml:"multiplicityunboundedmarker" json:"multiplicityunboundedmarker" xml:"multiplicityunboundedmarker"`
}

func NewMultiplicityInterval[T generics.Number]() *MultiplicityInterval {
	mi := new(MultiplicityInterval)
	mi.multiplicityRangeMarker = ".."
	mi.multiplicityUnboundedMarker = "*"
	return mi
}

// True if this interval imposes no constraints, i.e. is set to 0..*.
func (mi *MultiplicityInterval) IsOpen() bool {
	return false
}

// True if this interval expresses optionality, i.e. 0..1.
func (mi *MultiplicityInterval) IsOptional() bool {
	return false
}

// True if this interval expresses mandation, i.e. 1..1
func (mi *MultiplicityInterval) IsMandatory() bool {
	return false
}

// True if this interval is set to 0..0.
func (mi *MultiplicityInterval) IsProhibited() bool {
	return false
}
