package vocabulary

/**
External routine meta-type, containing sufficient meta-data to enable a routine
in an external library to be called.
*/

// Interface definition
type IBmmExternalRoutine interface {
	IBmmRoutineDefinition
	MetaData() map[string]string
	SetMetaData(metaData map[string]string) error
	ArgumentMapping() map[string]string
	SetArgumentMapping(argumentMapping map[string]string) error
}

// Struct definition
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

// BUILDER
type BmmExternalRoutineBuilder struct {
	bmmexternalroutine *BmmExternalRoutine
	errors             []error
}

func NewBmmExternalRoutineBuilder() *BmmExternalRoutineBuilder {
	return &BmmExternalRoutineBuilder{
		bmmexternalroutine: NewBmmExternalRoutine(),
		errors:             make([]error, 0),
	}
}

//BUILDER ATTRIBUTES
/**
External call general meta-data, including target routine name, type mapping
etc.
*/
func (i *BmmExternalRoutineBuilder) SetMetaData(v map[string]string) *BmmExternalRoutineBuilder {
	i.AddError(i.bmmexternalroutine.SetMetaData(v))
	return i
}

// Optional argument-mapping meta-data.
func (i *BmmExternalRoutineBuilder) SetArgumentMapping(v map[string]string) *BmmExternalRoutineBuilder {
	i.AddError(i.bmmexternalroutine.SetArgumentMapping(v))
	return i
}

func (i *BmmExternalRoutineBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmExternalRoutineBuilder) Build() *BmmExternalRoutine {
	return i.bmmexternalroutine
}

//FUNCTIONS
