// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfschema_test

import (
	"path/filepath"
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestHandlerSchema(t *testing.T) {
	testCases := []struct {
		TestDescription       string
		MetaSchemaPath        string
		ResourceSchemaPath    string
		ExpectError           bool
		ExpectedHandlerSchema int
	}{
		{
			TestDescription:    "no handlerSchema",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "AWS_CloudWatch_MetricStream.json",
		},
		{
			TestDescription:       "list handlerSchema",
			MetaSchemaPath:        "provider.definition.schema.v1.json",
			ResourceSchemaPath:    "AWS_NetworkManager_TransitGatewayRegistration.json",
			ExpectedHandlerSchema: 1,
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

			got := 0

			for _, handler := range resource.Handlers {
				if handler.HandlerSchema != nil {
					got++
				}
			}

			if actual, expected := got, testCase.ExpectedHandlerSchema; actual != expected {
				t.Errorf("expected %d handlerSchema elements, got: %d", expected, actual)
			}
		})
	}
}
