package vocabulary

/**
Built-in meta-type that expresses the type structure of any referenceable
element of a model. Consists of potential arguments and result , with
constraints in descendants determining the exact form.
*/

// Interface definition

// Struct definition
type BmmSignature struct {
	BmmBuiltinType
	// Attributes
	// result type of signature.
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
// result type of signature.
func (i *BmmSignatureBuilder) SetResultType(v IBmmType) *BmmSignatureBuilder {
	i.AddError(i.bmmsignature.SetResultType(v))
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
