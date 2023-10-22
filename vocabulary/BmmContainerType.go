package vocabulary

/**
	Meta-type that specifies linear containers with a generic parameter
	corresponding to the type of contained item, and whose container type is a
	generic type such as List<T> , Set<T> etc.
*/

type IBmmContainerType interface {
	TypeName (  )  string
	IsAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract
	FlattenedTypeList (  )  List <String>  Post_result : Result = item_type.flattened_type_list
	UnitaryType (  )  IBmmUnitaryType
	IsPrimitive (  )  Boolean  Post_result : Result = item_type.is_primitive
	EffectiveType (  )  IBmmEffectiveType
}

type BmmContainerType struct {
}
