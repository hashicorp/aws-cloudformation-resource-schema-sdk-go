// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfschema

type ResourceLink struct {
	Comment     *string           `json:"$comment,omitempty"`
	Mappings    map[string]string `json:"mappings,omitempty"`
	TemplateURI *string           `json:"templateUri,omitempty"`
}
