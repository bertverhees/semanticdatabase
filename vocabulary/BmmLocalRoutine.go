package vocabulary

// Meta-type for locally declared routine body.

// Interface definition
type IBmmLocalRoutine interface {
	IBmmRoutineDefinition
	//BMM_LOCAL_ROUTINE
	Locals() []IBmmLocal
	SetLocals(locals []IBmmLocal) error
	Body() IBmmStatementBlock
	SetBody(body IBmmStatementBlock) error
}

// Struct definition
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

// BUILDER
type BmmLocalRoutineBuilder struct {
	bmmlocalroutine *BmmLocalRoutine
	errors          []error
}

func NewBmmLocalRoutineBuilder() *BmmLocalRoutineBuilder {
	return &BmmLocalRoutineBuilder{
		bmmlocalroutine: NewBmmLocalRoutine(),
		errors:          make([]error, 0),
	}
}

// BUILDER ATTRIBUTES
// Local variables of the routine, if there is a body defined.
func (i *BmmLocalRoutineBuilder) SetLocals(v []IBmmLocal) *BmmLocalRoutineBuilder {
	i.AddError(i.bmmlocalroutine.SetLocals(v))
	return i
}

// body of routine declaration.
func (i *BmmLocalRoutineBuilder) SetBody(v IBmmStatementBlock) *BmmLocalRoutineBuilder {
	i.AddError(i.bmmlocalroutine.SetBody(v))
	return i
}

func (i *BmmLocalRoutineBuilder) AddError(e error) {
	if e != nil {
		i.errors = append(i.errors, e)
	}
}

func (i *BmmLocalRoutineBuilder) Build() *BmmLocalRoutine {
	return i.bmmlocalroutine
}

//FUNCTIONS
