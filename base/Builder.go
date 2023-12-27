package base

type Builder struct {
	errors []error
	object any
}

func (i *Builder) AddError(e error) {
	if e != nil {
		if i.errors == nil {
			i.errors = make([]error, 0)
		}
		i.errors = append(i.errors, e)
	}
}
