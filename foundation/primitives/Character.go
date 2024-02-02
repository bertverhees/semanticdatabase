package primitives

import (
	"math"
)

type Character struct {
	value rune
}

func NewCharacter(value rune) *Character {
	c := new(Character)
	c.value = value
	return c
}

func (p *Character) Value() rune {
	return p.value
}

func (p *Character) SetValue(value rune) {
	p.value = value
}

func (p *Character) returnCharacterFromIOrdered(ordered IOrdered) *Character {
	var r rune
	switch ordered.(type) {
	case *Double, *Real:
		f := ordered.(*Double).Value()
		r = rune(int(math.Round(f)))
	case *Integer:
		r = rune(ordered.(*Integer).Value())
	case *Integer64:
		r = rune(ordered.(*Integer64).Value())
	case *String:
		panic("Cannot convert *String to *Character")
	}
	return NewCharacter(r)
}

func (p *Character) IsEqual(b IAny) IAny {
	return NewBoolean(p.value == b.(*Character).value)
}

func (p *Character) LessThan(other IOrdered) *Boolean {
	return NewBoolean(p.value < p.returnCharacterFromIOrdered(other).Value())
}

func (p *Character) LessThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value <= p.returnCharacterFromIOrdered(other).Value())
}

func (p *Character) GreaterThan(other IOrdered) *Boolean {
	return NewBoolean(p.value > p.returnCharacterFromIOrdered(other).Value())
}

func (p *Character) GreaterThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value >= p.returnCharacterFromIOrdered(other).Value())
}
