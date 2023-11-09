package vocabulary

// A single tuple item, with an optional name.

// Interface definition
type IElTupleItem interface {
}

// Struct definition
type ElTupleItem struct {
	// Attributes
	/**
	Reference to value entity. If Void, this indicates that the item in this
	position is Void, e.g. within a routine call parameter list.
	*/
	Item IElExpression `yaml:"item" json:"item" xml:"item"`
	// Optional name of tuple item.
	Name string `yaml:"name" json:"name" xml:"name"`
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
}

func NewElTupleItemBuilder() *ElTupleItemBuilder {
	return &ElTupleItemBuilder{
		eltupleitem: NewElTupleItem(),
	}
}

//BUILDER ATTRIBUTES
/**
Reference to value entity. If Void, this indicates that the item in this
position is Void, e.g. within a routine call parameter list.
*/
func (i *ElTupleItemBuilder) SetItem(v IElExpression) *ElTupleItemBuilder {
	i.eltupleitem.Item = v
	return i
}

// Optional name of tuple item.
func (i *ElTupleItemBuilder) SetName(v string) *ElTupleItemBuilder {
	i.eltupleitem.Name = v
	return i
}

func (i *ElTupleItemBuilder) Build() *ElTupleItem {
	return i.eltupleitem
}

//FUNCTIONS
