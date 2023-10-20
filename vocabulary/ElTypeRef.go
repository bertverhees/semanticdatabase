package vocabulary

/**
Meta-type for reference to a non-abstract type as an object. Assumed to be
accessible at run-time. Typically represented syntactically as TypeName or
{TypeName} . May be used as a value, or as the qualifier for a function or
constant access.
*/

type IElTypeRef interface {
}

type ElTypeRef struct {
}
