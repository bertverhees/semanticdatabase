package vocabulary

/**
	Definition of a range-constrained class whose value range is defined by
	reference to a 'value set' within an external resource, e.g. a reference data
	service.
*/

// Interface definition
type IBmmValueSetSpec interface {
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
	ResourceId	string	`yaml:"resourceid" json:"resourceid" xml:"resourceid"`
	/**
		Identifier of a value set within the resource identified by resource_id , which
		specifies the set of legal values of a type. This might be a URI, but need not
		be.
	*/
	ValueSetId	string	`yaml:"valuesetid" json:"valuesetid" xml:"valuesetid"`
}

//CONSTRUCTOR
func NewBmmValueSetSpec() *BmmValueSetSpec {
	bmmvaluesetspec := new(BmmValueSetSpec)
	// Constants
	return bmmvaluesetspec
}
//BUILDER
type BmmValueSetSpecBuilder struct {
	bmmvaluesetspec *BmmValueSetSpec
}

func NewBmmValueSetSpecBuilder() *BmmValueSetSpecBuilder {
	 return &BmmValueSetSpecBuilder {
		bmmvaluesetspec : NewBmmValueSetSpec(),
	}
}

//BUILDER ATTRIBUTES
/**
	Identifier of a resource (typically available as a service) that contains legal
	values of a specific type. This is typically a URI but need not be.
*/
func (i *BmmValueSetSpecBuilder) SetResourceId ( v string ) *BmmValueSetSpecBuilder{
	i.bmmvaluesetspec.ResourceId = v
	return i
}
/**
	Identifier of a value set within the resource identified by resource_id , which
	specifies the set of legal values of a type. This might be a URI, but need not
	be.
*/
func (i *BmmValueSetSpecBuilder) SetValueSetId ( v string ) *BmmValueSetSpecBuilder{
	i.bmmvaluesetspec.ValueSetId = v
	return i
}

func (i *BmmValueSetSpecBuilder) Build() *BmmValueSetSpec {
	 return i.bmmvaluesetspec
}

//FUNCTIONS
