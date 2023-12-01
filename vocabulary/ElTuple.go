package vocabulary

// Defines an array of optionally named items each of any type.

// Interface definition
type IElTuple interface {
	IElExpression
	Items() []IElTupleItem
	SetItems(items []IElTupleItem) error
	Type() IBmmTupleType
	SetType(_type IBmmTupleType) error
}

// Struct definition
type ElTuple struct {
	ElExpression
	// Attributes
	/**
	items in the tuple, potentially with names. Typical use is to represent an
	argument list to routine call.
	*/
	items []IElTupleItem `yaml:"items" json:"items" xml:"items"`
	// Static type inferred from literal value.
	_type IBmmTupleType `yaml:"type" json:"type" xml:"type"`
}

func (e *ElTuple) Items() []IElTupleItem {
	return e.items
}

func (e *ElTuple) SetItems(items []IElTupleItem) error {
	e.items = items
	return nil
}

func (e *ElTuple) Type() IBmmTupleType {
	return e._type
}

func (e *ElTuple) SetType(_type IBmmTupleType) error {
	e._type = _type
	return nil
}

// CONSTRUCTOR
func NewElTuple() *ElTuple {
	eltuple := new(ElTuple)
	eltuple.items = make([]IElTupleItem, 0)
	// Constants
	return eltuple
}

// BUILDER
type ElTupleBuilder struct {
	eltuple *ElTuple
	errors  []error
}

func NewElTupleBuilder() *ElTupleBuilder {
	return &ElTupleBuilder{
		eltuple: NewElTuple(),
		errors:  make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
items in the tuple, potentially with names. Typical use is to represent an
argument list to routine call.
*/
func (i *ElTupleBuilder) SetItems(v []IElTupleItem) *ElTupleBuilder {
	i.AddError(i.eltuple.SetItems(v))
	return i
}

// Static type inferred from literal value.
func (i *ElTupleBuilder) SetType(v IBmmTupleType) *ElTupleBuilder {
	i.AddError(i.eltuple.SetType(v))
	return i
}

func (i *ElTupleBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElTupleBuilder) Build() *ElTuple {
	return i.eltuple
}

// FUNCTIONS
