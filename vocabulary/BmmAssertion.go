package vocabulary

/**
A statement that asserts the truth of its expression. If the expression
evaluates to False the execution generates an exception (depending on run-time
settings). May be rendered in syntax as assert condition or similar.
*/

type IBmmAssertion interface {
}

type BmmAssertion struct {
}
