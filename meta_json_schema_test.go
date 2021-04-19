package cfschema_test

import (
	"os"
	"path/filepath"
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestMetaJsonSchemaValidateResourceDocument(t *testing.T) {
	testCases := []struct {
		TestDescription    string
		MetaSchemaPath     string
		ResourceSchemaPath string
		ExpectError        bool
	}{
		{
			TestDescription:    "valid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "initech.tps.report.v1.json",
		},
		{
			TestDescription:    "invalid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "empty.json",
			ExpectError:        true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			metaSchema, err := cfschema.NewMetaJsonSchemaPath(filepath.Join("testdata", testCase.MetaSchemaPath))

			if err != nil {
				t.Fatalf("unexpected NewMetaJsonSchemaPath() error: %s", err)
			}

			resourceSchemaPath := filepath.Join("testdata", testCase.ResourceSchemaPath)
			file, err := os.ReadFile(resourceSchemaPath)

			if err != nil {
				t.Fatalf("unexpected error reading file (%s): %s", resourceSchemaPath, err)
			}

			err = metaSchema.ValidateResourceDocument(string(file))

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}
		})
	}
}

func TestMetaJsonSchemaValidateResourceJsonSchema(t *testing.T) {
	testCases := []struct {
		TestDescription    string
		MetaSchemaPath     string
		ResourceSchemaPath string
		ExpectError        bool
	}{
		{
			TestDescription:    "valid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "initech.tps.report.v1.json",
		},
		{
			TestDescription:    "invalid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "empty-object.json",
			ExpectError:        true,
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

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}
		})
	}
}

func TestMetaJsonSchemaValidateResourcePath(t *testing.T) {
	testCases := []struct {
		TestDescription    string
		MetaSchemaPath     string
		ResourceSchemaPath string
		ExpectError        bool
	}{
		{
			TestDescription:    "valid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "initech.tps.report.v1.json",
		},
		{
			TestDescription:    "invalid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "empty.json",
			ExpectError:        true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			metaSchema, err := cfschema.NewMetaJsonSchemaPath(filepath.Join("testdata", testCase.MetaSchemaPath))

			if err != nil {
				t.Fatalf("unexpected NewMetaJsonSchemaPath() error: %s", err)
			}

			err = metaSchema.ValidateResourcePath(filepath.Join("testdata", testCase.ResourceSchemaPath))

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}
		})
	}
}

func TestNewMetaJsonSchemaDocument(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Path            string
		ExpectError     bool
	}{
		{
			TestDescription: "valid",
			Path:            "provider.definition.schema.v1.json",
		},
		{
			TestDescription: "invalid",
			Path:            "empty.json",
			ExpectError:     true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			path := filepath.Join("testdata", testCase.Path)
			file, err := os.ReadFile(path)

			if err != nil {
				t.Fatalf("unexpected error reading file (%s): %s", path, err)
			}

			metaSchema, err := cfschema.NewMetaJsonSchemaDocument(string(file))

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if metaSchema == nil && !testCase.ExpectError {
				t.Error("expected result, got none")
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}
		})
	}
}

func TestNewMetaJsonSchemaPath(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Path            string
		ExpectError     bool
	}{
		{
			TestDescription: "valid",
			Path:            "provider.definition.schema.v1.json",
		},
		{
			TestDescription: "invalid",
			Path:            "empty.json",
			ExpectError:     true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			metaSchema, err := cfschema.NewMetaJsonSchemaPath(filepath.Join("testdata", testCase.Path))

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if metaSchema == nil && !testCase.ExpectError {
				t.Error("expected result, got none")
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}
		})
	}
}
