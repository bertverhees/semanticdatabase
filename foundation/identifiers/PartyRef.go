package identifiers

/* Identifier for parties in a demographic or identity service.
There are typically a number of subtypes of the PARTY class, including PERSON, ORGANISATION, etc.
Abstract supertypes are allowed if the referenced object is of a type not known by the current implementation of this class
(in other words, if the demographic model is changed by the addition of a new PARTY or ACTOR subtypes,
valid PARTY_REFs can still be constructed to them).
*/

type PartyRef struct {
	ObjectRef
}
