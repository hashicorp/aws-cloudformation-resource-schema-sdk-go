// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfschema

type HandlerSchema struct {
	AllOf      []*PropertySubschema `json:"allOf,omitempty"`
	AnyOf      []*PropertySubschema `json:"anyOf,omitempty"`
	OneOf      []*PropertySubschema `json:"oneOf,omitempty"`
	Properties map[string]*Property `json:"properties,omitempty"`
	Required   []string             `json:"required,omitempty"`
}
