package vocabulary

import (
	"reflect"
	"testing"
)

func TestBmmEnumerationBuilder_Build(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   *BmmEnumeration
		want1  []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			got, got1 := i.Build()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Build() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetAncestors(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmModelType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetAncestors(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAncestors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetConverters(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetConverters(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetConverters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetCreators(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetCreators(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCreators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetDocumentation(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetDocumentation(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDocumentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetExtensions(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetExtensions(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExtensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetFeatureGroups(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmFeatureGroup
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFeatureGroups(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFeatureGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetFeatures(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmFormalElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFeatures(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFeatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetFunctions(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmFunction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFunctions(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFunctions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetImmediateDescendants(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmClass
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetImmediateDescendants(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetImmediateDescendants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetInvariants(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmAssertion
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetInvariants(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInvariants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetIsAbstract(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsAbstract(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsAbstract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetIsOverride(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsOverride(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsOverride() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetIsPrimitive(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsPrimitive(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsPrimitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetItemNames(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetItemNames(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetItemNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetItemValues(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmPrimitiveValue
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetItemValues(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetItemValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetName(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetName(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetPackage(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v IBmmPackage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetPackage(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetProcedures(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetProcedures(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProcedures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetProperties(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProperty
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetProperties(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetScope(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v IBmmModelElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetScope(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetSourceSchemaId(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetSourceSchemaId(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSourceSchemaId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationBuilder_SetStaticProperties(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmStatic
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetStaticProperties(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStaticProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_Build(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   *BmmEnumerationInteger
		want1  []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			got, got1 := i.Build()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Build() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetAncestors(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmModelType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetAncestors(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAncestors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetConverters(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetConverters(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetConverters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetCreators(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetCreators(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCreators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetDocumentation(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetDocumentation(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDocumentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetExtensions(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetExtensions(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExtensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetFeatureGroups(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmFeatureGroup
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFeatureGroups(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFeatureGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetFeatures(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmFormalElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFeatures(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFeatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetFunctions(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmFunction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFunctions(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFunctions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetImmediateDescendants(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmClass
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetImmediateDescendants(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetImmediateDescendants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetInvariants(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmAssertion
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetInvariants(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInvariants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetIsAbstract(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsAbstract(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsAbstract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetIsOverride(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsOverride(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsOverride() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetIsPrimitive(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsPrimitive(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsPrimitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetItemValues(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmPrimitiveValue
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetItemValues(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetItemValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetName(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetName(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetPackage(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v IBmmPackage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetPackage(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetProcedures(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetProcedures(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProcedures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetProperties(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProperty
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetProperties(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetScope(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v IBmmModelElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetScope(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetSourceSchemaId(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetSourceSchemaId(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSourceSchemaId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationIntegerBuilder_SetStaticProperties(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmStatic
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationIntegerBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetStaticProperties(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStaticProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_Build(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   *BmmEnumerationString
		want1  []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			got, got1 := i.Build()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Build() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetAncestors(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmModelType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetAncestors(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAncestors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetConverters(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetConverters(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetConverters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetCreators(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetCreators(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCreators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetDocumentation(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetDocumentation(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDocumentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetExtensions(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetExtensions(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExtensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetFeatureGroups(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmFeatureGroup
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFeatureGroups(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFeatureGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetFeatures(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmFormalElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFeatures(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFeatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetFunctions(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmFunction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFunctions(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFunctions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetImmediateDescendants(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmClass
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetImmediateDescendants(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetImmediateDescendants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetInvariants(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmAssertion
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetInvariants(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInvariants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetIsAbstract(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsAbstract(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsAbstract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetIsOverride(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsOverride(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsOverride() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetIsPrimitive(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsPrimitive(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsPrimitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetItemValues(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmPrimitiveValue
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetItemValues(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetItemValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetName(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetName(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetPackage(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v IBmmPackage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetPackage(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetProcedures(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetProcedures(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProcedures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetProperties(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProperty
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetProperties(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetScope(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v IBmmModelElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetScope(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetSourceSchemaId(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetSourceSchemaId(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSourceSchemaId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmEnumerationStringBuilder_SetStaticProperties(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmStatic
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmEnumerationStringBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetStaticProperties(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStaticProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_Build(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   *BmmGenericClass
		want1  []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			got, got1 := i.Build()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Build() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetAncestors(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmModelType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetAncestors(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAncestors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetConverters(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetConverters(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetConverters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetCreators(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetCreators(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCreators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetDocumentation(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetDocumentation(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDocumentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetExtensions(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetExtensions(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExtensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetFeatureGroups(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmFeatureGroup
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFeatureGroups(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFeatureGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetFeatures(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmFormalElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFeatures(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFeatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetFunctions(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmFunction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFunctions(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFunctions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetGenericParameters(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmParameterType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetGenericParameters(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetGenericParameters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetImmediateDescendants(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmClass
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetImmediateDescendants(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetImmediateDescendants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetInvariants(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmAssertion
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetInvariants(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInvariants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetIsAbstract(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsAbstract(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsAbstract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetIsOverride(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsOverride(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsOverride() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetIsPrimitive(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsPrimitive(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsPrimitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetName(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetName(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetPackage(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v IBmmPackage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetPackage(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetProcedures(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetProcedures(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProcedures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetProperties(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProperty
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetProperties(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetScope(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v IBmmModelElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetScope(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetSourceSchemaId(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetSourceSchemaId(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSourceSchemaId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmGenericClassBuilder_SetStaticProperties(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmStatic
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmGenericClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetStaticProperties(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStaticProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_Build(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   *BmmSimpleClass
		want1  []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			got, got1 := i.Build()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Build() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetAncestors(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmModelType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetAncestors(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAncestors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetConverters(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetConverters(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetConverters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetCreators(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetCreators(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCreators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetDocumentation(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetDocumentation(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDocumentation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetExtensions(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetExtensions(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExtensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetFeatureGroups(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmFeatureGroup
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFeatureGroups(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFeatureGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetFeatures(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmFormalElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFeatures(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFeatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetFunctions(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmFunction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetFunctions(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFunctions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetImmediateDescendants(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmClass
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetImmediateDescendants(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetImmediateDescendants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetInvariants(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v []IBmmAssertion
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetInvariants(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetInvariants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetIsAbstract(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsAbstract(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsAbstract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetIsOverride(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsOverride(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsOverride() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetIsPrimitive(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetIsPrimitive(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsPrimitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetName(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetName(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetPackage(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v IBmmPackage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetPackage(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetProcedures(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProcedure
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetProcedures(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProcedures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetProperties(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmProperty
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetProperties(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetScope(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v IBmmModelElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetScope(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetSourceSchemaId(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetSourceSchemaId(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSourceSchemaId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmSimpleClassBuilder_SetStaticProperties(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v map[string]IBmmStatic
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmSimpleClassBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetStaticProperties(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStaticProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmValueSetSpecBuilder_Build(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	tests := []struct {
		name   string
		fields fields
		want   *BmmValueSetSpec
		want1  []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmValueSetSpecBuilder{
				Builder: tt.fields.Builder,
			}
			got, got1 := i.Build()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Build() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBmmValueSetSpecBuilder_SetResourceId(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmValueSetSpecBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmValueSetSpecBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetResourceId(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetResourceId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBmmValueSetSpecBuilder_SetValueSetId(t *testing.T) {
	type fields struct {
		Builder Builder
	}
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BmmValueSetSpecBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &BmmValueSetSpecBuilder{
				Builder: tt.fields.Builder,
			}
			if got := i.SetValueSetId(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetValueSetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBmmEnumerationBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *BmmEnumerationBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBmmEnumerationBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBmmEnumerationBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBmmEnumerationIntegerBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *BmmEnumerationIntegerBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBmmEnumerationIntegerBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBmmEnumerationIntegerBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBmmEnumerationStringBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *BmmEnumerationStringBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBmmEnumerationStringBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBmmEnumerationStringBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBmmGenericClassBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *BmmGenericClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBmmGenericClassBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBmmGenericClassBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBmmSimpleClassBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *BmmSimpleClassBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBmmSimpleClassBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBmmSimpleClassBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBmmValueSetSpecBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *BmmValueSetSpecBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBmmValueSetSpecBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBmmValueSetSpecBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
