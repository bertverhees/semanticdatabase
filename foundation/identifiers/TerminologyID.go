package identifiers

/* Identifier for terminologies such as accessed via a terminology query service.
In this class, the value attribute identifies the Terminology in the terminology service, e.g. SNOMED-CT .
A terminology is assumed to be in a particular language, which must be explicitly specified.

Lexical form: name [ '(' version ')' ].

*/

type TerminologyID struct {
	ObjectID
}
