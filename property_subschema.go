// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cfschema

type PropertySubschema struct {
	AllOf      []*PropertySubschema `json:"allOf,omitempty"`
	AnyOf      []*PropertySubschema `json:"anyOf,omitempty"`
	Default    any                  `json:"default,omitempty"`
	OneOf      []*PropertySubschema `json:"oneOf,omitempty"`
	Properties map[string]*Property `json:"properties,omitempty"`
	Ref        string               `json:"$ref,omitempty"`
	Required   []string             `json:"required,omitempty"`
}
