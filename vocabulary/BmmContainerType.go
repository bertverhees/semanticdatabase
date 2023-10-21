package vocabulary

/**
	Meta-type that specifies linear containers with a generic parameter
	corresponding to the type of contained item, and whose container type is a
	generic type such as List<T> , Set<T> etc.
*/

type IBmmContainerType interface {
	TypeName():String (  )  string
	IsAbstract():BooleanPostIsAbstract:Result=ContainerType.isAbstract (  )  Boolean  Post_is_abstract : Result = container_type.is_abstract
	FlattenedTypeList():List<string>PostResult:Result=ItemType.flattenedTypeList (  )  List <String>  Post_result : Result = item_type.flattened_type_list
	UnitaryType():BmmUnitaryType (  )  BMM_UNITARY_TYPE
	IsPrimitive():BooleanPostResult:Result=ItemType.isPrimitive (  )  Boolean  Post_result : Result = item_type.is_primitive
	EffectiveType():BmmEffectiveType (  )  BMM_EFFECTIVE_TYPE
}

type BmmContainerType struct {
}
