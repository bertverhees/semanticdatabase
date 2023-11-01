package vocabulary

/**
Meta-type defining a generalised module concept. Descendants define actual
structure and contents.
*/

// Interface definition
type IBmmModule interface {
	// From: BMM_MODEL_ELEMENT
	IsRootScope() bool
}

// Struct definition
type BmmModule struct {
	// embedded for Inheritance
	BmmModelElement
	// Constants
	// Attributes
	// List of feature groups in this class.
	FeatureGroups []IBmmFeatureGroup `yaml:"featuregroups" json:"featuregroups" xml:"featuregroups"`
	// Features of this module.
	Features []IBmmFormalElement `yaml:"features" json:"features" xml:"features"`
	// Model within which module is defined.
	Scope IBmmModel `yaml:"scope" json:"scope" xml:"scope"`
}

// CONSTRUCTOR
func NewBmmModule() *BmmModule {
	bmmmodule := new(BmmModule)
	bmmmodule.Features := make([]IBmmFormalElement,0)
	bmmmodule.FeatureGroups := make([]IBmmFeatureGroup,0)
	bmmmodule.Documentation := make(map[string]any)
	bmmmodule.Extensions := make(map[string]any)
	return bmmmodule
}

// BUILDER
type BmmModuleBuilder struct {
	bmmmodule *BmmModule
}

func NewBmmModuleBuilder() *BmmModuleBuilder {
	return &BmmModuleBuilder{
		bmmmodule: NewBmmModule(),
	}
}

// BUILDER ATTRIBUTES
// List of feature groups in this class.
func (i *BmmModuleBuilder) SetFeatureGroups(v []IBmmFeatureGroup) *BmmModuleBuilder {
	i.bmmmodule.FeatureGroups = v
	return i
}

// Features of this module.
func (i *BmmModuleBuilder) SetFeatures(v []IBmmFormalElement) *BmmModuleBuilder {
	i.bmmmodule.Features = v
	return i
}

// Model within which module is defined.
func (i *BmmModuleBuilder) SetScope(v IBmmModel) *BmmModuleBuilder {
	i.bmmmodule.Scope = v
	return i
}

// From: BmmModelElement
// Name of this model element.
func (i *BmmModuleBuilder) SetName(v string) *BmmModuleBuilder {
	i.bmmmodule.Name = v
	return i
}

// From: BmmModelElement
/**
Optional documentation of this element, as a keyed list. It is strongly
recommended to use the following key /type combinations for the relevant
purposes: "purpose": String "keywords": List<String> "use": String "misuse":
String "references": String Other keys and value types may be freely added.
*/
func (i *BmmModuleBuilder) SetDocumentation(v map[string]any) *BmmModuleBuilder {
	i.bmmmodule.Documentation = v
	return i
}

// From: BmmModelElement
/**
Optional meta-data of this element, as a keyed list. May be used to extend the
meta-model.
*/
func (i *BmmModuleBuilder) SetExtensions(v map[string]any) *BmmModuleBuilder {
	i.bmmmodule.Extensions = v
	return i
}

func (i *BmmModuleBuilder) Build() *BmmModule {
	return i.bmmmodule
}

//FUNCTIONS
// From: BMM_MODEL_ELEMENT
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmModule) IsRootScope() bool {
	return false
}
