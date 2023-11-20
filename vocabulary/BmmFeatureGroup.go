package vocabulary

/**
A logical group of features, with a name and set of properties that applies to
the group.
*/

// Interface definition
type IBmmFeatureGroup interface {
	Name() string
	SetName(name string) error
	Properties() map[string]string
	SetProperties(properties map[string]string) error
	Features() []IBmmFeature
	SetFeatures(features []IBmmFeature) error
	Visibility() IBmmVisibility
	SetVisibility(visibility IBmmVisibility) error
}

// Struct definition
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
	b.name = name
	return nil
}

func (b *BmmFeatureGroup) Properties() map[string]string {
	return b.properties
}

func (b *BmmFeatureGroup) SetProperties(properties map[string]string) error {
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

// BUILDER
type BmmFeatureGroupBuilder struct {
	bmmfeaturegroup *BmmFeatureGroup
	errors          []error
}

func NewBmmFeatureGroupBuilder() *BmmFeatureGroupBuilder {
	return &BmmFeatureGroupBuilder{
		bmmfeaturegroup: NewBmmFeatureGroup(),
		errors:          make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// name of this feature group; defaults to 'feature'.
func (i *BmmFeatureGroupBuilder) SetName(v string) *BmmFeatureGroupBuilder {
	i.AddError(i.bmmfeaturegroup.SetName(v))
	return i
}

/*
*
Set of properties of this group, represented as name/value pairs. These are
understood to apply logically to all of the features contained within the group.
*/
func (i *BmmFeatureGroupBuilder) SetProperties(v map[string]string) *BmmFeatureGroupBuilder {
	i.AddError(i.bmmfeaturegroup.SetProperties(v))
	return i
}

// Set of features in this group.
func (i *BmmFeatureGroupBuilder) SetFeatures(v []IBmmFeature) *BmmFeatureGroupBuilder {
	i.AddError(i.bmmfeaturegroup.SetFeatures(v))
	return i
}

// Optional visibility to apply to all features in this group.
func (i *BmmFeatureGroupBuilder) SetVisibility(v IBmmVisibility) *BmmFeatureGroupBuilder {
	i.AddError(i.bmmfeaturegroup.SetVisibility(v))
	return i
}

func (i *BmmFeatureGroupBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmFeatureGroupBuilder) Build() *BmmFeatureGroup {
	return i.bmmfeaturegroup
}

//FUNCTIONS
