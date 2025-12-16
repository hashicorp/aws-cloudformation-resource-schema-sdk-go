// Copyright IBM Corp. 2021, 2025
// SPDX-License-Identifier: MPL-2.0

package cfschema

type PropertyRelationshipRef struct {
	PropertyPath *PropertyJsonPointer `json:"propertyPath,omitempty"`
	TypeName     *string              `json:"typeName,omitempty"`
}
