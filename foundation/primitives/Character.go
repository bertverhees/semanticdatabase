package primitives

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

func (p *Character) IsEqual(b IAny) *Boolean {
	v := ConvertToCharacterFromIAny(b)
	return NewBoolean(p.value == v.value)
}

func (p *Character) NotEqual(b IAny) *Boolean {
	v := ConvertToCharacterFromIAny(b)
	return NewBoolean(p.value != v.value)
}

func (p *Character) LessThan(other IOrdered) *Boolean {
	return NewBoolean(p.value < ConvertToCharacterFromIOrdered(other).Value())
}

func (p *Character) LessThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value <= ConvertToCharacterFromIOrdered(other).Value())
}

func (p *Character) GreaterThan(other IOrdered) *Boolean {
	return NewBoolean(p.value > ConvertToCharacterFromIOrdered(other).Value())
}

func (p *Character) GreaterThanOrEqual(other IOrdered) *Boolean {
	return NewBoolean(p.value >= ConvertToCharacterFromIOrdered(other).Value())
}
