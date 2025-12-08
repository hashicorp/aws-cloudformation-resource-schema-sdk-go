// Copyright IBM Corp. 2021, 2025
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
	HandlerSchema    *HandlerSchema `json:"handlerSchema,omitempty"`
	Permissions      []string       `json:"permissions,omitempty"`
	TimeoutInMinutes int            `json:"timeoutInMinutes,omitempty"`
}
