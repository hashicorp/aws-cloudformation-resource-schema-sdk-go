package cfschema_test

import (
	"path/filepath"
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestResourceExpand(t *testing.T) {
	testCases := []struct {
		TestDescription     string
		MetaSchemaPath      string
		ResourceSchemaPath  string
		ExpectError         bool
		ExpectPropertyTypes map[string]cfschema.Type
	}{
		{
			TestDescription:    "valid",
			MetaSchemaPath:     "provider.definition.schema.v1.json",
			ResourceSchemaPath: "initech.tps.report.v1.json",
			ExpectPropertyTypes: map[string]cfschema.Type{
				"ApprovalDate":     cfschema.PropertyTypeString,
				"DueDate":          cfschema.PropertyTypeString,
				"Memo":             cfschema.PropertyTypeObject,
				"SecondCopyOfMemo": cfschema.PropertyTypeObject,
			},
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

			err = resource.Expand()

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}

			for propertyName, expectedPropertyType := range testCase.ExpectPropertyTypes {
				property, ok := resource.Properties[propertyName]

				if !ok {
					t.Errorf("expected resource property (%s), got none", propertyName)
					continue
				}

				if property.Type == nil {
					t.Errorf("unexpected resource property (%s) type, got none", propertyName)
					continue
				}

				if actual, expected := *property.Type, expectedPropertyType; actual != expected {
					t.Errorf("expected resource property (%s) type (%s), got: %s", propertyName, expected, actual)
				}
			}
		})
	}
}

func TestResourceExpand_NestedDefinition(t *testing.T) {
	testCases := []struct {
		TestDescription      string
		MetaSchemaPath       string
		ResourceSchemaPath   string
		ExpectError          bool
		PropertyPath         []string
		ExpectedPropertyType cfschema.Type
	}{
		{
			TestDescription:      "valid",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_ECS_Cluster.json",
			PropertyPath:         []string{"Configuration", "ExecuteCommandConfiguration", "LogConfiguration", "CloudWatchEncryptionEnabled"},
			ExpectedPropertyType: cfschema.PropertyTypeBoolean,
		},
		{
			TestDescription:      "valid in array",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_WAFv2_RuleGroup.json",
			PropertyPath:         []string{"Rules", "Action", "Allow", "CustomRequestHandling", "InsertHeaders", "Name"},
			ExpectedPropertyType: cfschema.PropertyTypeString,
		},
		{
			TestDescription:      "no type specified",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_IoT_JobTemplate.json",
			PropertyPath:         []string{"PresignedUrlConfig", "ExpiresInSec"},
			ExpectedPropertyType: cfschema.PropertyTypeInteger,
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

			err = resource.Expand()

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}

			properties := resource.Properties
			for i, propertyName := range testCase.PropertyPath {
				property, ok := properties[propertyName]

				if !ok {
					t.Fatalf("expected resource property (%s), got none", propertyName)
				}

				if property.Type == nil {
					if len(property.Properties) > 0 {
						properties = property.Properties
						continue
					} else {
						t.Fatalf("unexpected resource property (%s) type, got none", propertyName)
					}
				}

				if i == len(testCase.PropertyPath)-1 {
					if actual, expected := *property.Type, testCase.ExpectedPropertyType; actual != expected {
						t.Errorf("expected resource property (%s) type (%s), got: %s", propertyName, expected, actual)
					}
				} else {
					switch typ := (*property.Type).String(); typ {
					case cfschema.PropertyTypeArray:
						properties = property.Items.Properties
					case cfschema.PropertyTypeObject:
						properties = property.Properties
					default:
						t.Fatalf("resource property (%s) type (%s)", propertyName, typ)
					}
				}
			}
		})
	}
}

