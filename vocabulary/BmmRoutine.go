package vocabulary

// A feature defining a routine, scoped to a class.

type IBmmRoutine interface {
// Return number of arguments of this routine.
	arity (): Integer (  )  Integer
}

type BmmRoutine struct {
}
