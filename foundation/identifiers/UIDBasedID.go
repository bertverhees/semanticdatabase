package identifiers

/* Abstract model of UID-based identifiers consisting of a root part and an optional extension; lexical form: root '::' extension.
 */


type UIDBasedID struct {
	ObjectID
}
