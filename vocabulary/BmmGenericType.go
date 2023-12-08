package vocabulary

// Meta-type based on a non-container generic class, e.g. Packet<Header> .

// Interface definition

// Struct definition
type BmmGenericType struct {
	BmmModelType
	// Attributes
	/**
	Generic parameters of the root_type in this type specifier. The order must match
	the order of the owning class’s formal generic parameter declarations, and the
	types may be defined types or formal parameter types.
	*/
	genericParameters []IBmmUnitaryType `yaml:"genericparameters" json:"genericparameters" xml:"genericparameters"`
	// Defining generic class of this type.
	baseClass IBmmGenericClass `yaml:"baseclass" json:"baseclass" xml:"baseclass"`
}

func (b *BmmGenericType) GenericParameters() []IBmmUnitaryType {
	return b.genericParameters
}

func (b *BmmGenericType) SetGenericParameters(genericParameters []IBmmUnitaryType) error {
	b.genericParameters = genericParameters
	return nil
}

// CONSTRUCTOR
func NewBmmGenericType() *BmmGenericType {
	bmmgenerictype := new(BmmGenericType)
	//BMM_GENERIC_TYPE
	bmmgenerictype.genericParameters = make([]IBmmUnitaryType, 0)
	// (redefined) BaseClass IBmmGenericType
	return bmmgenerictype
}

// BUILDER
type BmmGenericTypeBuilder struct {
	bmmgenerictype *BmmGenericType
	errors         []error
}

func NewBmmGenericTypeBuilder() *BmmGenericTypeBuilder {
	return &BmmGenericTypeBuilder{
		bmmgenerictype: NewBmmGenericType(),
		errors:         make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
Generic parameters of the root_type in this type specifier. The order must match
the order of the owning class’s formal generic parameter declarations, and the
types may be defined types or formal parameter types.
*/
func (i *BmmGenericTypeBuilder) SetGenericParameters(v []IBmmUnitaryType) *BmmGenericTypeBuilder {
	i.AddError(i.bmmgenerictype.SetGenericParameters(v))
	return i
}

// Defining generic class of this type.
func (i *BmmGenericTypeBuilder) SetBaseClass(v IBmmGenericClass) *BmmGenericTypeBuilder {
	i.AddError(i.bmmgenerictype.SetBaseClass(v))
	return i
}

// From: BmmModelType
func (i *BmmGenericTypeBuilder) SetValueConstraint(v IBmmValueSetSpec) *BmmGenericTypeBuilder {
	i.AddError(i.bmmgenerictype.SetValueConstraint(v))
	return i
}

func (i *BmmGenericTypeBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmGenericTypeBuilder) Build() *BmmGenericType {
	return i.bmmgenerictype
}

//FUNCTIONS
/**
(effected) Return the full name of the type including generic parameters, e.g.
DV_INTERVAL<T> , TABLE<List<THING>,String> .
*/
func (b *BmmGenericType) TypeName() string {
	return ""
}

/*
*
(effected) Signature form of the type, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> .
*/
func (b *BmmGenericType) TypeSignature() string {
	return ""
}

// (effected) True if base_class.is_abstract or if any (non-open) parameter type is abstract.
func (b *BmmGenericType) IsAbstract() bool {
	return false
}

/*
*
(effected) result is base_class.name followed by names of all generic parameter type names,
which may be open or closed.
*/
func (b *BmmGenericType) FlattenedTypeList() []string {
	return nil
}

// Returns True if there is any substituted generic parameter.
func (b *BmmGenericType) IsPartiallyClosed() bool {
	return false
}

// Effective underlying class for this type, abstracting away any container type.
func (b *BmmGenericType) EffectiveBaseClass() IBmmGenericClass {
	return nil
}

/*
*
True if all generic parameters from ancestor generic types have been substituted
in this type.
*/
func (b *BmmGenericType) IsOpen() bool {
	return false
}
