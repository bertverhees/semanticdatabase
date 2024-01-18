package vocabulary

import (
	"errors"
	"semanticdatabase/base"
)

/* ------------------- BmmFormalElement ----------------------- */
// A formal element having a name, type and a type-based signature.
type BmmFormalElement struct {
	BmmModelElement
	// Attributes
	// Declared or inferred static type of the entity.
	_type IBmmType `yaml:"type" json:"type" xml:"type"`
	/**
	True if this element can be null (Void) at execution ISO8601. should be interpreted as
	optionality in subtypes..
	*/
	isNullable bool `yaml:"isnullable" json:"isnullable" xml:"isnullable"`
}

func (b *BmmFormalElement) Type() IBmmType {
	return b._type
}

func (b *BmmFormalElement) SetType(_type IBmmType) error {
	if _type == nil {
		return errors.New("Type-property BmmFormalElement should not be set to nil")
	}
	b._type = _type
	return nil
}

func (b *BmmFormalElement) IsNullable() bool {
	return b.isNullable
}

func (b *BmmFormalElement) SetIsNullable(isNullable bool) error {
	b.isNullable = isNullable
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder

//FUNCTIONS
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmFormalElement) Signature() IBmmSignature {
	return nil
}

/*
*
Post_result : result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmFormalElement) IsBoolean() bool {
	return false
}

/* ------------------- BmmFeature -------------------------- */
// A module-scoped formal element.
type BmmFeature struct {
	BmmFormalElement
	// Constants
	// Attributes
	/**
	True if this feature was synthesised due to generic substitution in an inherited
	type, or further constraining of a formal generic parameter.
	*/
	isSynthesisedGeneric bool `yaml:"issynthesisedgeneric" json:"issynthesisedgeneric" xml:"issynthesisedgeneric"`
	// extensions to feature-level meta-types.
	featureExtensions []IBmmFeatureExtension `yaml:"featureextensions" json:"featureextensions" xml:"featureextensions"`
	// group containing this feature.
	group IBmmFeatureGroup `yaml:"group" json:"group" xml:"group"`
	// Model element within which an element is declared.
	scope IBmmClass `yaml:"scope" json:"scope" xml:"scope"`
}

func (b *BmmFeature) IsSynthesisedGeneric() bool {
	return b.isSynthesisedGeneric
}

func (b *BmmFeature) SetIsSynthesisedGeneric(isSynthesisedGeneric bool) error {
	b.isSynthesisedGeneric = isSynthesisedGeneric
	return nil
}

func (b *BmmFeature) FeatureExtensions() []IBmmFeatureExtension {
	return b.featureExtensions
}

func (b *BmmFeature) SetFeatureExtensions(featureExtensions []IBmmFeatureExtension) error {
	b.featureExtensions = featureExtensions
	return nil
}

func (b *BmmFeature) Group() IBmmFeatureGroup {
	return b.group
}

func (b *BmmFeature) SetGroup(group IBmmFeatureGroup) error {
	if group == nil {
		return errors.New("Group-property BmmFeature should not be set to nil")
	}
	b.group = group
	return nil
}

func (b *BmmFeature) SetScope(v IBmmModelElement) error {
	s, ok := v.(IBmmClass)
	if !ok {
		return errors.New("_type-assertion to IBmmClass in BmmFeature->SetScope failed")
	} else {
		b.scope = s
		return nil
	}
}

/* ------------------- BmmFeatureGroup ---------------------- */
/**
A logical group of features, with a name and set of properties that applies to
the group.
*/
type BmmFeatureGroup struct {
	// Attributes
	// Name of this feature group; defaults to 'feature'.
	name string `yaml:"name" json:"name" xml:"name"`
	/**
	Set of properties of this group, represented as name/value pairs. These are
	understood to apply logically to all of the features contained within the group.
	*/
	properties map[string]string `yaml:"properties" json:"properties" xml:"properties"`
	// Set of features in this group.
	features []IBmmFeature `yaml:"features" json:"features" xml:"features"`
	// Optional visibility to apply to all features in this group.
	visibility IBmmVisibility `yaml:"visibility" json:"visibility" xml:"visibility"`
}

func (b *BmmFeatureGroup) Name() string {
	return b.name
}

