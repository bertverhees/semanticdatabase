package primitives

type Hex struct {
	value int16
}

func NewHex(value int16) *Hex {
	d := new(Hex)
	d.value = value
	return d
}

func (p *Hex) Value() int16 {
	return p.value
}

func (p *Hex) SetValue(value int16) {
	p.value = value
}

func (p *Hex) IsEqual(b IAny) IAny {
	return NewBoolean(p.value == b.(*Hex).value)
}

func (p *Hex) LessThan(other IOrdered) *Boolean {
	//f, ok := other.(*Hex)
	//if ok {
	//	return
	//}
	return nil
}

func (p *Hex) LessThanOrEqual(other IOrdered) *Boolean {
	return nil
}

func (p *Hex) GreaterThan(other IOrdered) *Boolean {
	return nil
}

func (p *Hex) GreaterThanOrEqual(other IOrdered) *Boolean {
	return nil
}
