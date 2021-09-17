package cfschema_test

import (
	"testing"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
)

func TestSanitize(t *testing.T) {
	testCases := []struct {
		TestDescription   string
		InputDocument     string
		SanitizedDocument string
	}{
		{
			TestDescription: "Nothing to sanitize",
			InputDocument: `
{
  "LogGroupName": {
    "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
    "type": "string",
    "minLength": 1,
    "maxLength": 512
  },
}
			`,
			SanitizedDocument: `
{
  "LogGroupName": {
    "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
    "type": "string",
    "minLength": 1,
    "maxLength": 512
  },
}
			`,
		},
		{
			TestDescription: "Sanitize pattern",
			InputDocument: `
{
  "LogGroupName": {
    "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
    "type": "string",
    "minLength": 1,
    "maxLength": 512,
    "pattern": "^[.\\-_/#A-Za-z0-9]{1,512}\\Z"
  },
  "KmsKeyId": {
    "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
    "type": "string",
    "pattern": "^arn:[a-z0-9-]+:kms:[a-z0-9-]+:\\d{12}:(key|alias)/.+\\Z",
    "maxLength": 256
  }
}
			`,
			SanitizedDocument: `
{
  "LogGroupName": {
    "description": "The name of the log group. If you don't specify a name, AWS CloudFormation generates a unique ID for the log group.",
    "type": "string",
    "minLength": 1,
    "maxLength": 512,
    "pattern": ""
  },
  "KmsKeyId": {
    "description": "The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.",
    "type": "string",
    "pattern": "",
    "maxLength": 256
  }
}
			`,
		},
		{
			TestDescription: "Sanitize patternProperties",
			InputDocument: `
{
  "BackupPlanTags": {
    "type": "object",
    "additionalProperties": false,
    "patternProperties": {
      "^.{1,128}$": {
        "type": "string"
      }
    }
  },
  "RecoveryPointTags": {
    "type": "object",
    "patternProperties": {
      "^.{1,128}$": {
        "type": "string"
      },
    "additionalProperties": false
    }
  }
}
			`,
			SanitizedDocument: `
{
  "BackupPlanTags": {
    "type": "object",
    "additionalProperties": false,
    "patternProperties": {
      "": {
        "type": "string"
      }
    }
  },
  "RecoveryPointTags": {
    "type": "object",
    "patternProperties": {
      "": {
        "type": "string"
      },
    "additionalProperties": false
    }
  }
}
			`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.TestDescription, func(t *testing.T) {
			got := cfschema.Sanitize(testCase.InputDocument)

			if got != testCase.SanitizedDocument {
				t.Errorf("expected: %s, got: %s", testCase.SanitizedDocument, got)
			}
		})
	}
}
