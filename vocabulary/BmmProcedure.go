package vocabulary

import "errors"

/**
A formal element with signature of the form: name ({arg:TArg}*):TStatus , where
TStatus is the built-in type BMM_STATUS_TYPE .. A procedure is a computed
(rather than data) element, generally assumed to be state-changing, and is
usually called in the form name ({arg:TArg}*) .
*/

// Interface definition
type IBmmProcedure interface {
	IBmmRoutine
}

// Struct definition
type BmmProcedure struct {
	BmmRoutine
	// Attributes
	// Declared or inferred static type of the entity.
	_type IBmmStatusType `yaml:"type" json:"type" xml:"type"`
}

func (b *BmmProcedure) SetType(_type IBmmType) error {
	s, ok := _type.(IBmmStatusType)
	if !ok {
		return errors.New("_type-assertion in BmmProcedure->SetType went wrong")
	} else {
		b._type = s
		return nil
	}
}

// CONSTRUCTOR
func NewBmmProcedure() *BmmProcedure {
	bmmprocedure := new(BmmProcedure)
	//BmmFormalElement
	//default, no constant
	bmmprocedure.isNullable = false
	//BmmModelElement
	bmmprocedure.documentation = make(map[string]any)
	bmmprocedure.extensions = make(map[string]any)
	//BmmFeature
	bmmprocedure.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmRoutine
	bmmprocedure.parameters = make([]IBmmParameter, 0)
	bmmprocedure.preConditions = make([]IBmmAssertion, 0)
	bmmprocedure.postConditions = make([]IBmmAssertion, 0)

	return bmmprocedure
}

// BUILDER
type BmmProcedureBuilder struct {
	bmmprocedure *BmmProcedure
	errors       []error
}

func NewBmmProcedureBuilder() *BmmProcedureBuilder {
	return &BmmProcedureBuilder{
		bmmprocedure: NewBmmProcedure(),
		errors:       make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Declared or inferred static type of the entity.
func (i *BmmProcedureBuilder) SetType(v IBmmStatusType) *BmmProcedureBuilder {
	i.bmmprocedure._type = v
	return i
}

// From: BmmRoutine
// Formal parameters of the routine.
func (i *BmmProcedureBuilder) SetParameters(v []IBmmParameter) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetParameters(v))
	return i
}

// From: BmmRoutine
/**
Boolean conditions that must evaluate to True for the routine to execute
correctly, May be used to generate exceptions if included in run-time build. A
False pre-condition implies an error in the passed parameters.
*/
func (i *BmmProcedureBuilder) SetPreConditions(v []IBmmAssertion) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetPreConditions(v))
	return i
}

// From: BmmRoutine
/**
Boolean conditions that will evaluate to True if the routine executed correctly,
May be used to generate exceptions if included in run-time build. A False
post-condition implies an error (i.e. bug) in routine code.
*/
func (i *BmmProcedureBuilder) SetPostConditions(v []IBmmAssertion) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetPostConditions(v))
	return i
}

// From: BmmRoutine
// Body of a routine, i.e. executable program.
func (i *BmmProcedureBuilder) SetDefinition(v IBmmRoutineDefinition) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetDefinition(v))
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmProcedureBuilder) SetIsSynthesisedGeneric(v bool) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetIsSynthesisedGeneric(v))
	return i
}

// From: BmmFeature
// extensions to feature-level meta-types.
func (i *BmmProcedureBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetFeatureExtensions(v))
	return i
}

// From: BmmFeature
// group containing this feature.
func (i *BmmProcedureBuilder) SetGroup(v IBmmFeatureGroup) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetGroup(v))
	return i
}

// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmProcedureBuilder) SetScope(v IBmmClass) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetScope(v))
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmProcedureBuilder) SetIsNullable(v bool) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetIsNullable(v))
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmProcedureBuilder) SetName(v string) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetName(v))
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmProcedureBuilder) SetDocumentation(v map[string]any) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetDocumentation(v))
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmProcedureBuilder) SetExtensions(v map[string]any) *BmmProcedureBuilder {
	i.AddError(i.bmmprocedure.SetExtensions(v))
	return i
}

func (i *BmmProcedureBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmProcedureBuilder) Build() *BmmProcedure {
	return i.bmmprocedure
}

// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmProcedure) Signature() IBmmSignature {
	return nil
}
