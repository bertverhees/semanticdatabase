package vocabulary

// Persistent form of BMM_CONTAINER_TYPE .

// Interface definition

// Struct definition

// BUILDER

//FUNCTIONS
/**
The target type; this converts to the first parameter in generic_parameters in
BMM_GENERIC_TYPE . Persisted attribute.
*/
func (p *PBmmContainerType) TypeRef() IPBmmBaseType {
	return nil
}

// From: P_BMM_TYPE
// Create appropriate BMM_XXX object; effected in descendants.
func (p *PBmmContainerType) CreateBmmType(a_schema IBmmModel, a_class_def IBmmClass) {
	return
}

// From: P_BMM_TYPE
// Formal name of the type for display.
func (p *PBmmContainerType) AsTypeString() string {
	return ""
}
