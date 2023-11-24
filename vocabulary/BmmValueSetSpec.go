package vocabulary

/**
definition of a range-constrained class whose value range is defined by
reference to a 'value set' within an external resource, e.g. a reference data
service.
*/

// Interface definition
type IBmmValueSetSpec interface {
	ResourceId() string
	SetResourceId(resourceId string) error
	ValueSetId() string
	SetValueSetId(valueSetId string) error
}

// Struct definition
type BmmValueSetSpec struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	/**
	Identifier of a resource (typically available as a service) that contains legal
	values of a specific type. This is typically a URI but need not be.
	*/
	resourceId string `yaml:"resourceid" json:"resourceid" xml:"resourceid"`
	/**
	Identifier of a value set within the resource identified by resource_id , which
	specifies the set of legal values of a type. This might be a URI, but need not
	be.
	*/
	valueSetId string `yaml:"valuesetid" json:"valuesetid" xml:"valuesetid"`
}

func (b *BmmValueSetSpec) ResourceId() string {
	return b.resourceId
}

func (b *BmmValueSetSpec) SetResourceId(resourceId string) error {
	b.resourceId = resourceId
	return nil
}

func (b *BmmValueSetSpec) ValueSetId() string {
	return b.valueSetId
}

func (b *BmmValueSetSpec) SetValueSetId(valueSetId string) error {
	b.valueSetId = valueSetId
	return nil
}

// CONSTRUCTOR
func NewBmmValueSetSpec() *BmmValueSetSpec {
	bmmvaluesetspec := new(BmmValueSetSpec)
	// Constants
	return bmmvaluesetspec
}

// BUILDER
type BmmValueSetSpecBuilder struct {
	bmmvaluesetspec *BmmValueSetSpec
	errors          []error
}

func NewBmmValueSetSpecBuilder() *BmmValueSetSpecBuilder {
	return &BmmValueSetSpecBuilder{
		bmmvaluesetspec: NewBmmValueSetSpec(),
		errors:          make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
Identifier of a resource (typically available as a service) that contains legal
values of a specific type. This is typically a URI but need not be.
*/
func (i *BmmValueSetSpecBuilder) SetResourceId(v string) *BmmValueSetSpecBuilder {
	i.AddError(i.bmmvaluesetspec.SetResourceId(v))
	return i
}

/*
*
Identifier of a value set within the resource identified by resource_id , which
specifies the set of legal values of a type. This might be a URI, but need not
be.
*/
func (i *BmmValueSetSpecBuilder) SetValueSetId(v string) *BmmValueSetSpecBuilder {
	i.AddError(i.bmmvaluesetspec.SetValueSetId(v))
	return i
}

func (i *BmmValueSetSpecBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmValueSetSpecBuilder) Build() *BmmValueSetSpec {
	return i.bmmvaluesetspec
}

//FUNCTIONS
