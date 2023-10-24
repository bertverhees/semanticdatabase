package vocabulary

import (
	"vocabulary"
)

/**
	A delayed routine call, whose arguments may be open, partially closed or closed.
	Generated by special reference to a routine, usually via a dedicated keyword,
	such as 'lambda' or 'agent', or other special syntax. Instances may include
	closed delayed calls like calculate_age (dob="1987-09-13", today="2019-06-03")
	but also partially open calls such as format_structure (struct=?, style=3) ,
	where struct is an open argument. Evaluation type (i.e. type of runtime
	evaluated form) is BMM_SIGNATURE .
*/

// Interface definition
type IElAgent interface {
	EvalType (  )  IBmmRoutineType
	IsCallable (  )  Boolean  Post_result_validity : Result = open_arguments = Void
	Reference (  )  string
	// From: EL_FEATURE_REF
	Reference (  )  string
	// From: EL_VALUE_GENERATOR
	Reference (  )  string
	// From: EL_SIMPLE
	// From: EL_TERMINAL
	// From: EL_EXPRESSION
	EvalType (  )  IBmmType
	IsBoolean (  )  bool
}

// Struct definition
type ElAgent struct {
	// embedded for Inheritance
	ElFeatureRef
	ElValueGenerator
	ElSimple
	ElTerminal
	ElExpression
	// Constants
	// Attributes
	// Closed arguments of a routine call as a tuple of objects.
	ClosedArgs	IElTuple	`yaml:"closedargs" json:"closedargs" xml:"closedargs"`
	/**
		Optional list of names of open arguments of the call. If not provided, and the
		name refers to a routine with more arguments than supplied in closed_args , the
		missing arguments are inferred from the definition .
	*/
	OpenArgs	[]string	`yaml:"openargs" json:"openargs" xml:"openargs"`
	// Reference to definition of a routine for which this is an agent, if one exists.
	Definition	IBmmRoutine	`yaml:"definition" json:"definition" xml:"definition"`
	// Name of the routine being called.
	Name	string	`yaml:"name" json:"name" xml:"name"`
	IsWritable	bool	`yaml:"iswritable" json:"iswritable" xml:"iswritable"`
}

//CONSTRUCTOR
func NewElAgent() *ElAgent {
	elagent := new(ElAgent)
	// Constants
	// From: ElFeatureRef
	// From: ElValueGenerator
	// From: ElSimple
	// From: ElTerminal
	// From: ElExpression
	return elagent
}
//BUILDER
type ElAgentBuilder struct {
	elagent *ElAgent
}

func NewElAgentBuilder() *ElAgentBuilder {
	 return &ElAgentBuilder {
		elagent : NewElAgent(),
	}
}

//BUILDER ATTRIBUTES
	// Closed arguments of a routine call as a tuple of objects.
func (i *ElAgentBuilder) SetClosedArgs ( v IElTuple ) *ElAgentBuilder{
	i.elagent.ClosedArgs = v
	return i
}
	/**
		Optional list of names of open arguments of the call. If not provided, and the
		name refers to a routine with more arguments than supplied in closed_args , the
		missing arguments are inferred from the definition .
	*/
func (i *ElAgentBuilder) SetOpenArgs ( v []string ) *ElAgentBuilder{
	i.elagent.OpenArgs = v
	return i
}
	// Reference to definition of a routine for which this is an agent, if one exists.
func (i *ElAgentBuilder) SetDefinition ( v IBmmRoutine ) *ElAgentBuilder{
	i.elagent.Definition = v
	return i
}
	// Name of the routine being called.
func (i *ElAgentBuilder) SetName ( v string ) *ElAgentBuilder{
	i.elagent.Name = v
	return i
}
func (i *ElAgentBuilder) SetIsWritable ( v bool ) *ElAgentBuilder{
	i.elagent.IsWritable = v
	return i
}

func (i *ElAgentBuilder) Build() *ElAgent {
	 return i.elagent
}

//FUNCTIONS
/**
	Post_result : Result = definition.signature. Eval type is the signature
	corresponding to the (remaining) open arguments and return type, if any.
*/
func (e *ElAgent) EvalType (  )  IBmmRoutineType {
	return nil
}
// True if there are no open arguments.
func (e *ElAgent) IsCallable (  )  Boolean  Post_result_validity : Result = open_arguments = Void {
	return nil
}
// Generated full reference name, including scoping elements.
func (e *ElAgent) Reference (  )  string {
	return nil
}
// From: EL_FEATURE_REF
/**
	Generated full reference name, consisting of scoping elements and name
	concatenated using dot notation.
*/
func (e *ElAgent) Reference (  )  string {
	return nil
}
// From: EL_VALUE_GENERATOR
/**
	Generated full reference name, based on constituent parts of the entity. Default
	version outputs name field.
*/
func (e *ElAgent) Reference (  )  string {
	return nil
}
// From: EL_EXPRESSION
/**
	Meta-type of expression entity used in type-checking and evaluation. Effected in
	descendants.
*/
func (e *ElAgent) EvalType (  )  IBmmType {
	return nil
}
// From: EL_EXPRESSION
/**
	Post_result : Result = eval_type().equal(
	{BMM_MODEL}.boolean_type_definition()). True if eval_type is notionally Boolean
	(i.e. a BMM_SIMPLE_TYPE with type_name() = Boolean ).
*/
func (e *ElAgent) IsBoolean (  )  bool {
	return nil
}
