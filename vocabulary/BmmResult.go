package vocabulary

/**
Automatically declared variable representing result of a Function call
(writable).
*/

// Interface definition
type IBmmResult interface {
	IBmmWritableVariable
	//BMM_RESULT
}

// Struct definition
type BmmResult struct {
	BmmWritableVariable
}

// CONSTRUCTOR
func NewBmmResult() *BmmResult {
	bmmresult := new(BmmResult)
	bmmresult.name = "Result"
	//BmmModelElement
	bmmresult.documentation = make(map[string]any)
	bmmresult.extensions = make(map[string]any)
	//BmmFormalElement
	//default, no constant
	bmmresult.isNullable = false

	return bmmresult
}

// BUILDER
type BmmResultBuilder struct {
	bmmresult *BmmResult
	errors    []error
}

func NewBmmResultBuilder() *BmmResultBuilder {
	return &BmmResultBuilder{
		bmmresult: NewBmmResult(),
		errors:    make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// name of this model element.
// From: BmmVariable
// Routine within which variable is defined.
func (i *BmmResultBuilder) SetScope(v IBmmRoutine) *BmmResultBuilder {
	i.AddError(i.bmmresult.SetScope(v))
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmResultBuilder) SetIsNullable(v bool) *BmmResultBuilder {
	i.AddError(i.bmmresult.SetIsNullable(v))
	return i
}

// Declared or inferred static type of the entity.
func (i *BmmResultBuilder) SetType(v IBmmType) *BmmResultBuilder {
	i.AddError(i.bmmresult.SetType(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmResultBuilder) SetName(v string) *BmmResultBuilder {
	i.AddError(i.bmmresult.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmResultBuilder) SetDocumentation(v map[string]any) *BmmResultBuilder {
	i.AddError(i.bmmresult.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmResultBuilder) SetExtensions(v map[string]any) *BmmResultBuilder {
	i.AddError(i.bmmresult.SetExtensions(v))
	return i
}

func (i *BmmResultBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmResultBuilder) Build() *BmmResult {
	return i.bmmresult
}

//FUNCTIONS
