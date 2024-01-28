package primitives

/*
Abstract ancestor class for all other classes.
Usually maps to a type like Any or Object in an object-oriented technology.
Defined here to provide value and reference equality semantics.
*/
type Any interface {
	IsEqual(any Any) Boolean
	InstanceOf(aType String) Boolean
	TypeOf(anObject Any) String
	NotEqual(other Ordered) Boolean
}