func (b *BmmFeatureGroup) SetName(name string) error {
	if name == "" {
		return errors.New("Name-property BmmFeatureGroup should not be set to empty")
	}
	b.name = name
	return nil
}

func (b *BmmFeatureGroup) Properties() map[string]string {
	return b.properties
}

func (b *BmmFeatureGroup) SetProperties(properties map[string]string) error {
	if properties == nil || len(properties) == 0 {
		return errors.New("Property-property BmmFeatureGroup should not be set to nil or empty")
	}
	b.properties = properties
	return nil
}

func (b *BmmFeatureGroup) Features() []IBmmFeature {
	return b.features
}

func (b *BmmFeatureGroup) SetFeatures(features []IBmmFeature) error {
	b.features = features
	return nil
}

func (b *BmmFeatureGroup) Visibility() IBmmVisibility {
	return b.visibility
}

func (b *BmmFeatureGroup) SetVisibility(visibility IBmmVisibility) error {
	b.visibility = visibility
	return nil
}

// CONSTRUCTOR
func NewBmmFeatureGroup() *BmmFeatureGroup {
	bmmfeaturegroup := new(BmmFeatureGroup)
	// Constants
	return bmmfeaturegroup
}

/* ------------------- BmmVisibility ---------------------- */
/**
Abstract parent of visibility representation. TODO: define schemes; probably
need to support C++/Java scheme as well as better type-based schemes.
*/
type BmmVisibility struct {
}

// CONSTRUCTOR
func NewBmmVisibility() *BmmVisibility {
	bmmvisibility := new(BmmVisibility)
	// Constants
	return bmmvisibility
}

/* ------------------- BmmInstantiableFeature ---------------------- */
/**
Meta-type representing instantiable features, i.e. features that are created as
value objects.
*/
type BmmInstantiableFeature struct {
	BmmFeature
}

/* ------------------- BmmStatic ---------------------- */
// Meta-type for static (i.e. read-only) properties.
type BmmStatic struct {
	BmmInstantiableFeature
	// Attributes
}

/* ------------------- BmmConstant ---------------------- */
/**
An immutable, static value-returning element scoped to a class. The value is the
result of the evaluation of the generator , which should be as simple as a literal
value, or should be any expression, including a function call.
*/
type BmmConstant struct {
	BmmStatic
	// Constants
	// Attributes
	// Literal value of the constant.
	generator IBmmLiteralValue[IBmmSimpleType] `yaml:"generator" json:"generator" xml:"generator"`
}

func (b *BmmConstant) Generator() IBmmLiteralValue[IBmmSimpleType] {
	return b.generator
}

func (b *BmmConstant) SetGenerator(generator IBmmLiteralValue[IBmmSimpleType]) error {
	if generator == nil {
		return errors.New("Generator-property BmmConstant should not be set to nil")
	}
	b.generator = generator
	return nil
}

// CONSTRUCTOR
func NewBmmConstant() *BmmConstant {
	bmmconstant := new(BmmConstant)
	//BmmFormalElement
	//default, no constant
	bmmconstant.isNullable = false
	//BmmModelElement
	bmmconstant.documentation = make(map[string]any)
	bmmconstant.extensions = make(map[string]any)
	//BmmFeature
	bmmconstant.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//BmmStatic
	return bmmconstant
}

/* ------------------- BmmSingleton ---------------------- */
// Meta-type for static value properties computed once by a function invocation.
type BmmSingleton struct {
	BmmStatic
	// Constants
	// Attributes
	// Generator of the value of this static property.
	generator IBmmRoutineDefinition `yaml:"generator" json:"generator" xml:"generator"`
}

func (b *BmmSingleton) Generator() IBmmRoutineDefinition {
	return b.generator
}

func (b *BmmSingleton) SetGenerator(generator IBmmRoutineDefinition) error {
	b.generator = generator
	return nil
}

// CONSTRUCTOR
func NewBmmSingleton() *BmmSingleton {
	bmmsingleton := new(BmmSingleton)
	//BmmFormalElement
	//default, no constant
	bmmsingleton.isNullable = false
	//BmmModelElement
	bmmsingleton.documentation = make(map[string]any)
	bmmsingleton.extensions = make(map[string]any)
	//BmmFeature
	bmmsingleton.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//BmmStatic

	return bmmsingleton
}

