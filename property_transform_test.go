package cfschema_test

import (
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestPropertyTransformValue(t *testing.T) {
	testCases := []struct {
		Name              string
		PropertyTransform cfschema.PropertyTransform
		Path              []string
		Expected          string
	}{
		{
			Name: "found",
			PropertyTransform: cfschema.PropertyTransform{
				"/properties/TestPath": "$lowercase(TestPath)",
			},
			Path:     []string{"TestPath"},
			Expected: "$lowercase(TestPath)",
		},
		{
			Name: "not found",
			PropertyTransform: cfschema.PropertyTransform{
				"/properties/TestPath": "$lowercase(TestPath)",
			},
			Path:     []string{"TestPath", "SubProperty"},
			Expected: "",
		},
		{
			Name: "found nested property",
			PropertyTransform: cfschema.PropertyTransform{
				"/properties/TestPath/SubProperty": "$lowercase(TestPath.SubProperty)",
			},
			Path:     []string{"TestPath", "SubProperty"},
			Expected: "$lowercase(TestPath.SubProperty)",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if actual, _ := testCase.PropertyTransform.Value(testCase.Path); actual != testCase.Expected {
				t.Fatalf("expected (%s), got: %s", testCase.Expected, actual)
			}
		})
	}
}
