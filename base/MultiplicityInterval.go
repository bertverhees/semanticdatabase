package base

import "vocabulary/generics"

type MultiplicityInterval[T generics.Number] struct {
	ProperInterval[T]
	multiplicityRangeMarker     string `yaml:"multiplicityrangemarker" json:"multiplicityrangemarker" xml:"multiplicityrangemarker"`
	multiplicityUnboundedMarker string `yaml:"multiplicityunboundedmarker" json:"multiplicityunboundedmarker" xml:"multiplicityunboundedmarker"`
}

func NewMultiplicityInterval[T generics.Number]() *MultiplicityInterval[T] {
	mi := new(MultiplicityInterval[T])
	mi.multiplicityRangeMarker = ".."
	mi.multiplicityUnboundedMarker = "*"
	return mi
}

// True if this interval imposes no constraints, i.e. is set to 0..*.
func (mi *MultiplicityInterval[T]) IsOpen() bool {
	return false
}

// True if this interval expresses optionality, i.e. 0..1.
func (mi *MultiplicityInterval[T]) IsOptional() bool {
	return false
}

// True if this interval expresses mandation, i.e. 1..1
func (mi *MultiplicityInterval[T]) IsMandatory() bool {
	return false
}

// True if this interval is set to 0..0.
func (mi *MultiplicityInterval[T]) IsProhibited() bool {
	return false
}