/* ------------------- BmmProperty ---------------------- */
/**
Meta-type of a writable property definition within a class definition of an
object model. The is_composition attribute indicates whether the property has
sub-part or an association semantics with respect to the owning class.
*/
type BmmProperty struct {
	BmmInstantiableFeature
	// Attributes
	// True if this property is marked with info model im_runtime property.
	isImRuntime bool `yaml:"isimruntime" json:"isimruntime" xml:"isimruntime"`
	// True if this property was marked with info model im_infrastructure flag.
	isImInfrastructure bool `yaml:"isiminfrastructure" json:"isiminfrastructure" xml:"isiminfrastructure"`
	/**
	True if this property instance is a compositional sub-part of the owning class
	instance. Equivalent to 'composition' in UML associations (but missing from UML
	properties without associations) and also 'cascade-delete' semantics in ER
	schemas.
	*/
	isComposition bool `yaml:"iscomposition" json:"iscomposition" xml:"iscomposition"`
}

func (b *BmmProperty) IsImRuntime() bool {
	return b.isImRuntime
}

func (b *BmmProperty) SetIsImRuntime(isImRuntime bool) error {
	b.isImRuntime = isImRuntime
	return nil
}

func (b *BmmProperty) IsImInfrastructure() bool {
	return b.isImInfrastructure
}

func (b *BmmProperty) SetIsImInfrastructure(isImInfrastructure bool) error {
	b.isImInfrastructure = isImInfrastructure
	return nil
}

func (b *BmmProperty) IsComposition() bool {
	return b.isComposition
}

