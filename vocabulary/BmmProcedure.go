package vocabulary

/**
A formal element with signature of the form: name ({arg:TArg}*):TStatus , where
TStatus is the built-in type BMM_STATUS_TYPE .. A procedure is a computed
(rather than data) element, generally assumed to be state-changing, and is
usually called in the form name ({arg:TArg}*) .
*/

// Interface definition
type IBmmProcedure interface {
	// From: BMM_MODEL_ELEMENT
	IBmmModelElement
	IBmmFormalElement
	IBmmFeature
	IBmmRoutine
	//BMM_ROUTINE
	Arity() int
}

// Struct definition
type BmmProcedure struct {
	// embedded for Inheritance
	BmmModelElement
	BmmFormalElement
	BmmFeature
	BmmRoutine
	// Constants
	// Attributes
	// Declared or inferred static type of the entity.
	Type IBmmStatusType `yaml:"type" json:"type" xml:"type"`
}

// CONSTRUCTOR
func NewBmmProcedure() *BmmProcedure {
	bmmprocedure := new(BmmProcedure)
	//BmmFormalElement
	//default, no constant
	bmmprocedure.IsNullable = false
	//BmmModelElement
	bmmprocedure.documentation = make(map[string]any)
	bmmprocedure.extensions = make(map[string]any)
	//BmmFeature
	bmmprocedure.FeatureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmRoutine
	bmmprocedure.Parameters = make([]IBmmParameter, 0)
	bmmprocedure.PreConditions = make([]IBmmAssertion, 0)
	bmmprocedure.PostConditions = make([]IBmmAssertion, 0)

	return bmmprocedure
}

// BUILDER
type BmmProcedureBuilder struct {
	bmmprocedure *BmmProcedure
}

func NewBmmProcedureBuilder() *BmmProcedureBuilder {
	return &BmmProcedureBuilder{
		bmmprocedure: NewBmmProcedure(),
	}
}

// BUILDER ATTRIBUTES
// Declared or inferred static type of the entity.
func (i *BmmProcedureBuilder) SetType(v IBmmStatusType) *BmmProcedureBuilder {
	i.bmmprocedure.Type = v
	return i
}

// From: BmmRoutine
// Formal parameters of the routine.
func (i *BmmProcedureBuilder) SetParameters(v []IBmmParameter) *BmmProcedureBuilder {
	i.bmmprocedure.Parameters = v
	return i
}

// From: BmmRoutine
/**
Boolean conditions that must evaluate to True for the routine to execute
correctly, May be used to generate exceptions if included in run-time build. A
False pre-condition implies an error in the passed parameters.
*/
func (i *BmmProcedureBuilder) SetPreConditions(v []IBmmAssertion) *BmmProcedureBuilder {
	i.bmmprocedure.PreConditions = v
	return i
}

// From: BmmRoutine
/**
Boolean conditions that will evaluate to True if the routine executed correctly,
May be used to generate exceptions if included in run-time build. A False
post-condition implies an error (i.e. bug) in routine code.
*/
func (i *BmmProcedureBuilder) SetPostConditions(v []IBmmAssertion) *BmmProcedureBuilder {
	i.bmmprocedure.PostConditions = v
	return i
}

// From: BmmRoutine
// Body of a routine, i.e. executable program.
func (i *BmmProcedureBuilder) SetDefinition(v IBmmRoutineDefinition) *BmmProcedureBuilder {
	i.bmmprocedure.Definition = v
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmProcedureBuilder) SetIsSynthesisedGeneric(v bool) *BmmProcedureBuilder {
	i.bmmprocedure.IsSynthesisedGeneric = v
	return i
}

// From: BmmFeature
// extensions to feature-level meta-types.
func (i *BmmProcedureBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmProcedureBuilder {
	i.bmmprocedure.FeatureExtensions = v
	return i
}

// From: BmmFeature
// Group containing this feature.
func (i *BmmProcedureBuilder) SetGroup(v IBmmFeatureGroup) *BmmProcedureBuilder {
	i.bmmprocedure.Group = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmProcedureBuilder) SetIsNullable(v bool) *BmmProcedureBuilder {
	i.bmmprocedure.IsNullable = v
	return i
}

// From: BmmModelElement
// name of this model element.
func (i *BmmProcedureBuilder) SetName(v string) *BmmProcedureBuilder {
	i.bmmprocedure.name = v
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
	i.bmmprocedure.documentation = v
	return i
}

// From: BmmModelElement
// Model element within which an element is declared.
func (i *BmmProcedureBuilder) SetScope(v IBmmModelElement) *BmmProcedureBuilder {
	i.bmmprocedure.BmmModelElement.scope = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmProcedureBuilder) SetExtensions(v map[string]any) *BmmProcedureBuilder {
	i.bmmprocedure.extensions = v
	return i
}

func (i *BmmProcedureBuilder) Build() *BmmProcedure {
	return i.bmmprocedure
}

// From: BMM_ROUTINE
// Return number of arguments of this routine.
func (b *BmmProcedure) Arity() int {
	return 0
}

// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmProcedure) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmProcedure) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmProcedure) IsRootScope() bool {
	return false
}
