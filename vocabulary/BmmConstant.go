package vocabulary

/**
An immutable, static value-returning element scoped to a class. The value is the
result of the evaluation of the generator , which may be as simple as a literal
value, or may be any expression, including a function call.
*/

// Interface definition
type IBmmConstant interface {
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
	// BMM_FORMAL_ELEMENT
	Signature() IBmmSignature
	IsBoolean() bool
	//BMM_FEATURE
	//BMM_INSTANTIABLE_FEATURE
	//BMM_STATIC
	//BMM_CONSTANT
}

// Struct definition
type BmmConstant struct {
	// embedded for Inheritance
	BmmModelElement
	BmmFormalElement
	BmmFeature
	BmmInstantiableFeature
	BmmStatic
	// Constants
	// Attributes
	// Literal value of the constant.
	Generator IBmmLiteralValue[IBmmSimpleType] `yaml:"generator" json:"generator" xml:"generator"`
}

// CONSTRUCTOR
func NewBmmConstant() *BmmConstant {
	bmmconstant := new(BmmConstant)
	//BmmFormalElement
	//default, no constant
	bmmconstant.IsNullable = false
	//BmmModelElement
	bmmconstant.Documentation = make(map[string]any)
	bmmconstant.Extensions = make(map[string]any)
	//BmmFeature
	bmmconstant.FeatureExtensions = make([]IBmmFeatureExtension, 0)
	//BmmInstantiableFeature
	//BmmStatic
	return bmmconstant
}

// BUILDER
type BmmConstantBuilder struct {
	bmmconstant *BmmConstant
}

func NewBmmConstantBuilder() *BmmConstantBuilder {
	return &BmmConstantBuilder{
		bmmconstant: NewBmmConstant(),
	}
}

// BUILDER ATTRIBUTES
// Literal value of the constant.
func (i *BmmConstantBuilder) SetGenerator(v IBmmLiteralValue[IBmmSimpleType]) *BmmConstantBuilder {
	i.bmmconstant.Generator = v
	return i
}

// From: BmmFeature
/**
True if this feature was synthesised due to generic substitution in an inherited
type, or further constraining of a formal generic parameter.
*/
func (i *BmmConstantBuilder) SetIsSynthesisedGeneric(v bool) *BmmConstantBuilder {
	i.bmmconstant.IsSynthesisedGeneric = v
	return i
}

// From: BmmFeature
// Extensions to feature-level meta-types.
func (i *BmmConstantBuilder) SetFeatureExtensions(v []IBmmFeatureExtension) *BmmConstantBuilder {
	i.bmmconstant.FeatureExtensions = v
	return i
}

// From: BmmFeature
// Group containing this feature.
func (i *BmmConstantBuilder) SetGroup(v IBmmFeatureGroup) *BmmConstantBuilder {
	i.bmmconstant.Group = v
	return i
}

// From: BmmFeature
// Model element within which an element is declared.
func (i *BmmConstantBuilder) SetScope(v IBmmClass) *BmmConstantBuilder {
	i.bmmconstant.BmmFeature.Scope = v
	return i
}

// From: BmmFormalElement
// Declared or inferred static type of the entity.
func (i *BmmConstantBuilder) SetType(v IBmmType) *BmmConstantBuilder {
	i.bmmconstant.Type = v
	return i
}

// From: BmmFormalElement
/**
True if this element can be null (Void) at execution time. May be interpreted as
optionality in subtypes..
*/
func (i *BmmConstantBuilder) SetIsNullable(v bool) *BmmConstantBuilder {
	i.bmmconstant.IsNullable = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmConstantBuilder) SetName(v string) *BmmConstantBuilder {
	i.bmmconstant.Name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmConstantBuilder) SetDocumentation(v map[string]any) *BmmConstantBuilder {
	i.bmmconstant.Documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmConstantBuilder) SetExtensions(v map[string]any) *BmmConstantBuilder {
	i.bmmconstant.Extensions = v
	return i
}

func (i *BmmConstantBuilder) Build() *BmmConstant {
	return i.bmmconstant
}

//FUNCTIONS
// From: BMM_FORMAL_ELEMENT
/**
Formal signature of this element, in the form: name [arg1_name: T_arg1,
…​][:T_value] Specific implementations in descendants.
*/
func (b *BmmConstant) Signature() IBmmSignature {
	return nil
}

// From: BMM_FORMAL_ELEMENT
/**
Post_result : Result = type().equal( {BMM_MODEL}.boolean_type_definition()).
True if type is notionally Boolean (i.e. a BMM_SIMPLE_TYPE with type_name() =
'Boolean' ).
*/
func (b *BmmConstant) IsBoolean() bool {
	return false
}

// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmConstant) IsRootScope() bool {
	return false
}
