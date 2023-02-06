// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfschema

const (
	HandlerTypeCreate = "create"
	HandlerTypeDelete = "delete"
	HandlerTypeList   = "list"
	HandlerTypeRead   = "read"
	HandlerTypeUpdate = "update"
)

type Handler struct {
	Permissions      []string `json:"permissions,omitempty"`
	TimeoutInMinutes int      `json:"timeoutInMinutes,omitempty"`
}
