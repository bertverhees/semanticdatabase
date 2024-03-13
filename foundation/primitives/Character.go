package primitives

type Character struct {
	Any
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

func (p *Character) LessThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsCharacter()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value < d.value), nil
}

func (p *Character) LessThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsCharacter()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value <= d.value), nil
}

func (p *Character) GreaterThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsCharacter()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value > d.value), nil
}

func (p *Character) GreaterThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsCharacter()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value >= d.value), nil
}
