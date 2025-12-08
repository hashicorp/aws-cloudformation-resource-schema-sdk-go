// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cfschema_test

import (
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestPropertySubschema_Resource(t *testing.T) {
	testCases := []struct {
		TestDescription    string
		MetaSchemaPath     string
		ResourceSchemaPath string
		ExpectError        bool
		ExpectedAllOf      int
		ExpectedAnyOf      int
		ExpectedOneOf      int
	}{
		{
			TestDescription:    "resource anyOf",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "AWS_CloudWatch_MetricStream.json",
			ExpectedAnyOf:      2,
		},
		{
			TestDescription:    "resource allOf",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "AWS_GameLift_Fleet.json",
			ExpectedAllOf:      3,
		},
		{
			TestDescription:    "resource no subschema",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "AWS_ECS_Cluster.json",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			resource := loadAndValidateResourceSchema(t, testCase.MetaSchemaPath, testCase.ResourceSchemaPath)

			if actual, expected := len(resource.AllOf), testCase.ExpectedAllOf; actual != expected {
				t.Errorf("expected %d allOf elements, got: %d", expected, actual)
			}
			if actual, expected := len(resource.AnyOf), testCase.ExpectedAnyOf; actual != expected {
				t.Errorf("expected %d anyOf elements, got: %d", expected, actual)
			}
			if actual, expected := len(resource.OneOf), testCase.ExpectedOneOf; actual != expected {
				t.Errorf("expected %d oneOf elements, got: %d", expected, actual)
			}
		})
	}
}

func TestPropertySubschema_Property(t *testing.T) {
	testCases := []struct {
		TestDescription    string
		MetaSchemaPath     string
		ResourceSchemaPath string
		ExpectError        bool
		PropertyPath       []string
		ExpectedAllOf      int
		ExpectedAnyOf      int
		ExpectedOneOf      int
	}{
		{
			TestDescription:    "property anyOf",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "AWS_S3Outposts_Bucket.json",
			PropertyPath:       []string{"LifecycleConfiguration", "Rules"},
			ExpectedAnyOf:      3,
		},
		{
			TestDescription:    "property oneOf",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "AWS_S3Outposts_Bucket.json",
			PropertyPath:       []string{"LifecycleConfiguration", "Rules", "Filter"},
			ExpectedOneOf:      3,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			resource := loadAndValidateResourceSchema(t, testCase.MetaSchemaPath, testCase.ResourceSchemaPath)

			err := resource.Expand()

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}

			properties := resource.Properties
			for i, propertyName := range testCase.PropertyPath {
				property, ok := properties[propertyName]

				if !ok {
					t.Fatalf("expected resource property (%s), got none", propertyName)
				}

				if property.Type == nil {
					t.Fatalf("unexpected resource property (%s) type, got none", propertyName)
				}

				if i == len(testCase.PropertyPath)-1 {
					if typ := (*property.Type).String(); typ == cfschema.PropertyTypeArray {
						property = property.Items
					}

					if actual, expected := len(property.AllOf), testCase.ExpectedAllOf; actual != expected {
						t.Errorf("expected %d allOf elements, got: %d", expected, actual)
					}
					if actual, expected := len(property.AnyOf), testCase.ExpectedAnyOf; actual != expected {
						t.Errorf("expected %d anyOf elements, got: %d", expected, actual)
					}
					if actual, expected := len(property.OneOf), testCase.ExpectedOneOf; actual != expected {
						t.Errorf("expected %d oneOf elements, got: %d", expected, actual)
					}
				} else {
					switch typ := (*property.Type).String(); typ {
					case cfschema.PropertyTypeArray:
						properties = property.Items.Properties
					case cfschema.PropertyTypeObject:
						properties = property.Properties
					default:
						t.Fatalf("resource property (%s) type (%s)", propertyName, typ)
					}
				}
			}
		})
	}
}
