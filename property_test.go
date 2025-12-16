// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cfschema_test

import (
	"path/filepath"
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestProperty_IsRequired(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Property        *cfschema.Property
		Name            string
		Expected        bool
	}{
		{
			TestDescription: "nil resource",
			Property:        nil,
			Name:            "test",
			Expected:        false,
		},
		{
			TestDescription: "not found",
			Property:        &cfschema.Property{},
			Name:            "test",
			Expected:        false,
		},
		{
			TestDescription: "found",
			Property: &cfschema.Property{
				Required: []string{"test"},
			},
			Name:     "test",
			Expected: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			if actual, expected := testCase.Property.IsRequired(testCase.Name), testCase.Expected; actual != expected {
				t.Fatalf("expected (%t), got: %t", expected, actual)
			}
		})
	}
}

func loadAndValidateResourceSchema(t *testing.T, metaSchemaPath, resourceSchemaPath string) *cfschema.Resource {
	metaSchema, err := cfschema.NewMetaJsonSchemaPath(filepath.Join("testdata", metaSchemaPath))

	if err != nil {
		t.Fatalf("unexpected NewMetaJsonSchemaPath() error: %s", err)
	}

	resourceSchema, err := cfschema.NewResourceJsonSchemaPath(filepath.Join("testdata", resourceSchemaPath))

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

	return resource
}
