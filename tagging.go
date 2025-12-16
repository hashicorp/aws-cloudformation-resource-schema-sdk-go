// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cfschema

type Tagging struct {
	Taggable                 *bool                `json:"taggable,omitempty"`
	TagOnCreate              *bool                `json:"tagOnCreate,omitempty"`
	TagUpdatable             *bool                `json:"tagUpdatable,omitempty"`
	CloudFormationSystemTags *bool                `json:"cloudFormationSystemTags,omitempty"`
	TagProperty              *PropertyJsonPointer `json:"tagProperty,omitempty"`
	Permissions              []string             `json:"permissions,omitempty"`
}
