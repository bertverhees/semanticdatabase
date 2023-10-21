package vocabulary

// A feature defining a routine, scoped to a class.

type IBmmRoutine interface {
	Arity():Integer (  )  Integer
}

type BmmRoutine struct {
}
