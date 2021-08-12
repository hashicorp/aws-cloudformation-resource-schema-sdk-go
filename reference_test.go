package cfschema_test

import (
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestReferenceField(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Reference       cfschema.Reference
		ExpectError     bool
		Expected        string
	}{
		{
			TestDescription: "empty",
			Reference:       cfschema.Reference(""),
			ExpectError:     true,
		},
		{
			TestDescription: "root",
			Reference:       cfschema.Reference("/"),
			ExpectError:     true,
		},
		{
			TestDescription: "definitions prefix only",
			Reference:       cfschema.Reference("/definitions"),
			ExpectError:     true,
		},
		{
			TestDescription: "properties prefix only",
			Reference:       cfschema.Reference("/properties"),
			ExpectError:     true,
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
			actual, err := testCase.Reference.Field()

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}

			if expected := testCase.Expected; actual != expected {
				t.Errorf("expected (%s), got: %s", expected, actual)
			}
		})
	}
}

func TestReferenceType(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Reference       cfschema.Reference
		ExpectError     bool
		Expected        string
	}{
		{
			TestDescription: "empty",
			Reference:       cfschema.Reference(""),
			ExpectError:     true,
		},
		{
			TestDescription: "root",
			Reference:       cfschema.Reference("/"),
			ExpectError:     true,
		},
		{
			TestDescription: "definitions prefix only",
			Reference:       cfschema.Reference("/definitions"),
			ExpectError:     true,
		},
		{
			TestDescription: "properties prefix only",
			Reference:       cfschema.Reference("/properties"),
			ExpectError:     true,
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
			actual, err := testCase.Reference.Type()

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}

			if expected := testCase.Expected; actual != expected {
				t.Errorf("expected (%s), got: %s", expected, actual)
			}
		})
	}
}
