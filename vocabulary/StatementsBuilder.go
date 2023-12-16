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
/* ================ BmmAssertion =========================== */
/* ================ BmmActionTable =========================== */
/* ================ BmmActionDecisionTable =========================== */
