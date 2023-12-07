package vocabulary

type BmmDefinitionsBuilder struct {
	Builder
	bmmdefinitions *BmmDefinitions
}

func NewBmmDefinitionsBuilder() *BmmDefinitionsBuilder {
	builder := &BmmDefinitionsBuilder{
		bmmdefinitions: NewBmmDefinitions(),
	}
	builder.errors = make([]error, 0)
	return builder
}

func (i *BmmDefinitionsBuilder) Build() *BmmDefinitions {
	return i.bmmdefinitions
}
