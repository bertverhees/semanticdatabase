package vocabulary

/**
Compound expression consisting of a list of value-range / expression pairs, and
an else member that as a whole, represents a case statement flavour of decision
table. Evaluated by iterating through items and for each one, comparing input to
the item value_range . If the input is in the range, the evaluation result of
the table is that item’s result evaluation result. If no member of items has a
True-returning condition , the evaluation result is the result of evaluating the
else expression.
*/

type ElCaseTable<t>er interface {
}

type ElCaseTable<t> struct {
}
