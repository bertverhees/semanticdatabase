package vocabulary

// _type reference to a single type i.e. not generic or container type.

// Interface definition

// Struct definition
type BmmSimpleType struct {
	BmmModelType
	// Attributes
	// Defining class of this type.
	//getter and setter in parent
	baseClass IBmmSimpleClass `yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

// CONSTRUCTOR
func NewBmmSimpleType() *BmmSimpleType {
	bmmsimpletype := new(BmmSimpleType)
	// BaseClass IBmmSimpleType (redefined)
	return bmmsimpletype
}

// BUILDER
type BmmSimpleTypeBuilder struct {
	bmmsimpletype *BmmSimpleType
	errors        []error
}

func NewBmmSimpleTypeBuilder() *BmmSimpleTypeBuilder {
	return &BmmSimpleTypeBuilder{
		bmmsimpletype: NewBmmSimpleType(),
		errors:        make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Defining class of this type.
func (i *BmmSimpleTypeBuilder) SetBaseClass(v IBmmSimpleClass) *BmmSimpleTypeBuilder {
	i.AddError(i.bmmsimpletype.SetBaseClass(v))
	return i
}

// From: BmmModelType
func (i *BmmSimpleTypeBuilder) SetValueConstraint(v IBmmValueSetSpec) *BmmSimpleTypeBuilder {
	i.AddError(i.bmmsimpletype.SetValueConstraint(v))
	return i
}

func (i *BmmSimpleTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmSimpleTypeBuilder) Build() *BmmSimpleType {
	return i.bmmsimpletype
}

// FUNCTIONS
// (effected) result is base_class.name .
func (b *BmmSimpleType) TypeName() string {
	return ""
}

// (effected) result is base_class.is_abstract .
func (b *BmmSimpleType) IsAbstract() bool {
	return false
}

// (effected) result is base_class.name .
func (b *BmmSimpleType) FlattenedTypeList() []string {
	return nil
}

// Main design class for this type, from which properties etc can be extracted.
func (b *BmmSimpleType) EffectiveBaseClass() IBmmSimpleClass {
	return nil
}
