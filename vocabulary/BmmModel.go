package vocabulary

/**
	Definition of the root of a BMM model (along with what is inherited from
	BMM_SCHEMA_CORE ).
*/

type IBmmModel interface {
	/**
		Identifier of this model, lower-case, formed from:
		<rm_publisher>_<model_name>_<rm_release> E.g. "openehr_ehr_1.0.4" .
	*/
	model_id (): String (  )  String
	/**
		Retrieve the class definition corresponding to a_type_name (which may contain a
		generic part).
	*/
	class_definition ( a_name: String[1] ): BMM_CLASS ( a_name String[1] )  BMM_CLASS
	/**
		Retrieve the class definition corresponding to a_type_name . If it contains a
		generic part, this will be removed if it is a fully open generic name, otherwise
		it will remain intact, i.e. if it is an effective generic name that identifies a
		BMM_GENERIC_CLASS_EFFECTIVE .
	*/
	type_definition (): BMM_CLASS (  )  BMM_CLASS
// True if a_class_name has a class definition in the model.
	has_class_definition ( a_class_name: String[1] ): Boolean ( a_class_name String[1] )  Boolean
	/**
		True if a_type_name is already concretely known in the system, including if it
		is generic, which may be open, partially open or closed.
	*/
	has_type_definition ( a_type_name: String[1] ): Boolean ( a_type_name String[1] )  Boolean
// Retrieve the enumeration definition corresponding to a_type_name .
	enumeration_definition ( a_name: String[1] ): BMM_ENUMERATION ( a_name String[1] )  BMM_ENUMERATION
// List of keys in class_definitions of items marked as primitive types.
	primitive_types (): List <String> (  )  List <String>
// List of keys in class_definitions of items that are enumeration types.
	enumeration_types (): List <String> (  )  List <String>
	/**
		Retrieve the property definition for a_prop_name in flattened class
		corresponding to a_type_name .
	*/
	property_definition (): BMM_PROPERTY (  )  BMM_PROPERTY
	/**
		True if a_ms_property_type is a valid 'MS' dynamic type for a_property in BMM
		type a_bmm_type_name . 'MS' conformance means 'model-semantic' conformance,
		which abstracts away container types like List<> , Set<> etc and compares the
		dynamic type with the relation target type in the UML sense, i.e. regardless of
		whether there is single or multiple containment.
	*/
	ms_conformant_property_type ( a_bmm_type_name: String[1] , a_bmm_property_name: String[1] , a_ms_property_name: String[1] ): Boolean ( a_bmm_type_name String[1], a_bmm_property_name String[1], a_ms_property_name String[1] )  Boolean
	/**
		Retrieve the property definition for a_property_path in flattened class
		corresponding to a_type_name .
	*/
	property_definition_at_path (): BMM_PROPERTY (  )  BMM_PROPERTY
	/**
		Retrieve the class definition for the class that owns the terminal attribute in
		a_prop_path in flattened class corresponding to a_type_name .
	*/
	class_definition_at_path ( a_type_name: String[1] , a_prop_path: String[1] ): BMM_CLASS ( a_type_name String[1], a_prop_path String[1] )  BMM_CLASS
	/**
		Return all ancestor types of a_class_name up to root class (usually Any , Object
		or something similar). Does not include current class. Returns empty list if
		none.
	*/
	all_ancestor_classes ( a_class: String[1] ): List <String> ( a_class String[1] )  List <String>
// True if a_class_name is a descendant in the model of a_parent_class_name .
	is_descendant_of ( a_class_name: String[1] , a_parent_class_name: String[1] ): Boolean ( a_class_name String[1], a_parent_class_name String[1] )  Boolean
	/**
		Check conformance of a_desc_type to an_anc_type ; the types may be generic, and
		may contain open generic parameters like 'T' etc. These are replaced with their
		appropriate constrainer types, or Any during the conformance testing process.
		Conformance is found if: [base class test] types are non-generic, and either
		type names are identical, or else a_desc_type has an_anc_type in its ancestors;
		both types are generic and pass base class test; number of generic params
		matches, and each generic parameter type, after 'open parameter' substitution,
		recursively passes; type_name_conforms_to test descendant type is generic and
		ancestor type is not, and they pass base classes test.
	*/
	type_conforms_to ( a_desc_type: String[1] , an_anc_type: String[1] ): Boolean ( a_desc_type String[1], an_anc_type String[1] )  Boolean
	/**
		Generate type substitutions for the supplied type, which may be simple, generic
		(closed, open or partially open), or a container type. In the generic and
		container cases, the result is the permutation of the base class type and type
		substitutions of all generic parameters. Parameters a_type Name of a type.
	*/
	subtypes ( a_type: String[1] ): List <String> ( a_type String[1] )  List <String>
	/**
		BMM_SIMPLE_CLASS instance for the Any class. This may be defined in the BMM
		schema, but if not, use BMM_DEFINITIONS. any_class .
	*/
	any_class_definition (): BMM_SIMPLE_CLASS (  )  BMM_SIMPLE_CLASS
// BMM_SIMPLE_TYPE instance for the Any type.
	any_type_definition (): BMM_SIMPLE_TYPE (  )  BMM_SIMPLE_TYPE
// BMM_SIMPLE_TYPE instance for the Boolean type.
	boolean_type_definition (): BMM_SIMPLE_TYPE (  )  BMM_SIMPLE_TYPE
}

type BmmModel struct {
}
