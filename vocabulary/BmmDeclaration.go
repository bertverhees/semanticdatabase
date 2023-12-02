package vocabulary

// Declaration of a writable variable, associating a name with a type.

// Interface definition
type IBmmDeclaration interface {
	IBmmSimpleStatement
	Name() string
	SetName(name string) error
	Result() IElWritableVariable
	SetResult(result IElWritableVariable) error
	Type() IBmmType
	SetType(_type IBmmType) error
}

// Struct definition
type BmmDeclaration struct {
	BmmSimpleStatement
	// Attributes
	name   string              `yaml:"name" json:"name" xml:"name"`
	result IElWritableVariable `yaml:"result" json:"result" xml:"result"`
	// The declared type of the variable.
	_type IBmmType `yaml:"type" json:"type" xml:"type"`
}

func (b *BmmDeclaration) Name() string {
	return b.name
}

func (b *BmmDeclaration) SetName(name string) error {
	b.name = name
	return nil
}

func (b *BmmDeclaration) Result() IElWritableVariable {
	return b.result
}

func (b *BmmDeclaration) SetResult(result IElWritableVariable) error {
	b.result = result
	return nil
}

func (b *BmmDeclaration) Type() IBmmType {
	return b._type
}

func (b *BmmDeclaration) SetType(_type IBmmType) error {
	b._type = _type
	return nil
}

// CONSTRUCTOR
func NewBmmDeclaration() *BmmDeclaration {
	bmmdeclaration := new(BmmDeclaration)
	// Constants
	return bmmdeclaration
}

// BUILDER
type BmmDeclarationBuilder struct {
	bmmdeclaration *BmmDeclaration
	errors         []error
}

func NewBmmDeclarationBuilder() *BmmDeclarationBuilder {
	return &BmmDeclarationBuilder{
		bmmdeclaration: NewBmmDeclaration(),
		errors:         make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
func (i *BmmDeclarationBuilder) SetName(v string) *BmmDeclarationBuilder {
	i.AddError(i.bmmdeclaration.SetName(v))
	return i
}
func (i *BmmDeclarationBuilder) SetResult(v IElWritableVariable) *BmmDeclarationBuilder {
	i.AddError(i.bmmdeclaration.SetResult(v))
	return i
}

// The declared type of the variable.
func (i *BmmDeclarationBuilder) SetType(v IBmmType) *BmmDeclarationBuilder {
	i.AddError(i.bmmdeclaration.SetType(v))
	return i
}

func (i *BmmDeclarationBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmDeclarationBuilder) Build() *BmmDeclaration {
	return i.bmmdeclaration
}

//FUNCTIONS
