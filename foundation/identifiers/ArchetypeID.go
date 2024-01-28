package identifiers

/* Identifier for archetypes. Ideally these would identify globally unique archetypes.

Lexical form: rm_originator '-' rm_name '-' rm_entity '.' concept_name { '-' specialisation }* '.v' number.
*/

type ArchetypeID struct {
	ObjectID
}
