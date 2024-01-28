package identifiers

/* Model of a reverse internet domain, as used to uniquely identify an internet domain.
In the form of a dot-separated string in the reverse order of a domain name, specified by IETF RFC 1034.
*/

type InternetID struct {
	UID
}
