package cfschema_test

import (
	"reflect"
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestPropertyJsonPointerEqualsPath(t *testing.T) {
	testCases := []struct {
		TestDescription     string
		PropertyJsonPointer cfschema.PropertyJsonPointer
		Path                []string
		Expected            bool
	}{
		{
			TestDescription:     "empty",
			PropertyJsonPointer: "",
			Path:                []string{"test"},
			Expected:            false,
		},
		{
			TestDescription:     "not found",
			PropertyJsonPointer: "",
			Path:                []string{"test"},
			Expected:            false,
		},
		{
			TestDescription:     "first level match",
			PropertyJsonPointer: "/properties/test",
			Path:                []string{"test"},
			Expected:            true,
		},
		{
			TestDescription:     "first level mismatch",
			PropertyJsonPointer: "/properties/nottest",
			Path:                []string{"test"},
			Expected:            false,
		},
		{
			TestDescription:     "first level mismatch with prefix",
			PropertyJsonPointer: "/properties/testing",
			Path:                []string{"test"},
			Expected:            false,
		},
		{
			TestDescription:     "multi level match",
			PropertyJsonPointer: "/properties/parent/nested",
			Path:                []string{"parent", "nested"},
			Expected:            true,
		},
		{
			TestDescription:     "multi level mismatch only front",
			PropertyJsonPointer: "/properties/parent/nested",
			Path:                []string{"parent"},
			Expected:            false,
		},
		{
			TestDescription:     "multi level mismatch only back",
			PropertyJsonPointer: "/properties/parent/nested",
			Path:                []string{"nested"},
			Expected:            false,
		},
		{
			TestDescription:     "multi level mismatch wrong front",
			PropertyJsonPointer: "/properties/parent/nested",
			Path:                []string{"notparent", "nested"},
			Expected:            false,
		},
		{
			TestDescription:     "multi level mismatch wrong back",
			PropertyJsonPointer: "/properties/parent/nested",
			Path:                []string{"parent", "notnested"},
			Expected:            false,
		},
		{
			TestDescription:     "multi level mismatch with prefix",
			PropertyJsonPointer: "/properties/parent/testing",
			Path:                []string{"parent", "test"},
			Expected:            false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			if actual, expected := testCase.PropertyJsonPointer.EqualsPath(testCase.Path), testCase.Expected; actual != expected {
				t.Fatalf("expected (%t), got: %t", expected, actual)
			}
		})
	}
}

func TestPropertyJsonPointerEqualsStringPath(t *testing.T) {
	testCases := []struct {
		TestDescription     string
		PropertyJsonPointer cfschema.PropertyJsonPointer
		Path                string
		Expected            bool
	}{
		{
			TestDescription:     "empty",
			PropertyJsonPointer: "",
			Path:                "/test",
			Expected:            false,
		},
		{
			TestDescription:     "not found",
			PropertyJsonPointer: "",
			Path:                "/test",
			Expected:            false,
		},
		{
			TestDescription:     "first level match",
			PropertyJsonPointer: "/properties/test",
			Path:                "/test",
			Expected:            true,
		},
		{
			TestDescription:     "first level mismatch",
			PropertyJsonPointer: "/properties/nottest",
			Path:                "/test",
			Expected:            false,
		},
		{
			TestDescription:     "first level mismatch with prefix",
			PropertyJsonPointer: "/properties/testing",
			Path:                "/test",
			Expected:            false,
		},
		{
			TestDescription:     "multi level match",
			PropertyJsonPointer: "/properties/parent/nested",
			Path:                "/parent/nested",
			Expected:            true,
		},
		{
			TestDescription:     "multi level mismatch only front",
			PropertyJsonPointer: "/properties/parent/nested",
			Path:                "/parent",
			Expected:            false,
		},
		{
			TestDescription:     "multi level mismatch only back",
			PropertyJsonPointer: "/properties/parent/nested",
			Path:                "/nested",
			Expected:            false,
		},
		{
			TestDescription:     "multi level mismatch wrong front",
			PropertyJsonPointer: "/properties/parent/nested",
			Path:                "/notparent/nested",
			Expected:            false,
		},
		{
			TestDescription:     "multi level mismatch wrong back",
			PropertyJsonPointer: "/properties/parent/nested",
			Path:                "/parent/notnested",
			Expected:            false,
		},
		{
			TestDescription:     "multi level mismatch with prefix",
			PropertyJsonPointer: "/properties/parent/testing",
			Path:                "/parent/test",
			Expected:            false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			if actual, expected := testCase.PropertyJsonPointer.EqualsStringPath(testCase.Path), testCase.Expected; actual != expected {
				t.Fatalf("expected (%t), got: %t", expected, actual)
			}
		})
	}
}

func TestPropertyJsonPointerPath(t *testing.T) {
	testCases := []struct {
		TestDescription     string
		PropertyJsonPointer cfschema.PropertyJsonPointer
		Expected            []string
	}{
		{
			TestDescription:     "empty",
			PropertyJsonPointer: "",
			Expected:            []string{""},
		},
		{
			TestDescription:     "one level",
			PropertyJsonPointer: "/properties/test",
			Expected:            []string{"test"},
		},
		{
			TestDescription:     "multi level",
			PropertyJsonPointer: "/properties/parent/nested",
			Expected:            []string{"parent", "nested"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			if actual, expected := testCase.PropertyJsonPointer.Path(), testCase.Expected; !reflect.DeepEqual(actual, expected) {
				t.Fatalf("expected (%#v), got: %#v", expected, actual)
			}
		})
	}
}
