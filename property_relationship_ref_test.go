// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cfschema_test

import (
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestPropertyRelationshipRef(t *testing.T) {
	testCases := []struct {
		TestDescription      string
		MetaSchemaPath       string
		ResourceSchemaPath   string
		PropertyPath         []string
		ExpectedPropertyPath string
		ExpectedTypeName     string
	}{
		{
			TestDescription:      "relationshipRef",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_S3_MultiRegionAccessPoint.json",
			PropertyPath:         []string{"Regions", "Bucket"},
			ExpectedPropertyPath: "/properties/BucketName",
			ExpectedTypeName:     "AWS::S3::Bucket",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			resource := loadAndValidateResourceSchema(t, testCase.MetaSchemaPath, testCase.ResourceSchemaPath)

			err := resource.Expand()

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			properties := resource.Properties
			for i, propertyName := range testCase.PropertyPath {
				property, ok := properties[propertyName]

				if !ok {
					t.Fatalf("expected resource property (%s), got none", propertyName)
				}

				if i == len(testCase.PropertyPath)-1 {
					if actual, expected := *property.RelationshipRef.TypeName, testCase.ExpectedTypeName; actual != expected {
						t.Errorf("expected resource property (%s) RelationshipRef.TypeName, (%s), got: %s", propertyName, expected, actual)
					}
					if actual, expected := *property.RelationshipRef.PropertyPath, testCase.ExpectedPropertyPath; actual.String() != expected {
						t.Errorf("expected resource property (%s) RelationshipRef.PropertyPath, (%s), got: %s", propertyName, expected, actual)
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
