package vocabulary

type Builder struct {
	errors []error
	object any
}

func (i *Builder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}
