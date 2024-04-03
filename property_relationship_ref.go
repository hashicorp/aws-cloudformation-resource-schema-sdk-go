// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfschema

type PropertyRelationshipRef struct {
	PropertyPath *PropertyJsonPointer `json:"propertyPath,omitempty"`
	TypeName     *string              `json:"typeName,omitempty"`
}
