package vocabulary

/*
Enumeration of parameter read/write direction values.
*/
type BmmParameterDirection int64

const (
	// Parameter is an input parameter, and treated as readonly by the receiving routine.
	in BmmParameterDirection = iota
	// Parameter is an output parameter, and treated as a reference to an entity writeable by the
	// receiving routine.
	out
	// Parameter is an input and output parameter, and treated as a reference to an entity readable
	// and writeable by the receiving routine.
	in_out
)