func TestResourceExpand_PatternProperties(t *testing.T) {
	testCases := []struct {
		TestDescription      string
		MetaSchemaPath       string
		ResourceSchemaPath   string
		ExpectError          bool
		PropertyPath         []string
		ExpectedPropertyType cfschema.Type
	}{
		{
			TestDescription:      "valid",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_GreengrassV2_ComponentVersion.json",
			PropertyPath:         []string{"LambdaFunction", "ComponentDependencies", "VersionRequirement"},
			ExpectedPropertyType: cfschema.PropertyTypeString,
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

			err = resource.Expand()

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}

			properties := resource.Properties
			var patternProperties map[string]*cfschema.Property
			for i, propertyName := range testCase.PropertyPath {
				property, ok := properties[propertyName]

				if !ok {
					if len(patternProperties) == 1 {
						for _, v := range patternProperties {
							property = v
						}

						if actual, expected := *property.Type, cfschema.Type(cfschema.PropertyTypeObject); actual != expected {
							t.Fatalf("expected resource property (%s) type (%s), got: %s", propertyName, expected, actual)
						}

						property, ok = property.Properties[propertyName]

						if !ok {
							t.Fatalf("expected resource property (%s), got none", propertyName)
						}
					} else {
						t.Fatalf("expected resource property (%s), got none", propertyName)
					}
				}

				if property.Type == nil {
					t.Fatalf("unexpected resource property (%s) type, got none", propertyName)
				}

				if i == len(testCase.PropertyPath)-1 {
					if actual, expected := *property.Type, testCase.ExpectedPropertyType; actual != expected {
						t.Errorf("expected resource property (%s) type (%s), got: %s", propertyName, expected, actual)
					}
				} else {
					if actual, expected := *property.Type, cfschema.Type(cfschema.PropertyTypeObject); actual != expected {
						t.Fatalf("expected resource property (%s) type (%s), got: %s", propertyName, expected, actual)
					}

					patternProperties = property.PatternProperties
					properties = property.Properties
				}
			}
		})
	}
}

func TestResourceExpand_SecondLevelNestedDefinition(t *testing.T) {
	testCases := []struct {
		TestDescription      string
		MetaSchemaPath       string
		ResourceSchemaPath   string
		ExpectError          bool
		PropertyPath         []string
		ExpectedPropertyType cfschema.Type
	}{
		{
			TestDescription:      "valid 1 level simple",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_ImageBuilder_DistributionConfiguration.json",
			PropertyPath:         []string{"Distributions", "Region"},
			ExpectedPropertyType: cfschema.PropertyTypeString,
		},
		{
			TestDescription:      "valid 1 level object",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_ImageBuilder_DistributionConfiguration.json",
			PropertyPath:         []string{"Distributions", "ContainerDistributionConfiguration", "TargetRepository"},
			ExpectedPropertyType: cfschema.PropertyTypeObject,
		},
		{
			TestDescription:      "valid 2 level simple",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_ImageBuilder_DistributionConfiguration.json",
			PropertyPath:         []string{"Distributions", "ContainerDistributionConfiguration", "TargetRepository", "RepositoryName"},
			ExpectedPropertyType: cfschema.PropertyTypeString,
		},
		{
			TestDescription:      "unwrap oneOf type object, single entry with ref",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_S3ObjectLambda_AccessPoint.json",
			PropertyPath:         []string{"ObjectLambdaConfiguration", "TransformationConfigurations", "ContentTransformation", "AwsLambda", "FunctionArn"},
			ExpectedPropertyType: cfschema.PropertyTypeString,
		},
		{
			TestDescription:      "unwrap oneOf type empty, single entry with refs",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_S3Outposts_Bucket.json",
			PropertyPath:         []string{"LifecycleConfiguration", "Rules", "Filter", "AndOperator", "Prefix"},
			ExpectedPropertyType: cfschema.PropertyTypeString,
		},
		{
			TestDescription:      "unwrap oneOf type object, multiple entries with refs",
			MetaSchemaPath:       "provider.definition.schema.v1.json",
			ResourceSchemaPath:   "AWS_S3_StorageLens.json",
			PropertyPath:         []string{"StorageLensConfiguration", "DataExport", "S3BucketDestination", "Encryption", "SSEKMS", "KeyId"},
			ExpectedPropertyType: cfschema.PropertyTypeString,
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

			err = resource.Expand()

			if err != nil && !testCase.ExpectError {
				t.Fatalf("unexpected error: %s", err)
			}

			if err == nil && testCase.ExpectError {
				t.Fatal("expected error, got none")
			}

			properties := resource.Properties
			for i, propertyName := range testCase.PropertyPath {
				property, ok := properties[propertyName]

				if !ok {
					t.Fatalf("expected resource property (%s), got none", propertyName)
				}

				if property.Type == nil {
					t.Fatalf("unexpected resource property (%s) type, got none", propertyName)
				}

				if i == len(testCase.PropertyPath)-1 {
					if actual, expected := *property.Type, testCase.ExpectedPropertyType; actual != expected {
						t.Errorf("expected resource property (%s) type (%s), got: %s", propertyName, expected, actual)
					}
				} else {
					switch typ := (*property.Type).String(); typ {
					case cfschema.PropertyTypeArray:
						properties = property.Items.Properties
					case cfschema.PropertyTypeObject:
						properties = property.Properties
					default:
						t.Fatalf("resource property (%s) type (%s)", propertyName, typ)
					}
				}
			}
		})
	}
}
