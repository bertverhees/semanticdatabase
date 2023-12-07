package vocabulary

import "semanticdatabase/base"

type IBmmDefinitions interface {
	base.IBasicDefinitions
	AnyClass() IBmmSimpleClass
	AnyType() IBmmSimpleType
	CreateSchemaId(a_model_publisher string, a_schema_name string, a_model_release string) string
}