func (b *BmmProperty) SetIsComposition(isComposition bool) error {
	b.isComposition = isComposition
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder

// FUNCTIONS
// Interval form of 0..1 , 1..1 etc, derived from is_nullable .
func (b *BmmProperty) Existence() *base.MultiplicityInterval[int] {
	return nil
}

// name of this property to display in UI.
func (b *BmmProperty) DisplayName() string {
	return ""
}

/* ------------------- BmmUnitaryProperty ---------------------- */
// Meta-type of for properties of unitary type.
type BmmUnitaryProperty struct {
	BmmProperty
	// Attributes
	// Declared or inferred static type of the entity.
}

func (b *BmmUnitaryProperty) SetType(_type IBmmType) error {
	s, ok := _type.(IBmmUnitaryType)
	if !ok {
		return errors.New("type-assertion in BmmUnitaryProperty->SetType failed")
	} else {
		b._type = s
		return nil
	}
}

// CONSTRUCTOR
func NewBmmUnitaryProperty() *BmmUnitaryProperty {
	bmmunitaryproperty := new(BmmUnitaryProperty)
	//BmmFormalElement
	//default, no constant
	bmmunitaryproperty.isNullable = false
	//BmmModelElement
	bmmunitaryproperty.documentation = make(map[string]any)
	bmmunitaryproperty.extensions = make(map[string]any)
	//BmmFeature
	bmmunitaryproperty.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//bmmProperty
	//default, no constant
	bmmunitaryproperty.isImRuntime = false
	//default, no constant
	bmmunitaryproperty.isImInfrastructure = false
	//default, no constant
	bmmunitaryproperty.isComposition = false
	return bmmunitaryproperty
}

/* ------------------- BmmContainerProperty ---------------------- */
// Meta-type of for properties of linear container type, such as List<T> etc.
type BmmContainerProperty struct {
	BmmProperty
	// Attributes
	// cardinality of this container.
	cardinality *base.MultiplicityInterval[int] `yaml:"cardinality" json:"cardinality" xml:"cardinality"`
	// Declared or inferred static type of the entity.
	_type IBmmContainerType `yaml:"type" json:"type" xml:"type"`
}

func (b *BmmContainerProperty) Cardinality() *base.MultiplicityInterval[int] {
	return b.cardinality
}

func (b *BmmContainerProperty) SetCardinality(cardinality *base.MultiplicityInterval[int]) error {
	b.cardinality = cardinality
	return nil
}

func (b *BmmContainerProperty) SetType(_type IBmmType) error {
	s, ok := _type.(IBmmContainerType)
	if !ok {
		return errors.New("type-assertion in BmmContainerProperty->SetType failed")
	} else {
		b._type = s
		return nil
	}
}

// CONSTRUCTOR
func NewBmmContainerProperty() *BmmContainerProperty {
	bmmcontainerproperty := new(BmmContainerProperty)
	//BmmFormalElement
	//default, no constant
	bmmcontainerproperty.isNullable = false
	//BmmModelElement
	bmmcontainerproperty.documentation = make(map[string]any)
	bmmcontainerproperty.extensions = make(map[string]any)
	//BmmFeature
	bmmcontainerproperty.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//bmmProperty
	//default, no constant
	bmmcontainerproperty.isImRuntime = false
	//default, no constant
	bmmcontainerproperty.isImInfrastructure = false
	//default, no constant
	bmmcontainerproperty.isComposition = false
	return bmmcontainerproperty
}

func (b *BmmContainerProperty) DisplayName() string {
	return ""
}

/* ------------------- BmmIndexedContainerProperty ---------------------- */
/**
Meta-type of for properties of linear container type, such as Hash<Index_type,
T> etc.
*/
type BmmIndexedContainerProperty struct {
	BmmContainerProperty
	// Attributes
	// Declared or inferred static type of the entity.
	_type IBmmIndexedContainerType `yaml:"type" json:"type" xml:"type"`
}

func (b *BmmIndexedContainerProperty) SetType(_type IBmmType) error {
	s, ok := _type.(IBmmIndexedContainerType)
	if !ok {
		return errors.New("type-assertion to IBmmIndexedContainerType in BmmIndexedContainerProperty->SetType failed")
	} else {
		b._type = s
		return nil
	}
}

// CONSTRUCTOR
func NewBmmIndexedContainerProperty() *BmmIndexedContainerProperty {
	bmmindexedcontainerproperty := new(BmmIndexedContainerProperty)
	//BmmFormalElement
	//default, no constant
	bmmindexedcontainerproperty.isNullable = false
	//BmmModelElement
	bmmindexedcontainerproperty.documentation = make(map[string]any)
	bmmindexedcontainerproperty.extensions = make(map[string]any)
	//BmmFeature
	bmmindexedcontainerproperty.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//bmmProperty
	//default, no constant
	bmmindexedcontainerproperty.isImRuntime = false
	//default, no constant
	bmmindexedcontainerproperty.isImInfrastructure = false
	//default, no constant
	bmmindexedcontainerproperty.isComposition = false
	return bmmindexedcontainerproperty
}

// name of this property in form name: ContainerTypeName<IndexTypeName, …​> .
func (b *BmmIndexedContainerProperty) DisplayName() string {
	return ""
}

/* ------------------- BmmRoutine ---------------------- */
// A feature defining a routine, scoped to a class.
type BmmRoutine struct {
	BmmFeature
	// Attributes
	// Formal parameters of the routine.
	parameters []IBmmParameter `yaml:"parameters" json:"parameters" xml:"parameters"`
	/**
	Boolean conditions that must evaluate to True for the routine to execute
	correctly, should be used to generate exceptions if included in run-ISO8601 build. A
	False pre-condition implies an error in the passed parameters.
	*/
	preConditions []IBmmAssertion `yaml:"preconditions" json:"preconditions" xml:"preconditions"`
	/**
	Boolean conditions that will evaluate to True if the routine executed correctly,
	should be used to generate exceptions if included in run-ISO8601 build. A False
	post-condition implies an error (i.e. bug) in routine code.
	*/
	postConditions []IBmmAssertion `yaml:"postconditions" json:"postconditions" xml:"postconditions"`
	// body of a routine, i.e. executable program.
	definition IBmmRoutineDefinition `yaml:"definition" json:"definition" xml:"definition"`
}

func (b *BmmRoutine) Parameters() []IBmmParameter {
	return b.parameters
}

func (b *BmmRoutine) SetParameters(parameters []IBmmParameter) error {
	b.parameters = parameters
	return nil
}

func (b *BmmRoutine) PreConditions() []IBmmAssertion {
	return b.preConditions
}

func (b *BmmRoutine) SetPreConditions(preConditions []IBmmAssertion) error {
	b.preConditions = preConditions
	return nil
}

func (b *BmmRoutine) PostConditions() []IBmmAssertion {
	return b.postConditions
}

func (b *BmmRoutine) SetPostConditions(postConditions []IBmmAssertion) error {
	b.postConditions = postConditions
	return nil
}

func (b *BmmRoutine) Definition() IBmmRoutineDefinition {
	return b.definition
}

func (b *BmmRoutine) SetDefinition(definition IBmmRoutineDefinition) error {
	b.definition = definition
	return nil
}

// CONSTRUCTOR
// abstract, no constructor, no builder

// FUNCTIONS
// Return number of arguments of this routine.
func (b *BmmRoutine) Arity() int {
	return 0
}

/* ------------------- BmmFunction ---------------------- */
/**
A formal element with signature of the form: name ({arg:TArg}*):TResult . A
function is a computed (rather than data) element, generally assumed to be
non-state-changing.
*/
type BmmFunction struct {
	BmmRoutine
	// Attributes
	/**
	Optional details enabling a function to be represented as an operator in a
	syntactic representation.
	*/
	operatorDefinition IBmmOperator `yaml:"operatordefinition" json:"operatordefinition" xml:"operatordefinition"`
	// Automatically created result variable, usable in body and post-condition.
	result IBmmResult `yaml:"result" json:"result" xml:"result"`
}

func (b *BmmFunction) OperatorDefinition() IBmmOperator {
	return b.operatorDefinition
}

func (b *BmmFunction) SetOperatorDefinition(operatorDefinition IBmmOperator) error {
	b.operatorDefinition = operatorDefinition
	return nil
}

func (b *BmmFunction) Result() IBmmResult {
	return b.result
}

func (b *BmmFunction) SetResult(result IBmmResult) error {
	if result == nil {
		return errors.New("result-property BmmFunction should not be set to nil")
	}
	b.result = result
	return nil
}

// CONSTRUCTOR
func NewBmmFunction() *BmmFunction {
	bmmfunction := new(BmmFunction)
	//BmmFormalElement
	//default, no constant
	bmmfunction.isNullable = false
	//BmmModelElement
	bmmfunction.documentation = make(map[string]any)
	bmmfunction.extensions = make(map[string]any)
	//BmmFeature
	bmmfunction.featureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmRoutine
	bmmfunction.parameters = make([]IBmmParameter, 0)
	bmmfunction.preConditions = make([]IBmmAssertion, 0)
	bmmfunction.postConditions = make([]IBmmAssertion, 0)

	return bmmfunction
}

/* ------------------- BmmOperator ---------------------- */
// definition of a symbolic operator associated with a function.
type BmmOperator struct {
	// embedded for Inheritance
	// Constants
	// Attributes
	// position of operator in syntactic representation.
	position BmmOperatorPosition `yaml:"position" json:"position" xml:"position"`
	/**
	Set of String symbols that should be used to represent this operator in a textual
	representation of a BMM model.
	*/
	symbols []string `yaml:"symbols" json:"symbols" xml:"symbols"`
	// Formal name of the operator, e.g. 'minus' etc.
	name string `yaml:"name" json:"name" xml:"name"`
}

func (b *BmmOperator) Position() *BmmOperatorPosition {
	return &b.position
}

func (b *BmmOperator) SetPosition(position *BmmOperatorPosition) error {
	b.position = *position
	return nil
}

func (b *BmmOperator) Symbols() []string {
	return b.symbols
}

func (b *BmmOperator) SetSymbols(symbols []string) error {
	b.symbols = symbols
	return nil
}

func (b *BmmOperator) Name() string {
	return b.name
}

func (b *BmmOperator) SetName(name string) error {
	b.name = name
	return nil
}

// CONSTRUCTOR
func NewBmmOperator() *BmmOperator {
	bmmoperator := new(BmmOperator)
	bmmoperator.symbols = make([]string, 0)
	return bmmoperator
}

/* ------------------- BmmOperatorPosition ---------------------- */
/*
Enumeration of possible position of operator in a syntactic representation
for operators associated with 1- and 2- degree functions
*/
type BmmOperatorPosition int64

const (
	// Prefix operator position: operator comes before operand.
	prefix BmmOperatorPosition = iota
	// Infix operator position: operator comes between left and right operands.
	infix
)

/* ------------------- BmmProcedure ---------------------- */
/**
A formal element with signature of the form: name ({arg:TArg}*):TStatus , where
TStatus is the built-in type BMM_STATUS_TYPE .. A procedure is a computed
(rather than data) element, generally assumed to be state-changing, and is
usually called in the form name ({arg:TArg}*) .
*/
type BmmProcedure struct {
	BmmRoutine
	// Attributes
	// Declared or inferred static type of the entity.
	_type IBmmStatusType `yaml:"type" json:"type" xml:"type"`
}

func (b *BmmProcedure) SetType(_type IBmmType) error {
	s, ok := _type.(IBmmStatusType)
	if !ok {
		return errors.New("_type-assertion in BmmProcedure->SetType failed")
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

/*
*
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmProcedure) Signature() IBmmSignature {
	return nil
}

/* ------------------- BmmVariable ---------------------- */
// A routine-scoped formal element.
type BmmVariable struct {
	BmmFormalElement
	// Attributes
	// Routine within which variable is defined.
	scope IBmmRoutine `yaml:"scope" json:"scope" xml:"scope"`
}

func (b *BmmVariable) SetScope(scope IBmmModelElement) error {
	if scope == nil {
		return errors.New("Scope-property BmmVariable should not be set to nil")
	}
	s, ok := scope.(IBmmRoutine)
	if !ok {
		return errors.New("_type-assertion in BmmVariable->SetScope failed")
	} else {
		b.scope = s
		return nil
	}
}

/* ------------------- BmmWritableVariable ---------------------- */
// Meta-type for writable variables, including the special variable result .
type BmmWritableVariable struct {
	BmmVariable
	// Attributes
}

func NewBmmWritableVariable() *BmmWritableVariable {
	bmmwritable := new(BmmWritableVariable)
	bmmwritable.documentation = make(map[string]any)
	bmmwritable.extensions = make(map[string]any)
	return bmmwritable
}

/* ------------------- BmmLocal ---------------------- */
// A routine local variable (writable).
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

/* ------------------- BmmResult ---------------------- */
/**
Automatically declared variable representing result of a Function call
(writable).
*/
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

/* ------------------- BmmReadonlyVariable ---------------------- */
/**
Meta-type for writable variables, including routine parameters and the special
variable Self .
*/
type BmmReadonlyVariable struct {
	BmmVariable
	// Attributes
}

func NewBmmReadonlyVariable() *BmmReadonlyVariable {
	bmmReadonly := new(BmmReadonlyVariable)
	bmmReadonly.documentation = make(map[string]any)
	bmmReadonly.extensions = make(map[string]any)

	return bmmReadonly
}

/* ------------------- BmmSelf ---------------------- */
/**
Meta-type for an automatically created variable referencing the current
instance. Typically called 'self' or 'this' in programming languages. Read-only.
*/
type BmmSelf struct {
	BmmReadonlyVariable
	// Constants
	// Attributes
	// name of this model element.
}

// CONSTRUCTOR
func NewBmmSelf() *BmmSelf {
	bmmself := new(BmmSelf)
	//BmmModelElement
	bmmself.documentation = make(map[string]any)
	bmmself.extensions = make(map[string]any)
	//BmmFormalElement
	//default, no constant
	bmmself.isNullable = false
	//default, no constant
	bmmself.name = "Self"
	return bmmself
}

/* ------------------- BmmParameter ---------------------- */
// A routine parameter variable (read-only).
type BmmParameter struct {
	// embedded for Inheritance
	BmmReadonlyVariable
	// Constants
	// Attributes
	/**
	Optional read/write direction of the parameter. If none-supplied, the parameter
	is treated as in , i.e. readable.
	*/
	direction BmmParameterDirection `yaml:"direction" json:"direction" xml:"direction"`
}

func (b *BmmParameter) Direction() *BmmParameterDirection {
	return &b.direction
}

func (b *BmmParameter) SetDirection(direction *BmmParameterDirection) error {
	b.direction = *direction
	return nil
}

// CONSTRUCTOR
func NewBmmParameter() *BmmParameter {
	bmmparameter := new(BmmParameter)
	//BmmModelElement
	bmmparameter.documentation = make(map[string]any)
	bmmparameter.extensions = make(map[string]any)
	//BmmFormalElement
	//default, no constant
	bmmparameter.isNullable = false
	return bmmparameter
}

/* ------------------- BmmParameterDirection ---------------------- */
/*
Enumeration of parameter read/write direction values.
*/
type BmmParameterDirection int64

const (
	// Parameter is an input parameter, and treated as readonly by the receiving routine.
	in BmmParameterDirection = iota
	// Parameter is an output parameter, and treated as a reference to an entity writeable by the
	// receiving routine.
	out
	// Parameter is an input and output parameter, and treated as a reference to an entity readable
	// and writeable by the receiving routine.
	in_out
)

/* ------------------- BmmRoutineDefinition ---------------------- */
// Abstract ancestor of routine body meta-types.
type BmmRoutineDefinition struct {
	// embedded for Inheritance
	// Constants
	// Attributes
}

/* ------------------- BmmLocalRoutine ---------------------- */
// Meta-type for locally declared routine body.
type BmmLocalRoutine struct {
	// embedded for Inheritance
	BmmRoutineDefinition
	// Constants
	// Attributes
	// Local variables of the routine, if there is a body defined.
	locals []IBmmLocal `yaml:"locals" json:"locals" xml:"locals"`
	// body of routine declaration.
	body IBmmStatementBlock `yaml:"body" json:"body" xml:"body"`
}

func (b *BmmLocalRoutine) Locals() []IBmmLocal {
	return b.locals
}

func (b *BmmLocalRoutine) SetLocals(locals []IBmmLocal) error {
	b.locals = locals
	return nil
}

func (b *BmmLocalRoutine) Body() IBmmStatementBlock {
	return b.body
}

func (b *BmmLocalRoutine) SetBody(body IBmmStatementBlock) error {
	if body == nil {
		return errors.New("Body property of BmmLocalRoutine should not be set nil")
	}
	b.body = body
	return nil
}

// CONSTRUCTOR
func NewBmmLocalRoutine() *BmmLocalRoutine {
	bmmlocalroutine := new(BmmLocalRoutine)
	//BmmLocal
	bmmlocalroutine.locals = make([]IBmmLocal, 0)
	return bmmlocalroutine
}

/* ------------------- BmmExternalRoutine ---------------------- */
/**
External routine meta-type, containing sufficient meta-data to enable a routine
in an external library to be called.
*/
type BmmExternalRoutine struct {
	BmmRoutineDefinition
	// Attributes
	/**
	External call general meta-data, including target routine name, type mapping
	etc.
	*/
	metaData map[string]string `yaml:"metadata" json:"metadata" xml:"metadata"`
	// Optional argument-mapping meta-data.
	argumentMapping map[string]string `yaml:"argumentmapping" json:"argumentmapping" xml:"argumentmapping"`
}

func (b *BmmExternalRoutine) MetaData() map[string]string {
	return b.metaData
}

func (b *BmmExternalRoutine) SetMetaData(metaData map[string]string) error {
	if metaData == nil || len(metaData) == 0 {
		return errors.New("metaData-property BmmExternalRoutine should not be set to nil or empty")
	}
	b.metaData = metaData
	return nil
}

func (b *BmmExternalRoutine) ArgumentMapping() map[string]string {
	return b.argumentMapping
}

func (b *BmmExternalRoutine) SetArgumentMapping(argumentMapping map[string]string) error {
	b.argumentMapping = argumentMapping
	return nil
}

// CONSTRUCTOR
func NewBmmExternalRoutine() *BmmExternalRoutine {
	bmmexternalroutine := new(BmmExternalRoutine)
	bmmexternalroutine.metaData = make(map[string]string)
	bmmexternalroutine.argumentMapping = make(map[string]string)
	return bmmexternalroutine
}

/* ------------------- BmmFeatureExtension ---------------------- */
type BmmFeatureExtension struct {
	// embedded for Inheritance
	// Constants
	// Attributes
}

// CONSTRUCTOR
func NewBmmFeatureExtension() *BmmFeatureExtension {
	bmmFeatureExtension := new(BmmFeatureExtension)
	// Constants
	return bmmFeatureExtension
}
