package cfschema_test

import (
	"path/filepath"
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestResourceExpand(t *testing.T) {
	testCases := []struct {
		TestDescription     string
		MetaSchemaPath      string
		ResourceSchemaPath  string
		ExpectError         bool
		ExpectPropertyTypes map[string]cfschema.Type
	}{
		{
			TestDescription:    "valid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "initech.tps.report.v1.json",
			ExpectPropertyTypes: map[string]cfschema.Type{
				"ApprovalDate":     cfschema.PropertyTypeString,
				"DueDate":          cfschema.PropertyTypeString,
				"Memo":             cfschema.PropertyTypeObject,
				"SecondCopyOfMemo": cfschema.PropertyTypeObject,
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			metaSchema, err := cfschema.NewMetaJsonSchemaPath(filepath.Join("testdata", testCase.MetaSchemaPath))

			if err != nil {
				t.Fatalf("unexpected NewMetaJsonSchemaPath() error: %s", err)
			}

			resourceSchema, err := cfschema.NewResourceJsonSchemaPath(filepath.Join("testdata", testCase.ResourceSchemaPath))

			if err != nil {
				t.Fatalf("unexpected NewResourceJsonSchemaPath() error: %s", err)
			}

			err = metaSchema.ValidateResourceJsonSchema(resourceSchema)

			if err != nil {
				t.Fatalf("unexpected ValidateResourceJsonSchema() error: %s", err)
			}

			resource, err := resourceSchema.Resource()

			if err != nil {
				t.Fatalf("unexpected Resource() error: %s", err)
			}

			err = resource.Expand()

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}

			for propertyName, expectedPropertyType := range testCase.ExpectPropertyTypes {
				property, ok := resource.Properties[propertyName]

				if !ok {
					t.Errorf("expected resource property (%s), got none", propertyName)
					continue
				}

				if property.Type == nil {
					t.Errorf("unexpected resource property (%s) type, got none", propertyName)
					continue
				}

				if actual, expected := *property.Type, expectedPropertyType; actual != expected {
					t.Errorf("expected resource property (%s) type (%s), got: %s", propertyName, expected, actual)
				}
			}
		})
	}
}
