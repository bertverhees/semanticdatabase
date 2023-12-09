package vocabulary

type BmmDefinitionsBuilder struct {
	Builder
}

func NewBmmDefinitionsBuilder() *BmmDefinitionsBuilder {
	builder := &BmmDefinitionsBuilder{}
	builder.object = NewBmmDefinitions()
	builder.errors = make([]error, 0)
	return builder
}

func (i *BmmDefinitionsBuilder) Build() (*BmmDefinitions, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmDefinitions), nil
	}
}
