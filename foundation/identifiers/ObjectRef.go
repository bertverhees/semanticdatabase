package identifiers

/* Class describing a reference to another object, which may exist locally or be maintained outside the current namespace,
e.g. in another service. Services are usually external, e.g. available in a LAN (including on the same host)
or the internet via Corba, SOAP, or some other distributed protocol.
However, in small systems they may be part of the same executable as the data containing the Id.
*/

type ObjectRef struct{}
