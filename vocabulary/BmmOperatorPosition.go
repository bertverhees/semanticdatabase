package vocabulary

/*
Enumeration of possible position of operator in a syntactic representation
for operators associated with 1- and 2- degree functions
*/
type BmmOperatorPosition int64

const (
	// Prefix operator position: operator comes before operand.
	prefix BmmOperatorPosition = iota
	// Infix operator position: operator comes between left and right operands.
	infix
)
