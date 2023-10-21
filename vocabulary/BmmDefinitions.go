package vocabulary

// Definitions used by all BMM packages.

type IBmmDefinitions interface {
// built-in class definition corresponding to the top `Any' class.
	Any_class (): BMM_SIMPLE_CLASS (  )  BMM_SIMPLE_CLASS
// Built-in type definition corresponding to the top `Any' type.
	Any_type (): BMM_SIMPLE_TYPE (  )  BMM_SIMPLE_TYPE
	/**
		Create schema id, formed from: a_model_publisher '-' a_schema_name '-'
		a_model_release e.g. openehr_rm_1.0.3 , openehr_test_1.0.1 ,
		iso_13606_1_2008_2.1.2 .
	*/
	create_schema_id ( a_model_publisher, a_schema_name, a_model_release: String[1] ): String ( a_model_publisher String[1], a_schema_name String[1], a_model_release String[1] )  String
}

type BmmDefinitions struct {
}
