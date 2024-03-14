package primitives

type Uri struct {
	Any
	value string
}

func NewUri(value string) *Uri {
	return &Uri{value: value}
}

func (p *Uri) LessThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsUri()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value < d.value), nil
}

func (p *Uri) LessThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsUri()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value <= d.value), nil
}

func (p *Uri) GreaterThan(other IOrdered) (*Boolean, error) {
	d, e := other.AsUri()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value > d.value), nil
}

func (p *Uri) GreaterThanOrEqual(other IOrdered) (*Boolean, error) {
	d, e := other.AsUri()
	if e != nil {
		return nil, e
	}
	return NewBoolean(p.value >= d.value), nil
}
