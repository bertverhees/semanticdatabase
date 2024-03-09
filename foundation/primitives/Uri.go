package primitives

type Uri struct {
	String
}

func (p *Uri) IsEqual(b IAny) *Boolean {
	v := ConvertToUriFromIAny(b)
	return NewBoolean(p.value == v.value)
}

func (p *Uri) NotEqual(b IAny) *Boolean {
	v := ConvertToUriFromIAny(b)
	return NewBoolean(p.value != v.value)
}
