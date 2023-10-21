package vocabulary

/**
	Abstract meta-type of BMM declared model elements. A declaration is a an element
	of a model within a context, which defines the scope of the element. Thus, a
	class definition and its property and routine definitions are model elements,
	but Types are not, since they are derived from model elements.
*/

type IBmmModelElement interface {
// True if this model element is the root of a model structure hierarchy.
	is_root_scope (): Boolean  Post_result : Result = (scope = self) (  )  Boolean  Post_result : Result = (scope = self)
}

type BmmModelElement struct {
}
