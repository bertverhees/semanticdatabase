package v2

// Persistent form of BMM_MODEL_ELEMENT .

type IPBmmModelElement interface {
}

type PBmmModelElement struct {
	// Optional documentation of this element.
	Documentation string `yaml:"documentation" json:"documentation" xml:"documentation"`
}
