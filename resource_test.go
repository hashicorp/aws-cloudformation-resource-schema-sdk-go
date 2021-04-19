package cfschema_test

import (
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestResourceIsRequired(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Resource        *cfschema.Resource
		Name            string
		Expected        bool
	}{
		{
			TestDescription: "nil resource",
			Resource:        nil,
			Name:            "test",
			Expected:        false,
		},
		{
			TestDescription: "not found",
			Resource:        &cfschema.Resource{},
			Name:            "test",
			Expected:        false,
		},
		{
			TestDescription: "found",
			Resource: &cfschema.Resource{
				Required: []string{"test"},
			},
			Name:     "test",
			Expected: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			if actual, expected := testCase.Resource.IsRequired(testCase.Name), testCase.Expected; actual != expected {
				t.Fatalf("expected (%t), got: %t", expected, actual)
			}
		})
	}
}

func TestResourceResolveProperty(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Resource        *cfschema.Resource
		Property        *cfschema.Property
		Expected        bool
		ExpectRef       bool
		ExpectType      bool
	}{
		{
			TestDescription: "nil resource",
			Resource:        nil,
			Expected:        false,
			ExpectRef:       false,
			ExpectType:      false,
		},
		{
			TestDescription: "nil property",
			Resource:        &cfschema.Resource{},
			Expected:        false,
			ExpectRef:       false,
			ExpectType:      false,
		},
		{
			TestDescription: "passthrough",
			Resource:        &cfschema.Resource{},
			Property:        &cfschema.Property{},
			Expected:        true,
			ExpectRef:       false,
			ExpectType:      false,
		},
		{
			TestDescription: "missing definition",
			Resource:        &cfschema.Resource{},
			Property: &cfschema.Property{
				Ref: testReference("/definitions/test"),
			},
			Expected:   false,
			ExpectRef:  false,
			ExpectType: false,
		},
		{
			TestDescription: "missing property",
			Resource:        &cfschema.Resource{},
			Property: &cfschema.Property{
				Ref: testReference("/properties/test"),
			},
			Expected:   false,
			ExpectRef:  false,
			ExpectType: false,
		},
		{
			TestDescription: "definition ref",
			Resource: &cfschema.Resource{
				Definitions: map[string]*cfschema.Property{
					"test": {
						Ref: testReference("/definitions/test2"),
					},
					"test2": {
						Type: testType(cfschema.PropertyTypeBoolean),
					},
				},
			},
			Property: &cfschema.Property{
				Ref: testReference("/definitions/test"),
			},
			Expected:   true,
			ExpectRef:  true,
			ExpectType: false,
		},
		{
			TestDescription: "definition type",
			Resource: &cfschema.Resource{
				Definitions: map[string]*cfschema.Property{
					"test": {
						Type: testType(cfschema.PropertyTypeBoolean),
					},
				},
			},
			Property: &cfschema.Property{
				Ref: testReference("/definitions/test"),
			},
			Expected:   true,
			ExpectRef:  false,
			ExpectType: true,
		},
		{
			TestDescription: "property ref",
			Resource: &cfschema.Resource{
				Properties: map[string]*cfschema.Property{
					"test": {
						Type: testType(cfschema.PropertyTypeBoolean),
					},
				},
			},
			Property: &cfschema.Property{
				Ref: testReference("/properties/test"),
			},
			Expected:   true,
			ExpectRef:  false,
			ExpectType: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			actualProperty := testCase.Resource.ResolveProperty(testCase.Property)

			if actualProperty != nil && !testCase.Expected {
				t.Fatalf("expected no property, got one")
			}

			if actualProperty == nil && testCase.Expected {
				t.Fatalf("expected property, got none")
			}

			if actualProperty != nil {
				if actualProperty.Ref != nil && !testCase.ExpectRef {
					t.Fatalf("expected no property ref, got one")
				}

				if actualProperty.Ref == nil && testCase.ExpectRef {
					t.Fatalf("expected property ref, got none")
				}

				if actualProperty.Type != nil && !testCase.ExpectType {
					t.Fatalf("expected no property type, got one")
				}

				if actualProperty.Type == nil && testCase.ExpectType {
					t.Fatalf("expected property type, got none")
				}
			}
		})
	}
}

func testReference(r string) *cfschema.Reference {
	result := cfschema.Reference(r)

	return &result
}

func testType(t string) *cfschema.Type {
	result := cfschema.Type(t)

	return &result
}
