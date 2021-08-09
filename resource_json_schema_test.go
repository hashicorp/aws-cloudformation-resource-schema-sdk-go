package cfschema_test

import (
	"os"
	"path/filepath"
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestResourceJsonSchemaResource(t *testing.T) {
	testCases := []struct {
		TestDescription     string
		MetaSchemaPath      string
		ResourceSchemaPath  string
		ExpectError         bool
		ExpectedNumHandlers int
	}{
		{
			TestDescription:     "valid",
			MetaSchemaPath:      "provider.definition.schema.v1.json",
			ResourceSchemaPath:  "initech.tps.report.v1.json",
			ExpectedNumHandlers: 5,
		},
		{
			TestDescription:     "valid with negative-lookahead in pattern",
			MetaSchemaPath:      "provider.definition.schema.v1.json",
			ResourceSchemaPath:  "initech.tps.report.negative-lookahead.json",
			ExpectedNumHandlers: 5,
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

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if resource == nil && !testCase.ExpectError {
				t.Error("expected result, got none")
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}

			if got := len(resource.Handlers); got != testCase.ExpectedNumHandlers {
				t.Fatalf("expected %d handlers, got %d", testCase.ExpectedNumHandlers, got)
			}
		})
	}
}

func TestResourceJsonSchemaValidateConfigurationDocument(t *testing.T) {
	testCases := []struct {
		TestDescription    string
		MetaSchemaPath     string
		ResourceSchemaPath string
		ConfigurationPath  string
		ExpectError        bool
	}{
		{
			TestDescription:    "valid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "initech.tps.report.v1.json",
			ConfigurationPath:  "valid-initech-tps-report-v1-configuration.json",
		},
		{
			TestDescription:    "invalid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "initech.tps.report.v1.json",
			ConfigurationPath:  "empty.json",
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

			if err != nil {
				t.Fatalf("unexpected ValidateResourceJsonSchema() error: %s", err)
			}

			configurationPath := filepath.Join("testdata", testCase.ConfigurationPath)
			file, err := os.ReadFile(configurationPath)

			if err != nil {
				t.Fatalf("unexpected error reading file (%s): %s", configurationPath, err)
			}

			err = resourceSchema.ValidateConfigurationDocument(string(file))

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}
		})
	}
}

func TestResourceJsonSchemaValidateConfigurationPath(t *testing.T) {
	testCases := []struct {
		TestDescription    string
		MetaSchemaPath     string
		ResourceSchemaPath string
		ConfigurationPath  string
		ExpectError        bool
	}{
		{
			TestDescription:    "valid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "initech.tps.report.v1.json",
			ConfigurationPath:  "valid-initech-tps-report-v1-configuration.json",
		},
		{
			TestDescription:    "invalid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "initech.tps.report.v1.json",
			ConfigurationPath:  "empty.json",
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

			if err != nil {
				t.Fatalf("unexpected ValidateResourceJsonSchema() error: %s", err)
			}

			err = resourceSchema.ValidateConfigurationPath(filepath.Join("testdata", testCase.ConfigurationPath))

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}
		})
	}
}

func TestNewResourceJsonSchemaDocument(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Path            string
		ExpectError     bool
	}{
		{
			TestDescription: "valid",
			Path:            "initech.tps.report.v1.json",
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
			file, err := os.ReadFile(filepath.Join("testdata", testCase.Path))

			if err != nil {
				t.Fatalf("error reading file (%s): %s", testCase.Path, err)
			}

			metaSchema, err := cfschema.NewResourceJsonSchemaDocument(string(file))

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

func TestNewResourceJsonSchemaPath(t *testing.T) {
	testCases := []struct {
		TestDescription string
		Path            string
		ExpectError     bool
	}{
		{
			TestDescription: "valid",
			Path:            "initech.tps.report.v1.json",
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
			metaSchema, err := cfschema.NewResourceJsonSchemaPath(filepath.Join("testdata", testCase.Path))

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
