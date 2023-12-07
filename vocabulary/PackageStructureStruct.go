package vocabulary

import (
	"semanticdatabase/base"
	"strings"
)

/* ----------  BmmDefinitions  -----------*/
type BmmDefinitions struct {
	base.BasicDefinitions
}

func NewBmmDefinitions() *BmmDefinitions {
	bmmdefinitions := new(BmmDefinitions)
	return bmmdefinitions
}

// FUNCTIONS
// While a BMM model defines a model in terms of class declarations, it must always have a top class named Any from which all others inherit.
// Similar to the root class in a typical OOPL type systems (sometimes called 'Object'), the Any class defines semantics true for all objects such as equality
// (i.e. semantics for an '=' operator) and copying.
//
// A BMM model may define its own Any class, but if it does not, the BMM_MODEL instance representing the model will produce a standard 'Any' class via the
// any_class_definition() method. This will create the following structure, including a default package structure, and an Any type.
// built-in class definition corresponding to the top `Any' class.
func (b *BmmDefinitions) AnyClass() IBmmSimpleClass {
	//return NewBmmSimpleClassBuilder().SetName("Any").SetPackage(NewBmmPackageBuilder().SetName("foundation_types").Build()).Build()
	return nil
}

// Built-in type definition corresponding to the top `Any' type.
func (b *BmmDefinitions) AnyType() IBmmSimpleType {
	return NewBmmSimpleTypeBuilder().SetBaseClass(b.AnyClass()).Build()
}

/*
Create schema id, formed from: a_model_publisher '_' a_schema_name '_' a_model_release e.g. openehr_rm_1.0.3 , openehr_test_1.0.1 , iso_13606_1_2008_2.1.2 .
*/
func (b *BmmDefinitions) CreateSchemaId(aModelPublisher string, aSchemaName string, aModelRelease string) string {
	result := strings.Builder{}
	if strings.TrimSpace(aModelPublisher) == "" {
		result.WriteString("unknown")
	} else {
		result.WriteString(aModelPublisher)
	}
	result.WriteString("_")
	if strings.TrimSpace(aSchemaName) == "" {
		result.WriteString("unknown")
	} else {
		result.WriteString(aSchemaName)
	}
	result.WriteString("_")
	if strings.TrimSpace(aModelRelease) == "" {
		result.WriteString("unknown")
	} else {
		result.WriteString(aModelRelease)
	}
	return strings.ToLower(result.String())
}
