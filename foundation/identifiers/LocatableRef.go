package identifiers

/* Reference to a LOCATABLE instance inside the top-level content structure inside a VERSION<T> identified by the id attribute.
The path attribute is applied to the object that VERSION.data points to.
*/

type LocatableRef struct {
	ObjectRef
}
