package identifiers

/* Generic identifier type for identifiers whose format is otherwise unknown to openEHR.
Includes an attribute for naming the identification scheme (which may well be local).
*/

type GenericID struct {
	ObjectID
}
