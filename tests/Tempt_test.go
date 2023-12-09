package tests

import (
	"fmt"
	"semanticdatabase/vocabulary"
	"testing"
)

func TestBmmSchemaDescriptorBuilder(t *testing.T) {
	test, e := vocabulary.NewBmmSchemaDescriptorBuilder().SetSchemaId("test").Build()
	if e==nil {
		fmt.Println(test.SchemaId())
	}

}
