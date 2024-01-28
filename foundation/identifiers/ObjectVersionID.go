package identifiers

/* Globally unique identifier for one version of a versioned object; lexical form:
object_id '::' creating_system_id '::' version_tree_id.
*/

type ObjectVersionID struct {
	UIDBasedID
}
