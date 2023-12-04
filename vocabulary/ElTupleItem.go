package vocabulary

// A single tuple item, with an optional name.

// Interface definition
type IElTupleItem interface {
	Item() IElExpression
	SetItem(item IElExpression) error
	Name() string
	SetName(name string) error
}

// Struct definition
type ElTupleItem struct {
	// Attributes
	/**
	Reference to value entity. If Void, this indicates that the item in this
	position is Void, e.g. within a routine call parameter list.
	*/
	item IElExpression `yaml:"item" json:"item" xml:"item"`
	// Optional name of tuple item.
	name string `yaml:"name" json:"name" xml:"name"`
}

func (e *ElTupleItem) Item() IElExpression {
	return e.item
}

func (e *ElTupleItem) SetItem(item IElExpression) error {
	e.item = item
	return nil
}

func (e *ElTupleItem) Name() string {
	return e.name
}

func (e *ElTupleItem) SetName(name string) error {
	e.name = name
	return nil
}

// CONSTRUCTOR
func NewElTupleItem() *ElTupleItem {
	eltupleitem := new(ElTupleItem)
	// Constants
	return eltupleitem
}

// BUILDER
type ElTupleItemBuilder struct {
	eltupleitem *ElTupleItem
	errors      []error
}

func NewElTupleItemBuilder() *ElTupleItemBuilder {
	return &ElTupleItemBuilder{
		eltupleitem: NewElTupleItem(),
		errors:      make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
Reference to value entity. If Void, this indicates that the item in this
position is Void, e.g. within a routine call parameter list.
*/
func (i *ElTupleItemBuilder) SetItem(v IElExpression) *ElTupleItemBuilder {
	i.AddError(i.eltupleitem.SetItem(v))
	return i
}

// Optional name of tuple item.
func (i *ElTupleItemBuilder) SetName(v string) *ElTupleItemBuilder {
	i.AddError(i.eltupleitem.SetName(v))
	return i
}

func (i *ElTupleItemBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *ElTupleItemBuilder) Build() *ElTupleItem {
	return i.eltupleitem
}

//FUNCTIONS
