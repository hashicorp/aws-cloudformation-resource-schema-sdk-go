// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cfschema_test

import (
	"testing"
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
			resource := loadAndValidateResourceSchema(t, testCase.MetaSchemaPath, testCase.ResourceSchemaPath)

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
