package primitives

type Octet struct {
	Any
	value uint8
}

func NewOctet(value uint8) *Octet {
	d := new(Octet)
	d.value = value
	return d
}

func (p *Octet) Value() uint8 {
	return p.value
}

func (p *Octet) SetValue(value uint8) {
	p.value = value
}

func (p *Octet) LessThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsOctet()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value < d.value), nil
}

func (p *Octet) LessThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsOctet()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value <= d.value), nil
}

func (p *Octet) GreaterThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsOctet()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value > d.value), nil
}

func (p *Octet) GreaterThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsOctet()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value >= d.value), nil
}
