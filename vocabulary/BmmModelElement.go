package vocabulary

/**
Abstract meta-type of BMM declared model elements. A declaration is a an element
of a model within a context, which defines the scope of the element. Thus, a
class definition and its property and routine definitions are model elements,
but Types are not, since they are derived from model elements.
*/

// Interface definition
type IBmmModelElement interface {
	IsRootScope() bool
	Name() (string, error)
	SetName(name string) error
	Documentation() (map[string]any, error)
	SetDocumentation(v map[string]any) error
	Scope() (IBmmModelElement, error)
	SetScope(v IBmmModelElement) error
	Extensions() (map[string]any, error)
	SetExtensions(v map[string]any) error
}

// Struct definition
type BmmModelElement struct {
	// Attributes
	// name of this model element.
	name string `yaml:"name" json:"name" xml:"name"`
	/**
	Optional documentation of this element, as a keyed list. It is strongly
	recommended to use the following key /type combinations for the relevant
	purposes:
	"purpose": String
	"keywords": List<String>
	"use": String
	"misuse": String
	"references": String
	Other keys and value types may be freely added.
	*/
	documentation map[string]any `yaml:"documentation" json:"documentation" xml:"documentation"`
	// Model element within which an element is declared.
	scope IBmmModelElement `yaml:"scope" json:"scope" xml:"scope"`
	/**
	Optional meta-data of this element, as a keyed list. May be used to extend the
	meta-model.
	*/
	extensions map[string]any `yaml:"extensions" json:"extensions" xml:"extensions"`
}

// getters/setters
func (b *BmmModelElement) Name() (string, error) {
	return b.name, nil
}

func (b *BmmModelElement) SetName(name string) error {
	b.name = name
	return nil
}

func (b *BmmModelElement) Documentation() (map[string]any, error) {
	return b.documentation, nil
}

func (b *BmmModelElement) SetDocumentation(v map[string]any) error {
	b.documentation = v
	return nil
}

func (b *BmmModelElement) Scope() (IBmmModelElement, error) {
	return b.scope, nil
}

func (b *BmmModelElement) SetScope(v IBmmModelElement) error {
	b.scope = v
	return nil
}

func (b *BmmModelElement) Extensions() (map[string]any, error) {
	return b.extensions, nil
}

func (b *BmmModelElement) SetExtensions(v map[string]any) error {
	b.extensions = v
	return nil
}

// CONSTRUCTOR
// No constructor or Builder, class is abstract

//FUNCTIONS
/**
Post_result : Result = (scope = self). True if this model element is the root of
a model structure hierarchy.
*/
func (b *BmmModelElement) IsRootScope() bool {
	return false
}
