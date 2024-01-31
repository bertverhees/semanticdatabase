package primitives

type Character struct {
	value rune
}

func (p *Character) Value() rune {
	return p.value
}

func (p *Character) SetValue(value rune) {
	p.value = value
}

func (p *Character) IsEqual(b IAny) IAny {
	return NewBoolean(p.value == b.(*Character).value)
}

func (p *Character) LessThan(other IOrdered) *Boolean {
	//f, ok := other.(*Real)
	//if ok {
	//	return
	//}
	return nil
}

func (p *Character) LessThanOrEqual(other IOrdered) *Boolean {
	return nil
}

func (p *Character) GreaterThan(other IOrdered) *Boolean {
	return nil
}

func (p *Character) GreaterThanOrEqual(other IOrdered) *Boolean {
	return nil
}
