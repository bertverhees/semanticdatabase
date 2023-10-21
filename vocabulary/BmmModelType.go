package vocabulary

// A type that is defined by a class (or classes) in the model.

type IBmmModelType interface {
	TypeBaseName():String (  )  string
	IsPrimitive():Boolean (  )  Boolean
}

type BmmModelType struct {
}
