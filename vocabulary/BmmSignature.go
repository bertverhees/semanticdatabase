package vocabulary

/**
Built-in meta-type that expresses the type structure of any referenceable
element of a model. Consists of potential arguments and result , with
constraints in descendants determining the exact form.
*/

// Interface definition
type IBmmSignature interface {
	IBmmBuiltinType
	//BMM_SIGNATURE
	ResultType() IBmmType
	SetResultType(resultType IBmmType) error
	FlattenedTypeList() []string
}

// Struct definition
type BmmSignature struct {
	BmmBuiltinType
	// Attributes
	// Result type of signature.
	resultType IBmmType `yaml:"resulttype" json:"resulttype" xml:"resulttype"`
}

func (b *BmmSignature) ResultType() IBmmType {
	return b.resultType
}

func (b *BmmSignature) SetResultType(resultType IBmmType) error {
	b.resultType = resultType
	return nil
}

// CONSTRUCTOR
func NewBmmSignature() *BmmSignature {
	bmmsignature := new(BmmSignature)
	// Base name (built-in).
	bmmsignature.baseName = "Signature"
	return bmmsignature
}

// BUILDER
type BmmSignatureBuilder struct {
	bmmsignature *BmmSignature
	errors       []error
}

func NewBmmSignatureBuilder() *BmmSignatureBuilder {
	return &BmmSignatureBuilder{
		bmmsignature: NewBmmSignature(),
		errors:       make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Result type of signature.
func (i *BmmSignatureBuilder) SetResultType(v IBmmType) *BmmSignatureBuilder {
	i.AddError(i.bmmsignature.SetResultType(v))
	return i
}

func (i *BmmSignatureBuilder) SetBaseName(v string) *BmmSignatureBuilder {
	i.AddError(i.bmmsignature.SetBaseName(v))
	return i
}

func (i *BmmSignatureBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmSignatureBuilder) Build() *BmmSignature {
	return i.bmmsignature
}

//FUNCTIONS
/**
(effected) Return the logical set (i.e. unique items) consisting of
argument_types.flattened_type_list () and result_type.flattened_type_list () .
*/
func (b *BmmSignature) FlattenedTypeList() []string {
	return nil
}

// From: BMM_BUILTIN_TYPE
// Return False.
func (b *BmmSignature) IsAbstract() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// Return True.
func (b *BmmSignature) IsPrimitive() bool {
	return false
}

// From: BMM_BUILTIN_TYPE
// (effected) Return base_name .
func (b *BmmSignature) TypeBaseName() string {
	return b.BaseName()
}

// From: BMM_BUILTIN_TYPE
// Return base_name .
func (b *BmmSignature) TypeName() string {
	return ""
}

// From: BMM_EFFECTIVE_TYPE
// Result = self.
func (b *BmmSignature) EffectiveType() IBmmEffectiveType {
	return nil
}

// From: BMM_UNITARY_TYPE
// Result = self.
func (b *BmmSignature) UnitaryType() IBmmUnitaryType {
	return nil
}

// From: BMM_TYPE
/**
Signature form of the type name, which for generics includes generic parameter
constrainer types E.g. Interval<T:Ordered> . Defaults to the value of
type_name() .
*/
func (b *BmmSignature) TypeSignature() string {
	return ""
}
