package cfschema_test

import (
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestReferenceField(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Reference       cfschema.Reference
		Expected        string
	}{
		{
			TestDescription: "empty",
			Reference:       cfschema.Reference(""),
			Expected:        "",
		},
		{
			TestDescription: "root",
			Reference:       cfschema.Reference("/"),
			Expected:        "",
		},
		{
			TestDescription: "definitions prefix only",
			Reference:       cfschema.Reference("/definitions"),
			Expected:        "",
		},
		{
			TestDescription: "properties prefix only",
			Reference:       cfschema.Reference("/properties"),
			Expected:        "",
		},
		{
			TestDescription: "definition",
			Reference:       cfschema.Reference("/definitions/test"),
			Expected:        "test",
		},
		{
			TestDescription: "property",
			Reference:       cfschema.Reference("/properties/test"),
			Expected:        "test",
		},
		{
			TestDescription: "definition with prefix",
			Reference:       cfschema.Reference("#/definitions/test"),
			Expected:        "test",
		},
		{
			TestDescription: "property with prefix",
			Reference:       cfschema.Reference("#/properties/test"),
			Expected:        "test",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			if actual, expected := testCase.Reference.Field(), testCase.Expected; actual != expected {
				t.Errorf("expected (%s), got: %s", expected, actual)
			}
		})
	}
}

func TestReferenceType(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Reference       cfschema.Reference
		Expected        string
	}{
		{
			TestDescription: "empty",
			Reference:       cfschema.Reference(""),
			Expected:        "",
		},
		{
			TestDescription: "root",
			Reference:       cfschema.Reference("/"),
			Expected:        "",
		},
		{
			TestDescription: "definitions prefix only",
			Reference:       cfschema.Reference("/definitions"),
			Expected:        "",
		},
		{
			TestDescription: "properties prefix only",
			Reference:       cfschema.Reference("/properties"),
			Expected:        "",
		},
		{
			TestDescription: "definition",
			Reference:       cfschema.Reference("/definitions/test"),
			Expected:        cfschema.ReferenceTypeDefinitions,
		},
		{
			TestDescription: "property",
			Reference:       cfschema.Reference("/properties/test"),
			Expected:        cfschema.ReferenceTypeProperties,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			if actual, expected := testCase.Reference.Type(), testCase.Expected; actual != expected {
				t.Errorf("expected (%s), got: %s", expected, actual)
			}
		})
	}
}
