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
}

func NewBmmLocalBuilder() *BmmLocalBuilder {
	return &BmmLocalBuilder{
		bmmlocal: NewBmmLocal(),
	}
}

// BUILDER ATTRIBUTES
// From: BmmVariable
// Routine within which variable is defined.
func (i *BmmLocalBuilder) SetScope(v IBmmRoutine) *BmmLocalBuilder {
	i.bmmlocal.BmmModelElement.scope = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmLocalBuilder) SetType(v IBmmType) *BmmLocalBuilder {
	i.bmmlocal._type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmLocalBuilder) SetIsNullable(v bool) *BmmLocalBuilder {
	i.bmmlocal.isNullable = v
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmLocalBuilder) SetName(v string) *BmmLocalBuilder {
	i.bmmlocal.name = v
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
	i.bmmlocal.documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmLocalBuilder) SetExtensions(v map[string]any) *BmmLocalBuilder {
	i.bmmlocal.extensions = v
	return i
}

func (i *BmmLocalBuilder) Build() *BmmLocal {
	return i.bmmlocal
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmLocal) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmLocal) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmLocal) IsRootScope() bool {
	return false
}
