package vocabulary

/* ================ BmmStatementItem =========================== */
type BmmStatementItemBuilder struct {
	Builder
}

/* ================ BmmStatementBlock =========================== */
type BmmStatementBlockBuilder struct {
	BmmStatementItemBuilder
}

func NewBmmStatementBlockBuilder() *BmmStatementBlockBuilder {
	builder := &BmmStatementBlockBuilder{}
	builder.object = NewBmmStatementBlock()
	return builder
}

func (i *BmmStatementBlockBuilder) SetItems(v []IBmmStatementItem) *BmmStatementBlockBuilder {
	i.AddError(i.object.(*BmmStatementBlock).SetItems(v))
	return i
}

func (i *BmmStatementBlockBuilder) Build() (*BmmStatementBlock, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmStatementBlock), nil
	}
}

/* ================ BmmStatement =========================== */
type BmmStatementBuilder struct {
	BmmStatementItemBuilder
}

/* ================ BmmSimpleStatement =========================== */
type BmmSimpleStatementBuilder struct {
	BmmStatementBuilder
}

/* ================ BmmDeclaration =========================== */
type BmmDeclarationBuilder struct {
	BmmSimpleStatementBuilder
}

func NewBmmDeclarationBuilder() *BmmDeclarationBuilder {
	builder := &BmmDeclarationBuilder{}
	builder.object = NewBmmDeclaration()
	return builder
}

// BUILDER ATTRIBUTES
func (i *BmmDeclarationBuilder) SetName(v string) *BmmDeclarationBuilder {
	i.AddError(i.object.(*BmmDeclaration).SetName(v))
	return i
}
func (i *BmmDeclarationBuilder) SetResult(v IElWritableVariable) *BmmDeclarationBuilder {
	i.AddError(i.object.(*BmmDeclaration).SetResult(v))
	return i
}

// The declared type of the variable.
func (i *BmmDeclarationBuilder) SetType(v IBmmType) *BmmDeclarationBuilder {
	i.AddError(i.object.(*BmmDeclaration).SetType(v))
	return i
}

func (i *BmmDeclarationBuilder) Build() (*BmmDeclaration, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmDeclaration), nil
	}
}

/* ================ BmmAssignment =========================== */
type BmmAssignmentBuilder struct {
	BmmSimpleStatementBuilder
}

func NewBmmAssignmentBuilder() *BmmAssignmentBuilder {
	builder := &BmmAssignmentBuilder{}
	builder.object = NewBmmAssignment()
	return builder
}

// BUILDER ATTRIBUTES
// The target variable on the notional left-hand side of this assignment.
func (i *BmmAssignmentBuilder) SetTarget(v IElValueGenerator) *BmmAssignmentBuilder {
	i.AddError(i.object.(*BmmAssignment).SetTarget(v))
	return i
}

// source right hand side) of the assignment.
func (i *BmmAssignmentBuilder) SetSource(v IElExpression) *BmmAssignmentBuilder {
	i.AddError(i.object.(*BmmAssignment).SetSource(v))
	return i
}

func (i *BmmAssignmentBuilder) Build() (*BmmAssignment, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmAssignment), nil
	}
}

/* ================ BmmProcedureCall =========================== */
type BmmProcedureCallBuilder struct {
	ElAgentCallBuilder
	BmmSimpleStatementBuilder
}

func NewBmmProcedureCallBuilder() *BmmProcedureCallBuilder {
	builder := &BmmProcedureCallBuilder{}
	builder.object = NewBmmProcedureCall()
	return builder
}

// BUILDER ATTRIBUTES
// The procedure agent being called.
func (i *BmmProcedureCallBuilder) SetAgent(v IElProcedureAgent) *BmmProcedureCallBuilder {
	i.AddError(i.object.(*BmmProcedureCall).SetAgent(v))
	return i
}

func (i *BmmProcedureCallBuilder) Build() (*BmmProcedureCall, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmProcedureCall), nil
	}
}

/* ================ BmmAssertion =========================== */
type BmmAssertionBuilder struct {
	BmmSimpleStatementBuilder
}

func NewBmmAssertionBuilder() *BmmAssertionBuilder {
	builder := &BmmAssertionBuilder{}
	builder.object = NewBmmAssertion()
	return builder
}

// BUILDER ATTRIBUTES
// Boolean-valued expression of the assertion.
func (i *BmmAssertionBuilder) SetExpression(v IElBooleanExpression) *BmmAssertionBuilder {
	i.AddError(i.object.(*BmmAssertion).SetExpression(v))
	return i
}

/*
*
Optional tag, typically used to designate design intention of the assertion,
e.g. "Inv_all_members_valid" .
*/
func (i *BmmAssertionBuilder) SetTag(v string) *BmmAssertionBuilder {
	i.AddError(i.object.(*BmmAssertion).SetTag(v))
	return i
}

func (i *BmmAssertionBuilder) Build() (*BmmAssertion, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmAssertion), nil
	}
}

/* ================ BmmActionTable =========================== */
type BmmActionTableBuilder struct {
	BmmStatementBuilder
}

func NewBmmActionTableBuilder() *BmmActionTableBuilder {
	builder := &BmmActionTableBuilder{}
	builder.object = NewBmmActionTable()
	return builder
}

//BUILDER ATTRIBUTES
/**
A specialised decision table whose outputs can only be procedure agents. In
execution, the matched agent will be invoked.
*/
func (i *BmmActionTableBuilder) SetDecisionTable(v IBmmActionDecisionTable) *BmmActionTableBuilder {
	i.AddError(i.object.(*BmmActionTable).SetDecisionTable(v))
	return i
}

func (i *BmmActionTableBuilder) Build() (*BmmActionTable, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmActionTable), nil
	}
}

/* ================ BmmActionDecisionTable =========================== */
type BmmActionDecisionTableBuilder struct {
	Builder
}

func NewBmmActionDecisionTableBuilder() *BmmActionDecisionTableBuilder {
	builder := &BmmActionDecisionTableBuilder{}
	builder.object = NewBmmActionDecisionTable()
	return builder
}

func (i *BmmActionDecisionTableBuilder) Build() (*BmmActionDecisionTable, []error) {
	if len(i.errors) > 0 {
		return nil, i.errors
	} else {
		return i.object.(*BmmActionDecisionTable), nil
	}
}
