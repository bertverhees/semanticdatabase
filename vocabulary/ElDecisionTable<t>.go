package vocabulary

/**
Meta-type for decision tables. Generic on the meta-type of the result attribute
of the branches, to allow specialised forms of if/else and case structures to be
created.
*/

type IElDecisionTable<t> interface {
}

type ElDecisionTable<t> struct {
}
