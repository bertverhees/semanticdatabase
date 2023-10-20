package vocabulary

/**
Definition of an enumeration class, understood as a class whose value range is
constrained extensionally, i.e. by an explicit enumeration of named singleton
instances. Only one inheritance ancestor is allowed in order to provide the base
type to which the range constraint is applied. The common notion of a set of
literals with no explicit defined values is represented as the degenerate
subtype BMM_ENUMERATION_INTEGER , whose values are 0, 1, …​
*/

type IBmmEnumeration interface {
}

type BmmEnumeration struct {
}
