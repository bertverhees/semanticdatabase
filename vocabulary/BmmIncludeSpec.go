package vocabulary

// Schema inclusion structure.

// Interface definition
type IBmmIncludeSpec interface {
	Id() string
	SetId(id string) error
}

// Struct definition
type BmmIncludeSpec struct {
	id string `yaml:"id" json:"id" xml:"id"`
}

func (b *BmmIncludeSpec) Id() string {
	return b.id
}

func (b *BmmIncludeSpec) SetId(id string) error {
	b.id = id
	return nil
}

// CONSTRUCTOR
func NewBmmIncludeSpec() *BmmIncludeSpec {
	bmmincludespec := new(BmmIncludeSpec)
	return bmmincludespec
}

// BUILDER
type BmmIncludeSpecBuilder struct {
	bmmincludespec *BmmIncludeSpec
	errors         []error
}

func NewBmmIncludeSpecBuilder() *BmmIncludeSpecBuilder {
	return &BmmIncludeSpecBuilder{
		bmmincludespec: NewBmmIncludeSpec(),
		errors:         make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Full identifier of the included schema, e.g. "openehr_primitive_types_1.0.2" .
func (i *BmmIncludeSpecBuilder) SetId(v string) *BmmIncludeSpecBuilder {
	i.AddError(i.bmmincludespec.SetId(v))
	return i
}

func (i *BmmIncludeSpecBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmIncludeSpecBuilder) Build() *BmmIncludeSpec {
	return i.bmmincludespec
}

//FUNCTIONS
