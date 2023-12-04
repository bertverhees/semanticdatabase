package vocabulary

// A routine local variable (writable).

// Interface definition
type IBmmLocal interface {
	IBmmWritableVariable
	//BMM_LOCAL
}

// Struct definition
type BmmLocal struct {
	BmmWritableVariable
	// Attributes
}

// CONSTRUCTOR
func NewBmmLocal() *BmmLocal {
	bmmlocal := new(BmmLocal)
	//BmmModelElement
	bmmlocal.documentation = make(map[string]any)
	bmmlocal.extensions = make(map[string]any)
	//BmmFormalElement
	//default, no constant
	bmmlocal.isNullable = false

	return bmmlocal
}

// BUILDER
type BmmLocalBuilder struct {
	bmmlocal *BmmLocal
	errors   []error
}

func NewBmmLocalBuilder() *BmmLocalBuilder {
	return &BmmLocalBuilder{
		bmmlocal: NewBmmLocal(),
		errors:   make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// From: BmmVariable
// Routine within which variable is defined.
func (i *BmmLocalBuilder) SetScope(v IBmmRoutine) *BmmLocalBuilder {
	i.AddError(i.bmmlocal.SetScope(v))
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmLocalBuilder) SetIsNullable(v bool) *BmmLocalBuilder {
	i.AddError(i.bmmlocal.SetIsNullable(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmLocalBuilder) SetName(v string) *BmmLocalBuilder {
	i.AddError(i.bmmlocal.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmLocalBuilder) SetDocumentation(v map[string]any) *BmmLocalBuilder {
	i.AddError(i.bmmlocal.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmLocalBuilder) SetExtensions(v map[string]any) *BmmLocalBuilder {
	i.AddError(i.bmmlocal.SetExtensions(v))
	return i
}

func (i *BmmLocalBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmLocalBuilder) Build() *BmmLocal {
	return i.bmmlocal
}

//FUNCTIONS
