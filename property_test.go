package cfschema_test

import (
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestPropertyIsRequired(t *testing.T) {
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
