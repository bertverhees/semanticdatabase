package vocabulary

type Builder struct {
	errors          []error
}

func NewBuilder() *Builder {
	return &Builder{
		errors:                 make([]error, 0),
	}
}

func (i *Builder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

