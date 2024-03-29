package identifiers

/* Model of the DCE Universal Unique Identifier or UUID which takes the form of hexadecimal integers separated by hyphens,
following the pattern 8-4-4-4-12 as defined by the Open Group, CDE 1.1 Remote Procedure Call specification,
Appendix A. Also known as a GUID.
*/

type UUID struct {
	UID
}
