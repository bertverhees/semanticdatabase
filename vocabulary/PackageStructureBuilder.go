package vocabulary

type BmmDefinitionsBuilder struct {
	Builder
}

func NewBmmDefinitionsBuilder() *BmmDefinitionsBuilder {
	builder := &BmmDefinitionsBuilder{}
	builder.object = NewBmmDefinitions()
	return builder
}

func (i *BmmDefinitionsBuilder) Build() (*BmmDefinitions, []error) {
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*BmmDefinitions), nil
	}
}
