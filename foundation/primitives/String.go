package primitives

type String struct {
	Any
	value string
}

func NewString(value string) *String {
	return &String{value: value}
}

func (p *String) Value() string {
	return p.value
}

func (p *String) SetValue(value string) {
	p.value = value
}

func (p *String) LessThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsString()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value < d.value), nil
}

func (p *String) LessThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsString()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value <= d.value), nil
}

func (p *String) GreaterThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsString()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value > d.value), nil
}

func (p *String) GreaterThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsString()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value >= d.value), nil
}