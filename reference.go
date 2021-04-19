package cfschema

import "strings"

const (
	ReferenceAnchor          = "#"
	ReferenceSeparator       = "/"
	ReferenceTypeDefinitions = "definitions"
	ReferenceTypeProperties  = "properties"
)

// Reference is an internal implementation for RFC 6901 JSON Pointer values.
type Reference string

// Field returns the JSON Pointer string path after the type.
func (r *Reference) Field() string {
	if r == nil {
		return ""
	}

	referenceParts := strings.Split(strings.TrimPrefix(string(*r), ReferenceAnchor), ReferenceSeparator)

	if len(referenceParts) != 3 {
		return ""
	}

	return referenceParts[2]
}

// String returns the string representation of Reference.
func (r *Reference) String() string {
	if r == nil {
		return ""
	}

	return string(*r)
}

// Type returns the first path part of the JSON Pointer.
//
// In CloudFormation Resources, this should be definitions or properties.
func (r *Reference) Type() string {
	if r == nil {
		return ""
	}

	referenceParts := strings.Split(strings.TrimPrefix(string(*r), ReferenceAnchor), ReferenceSeparator)

	if len(referenceParts) != 3 {
		return ""
	}

	return referenceParts[1]
}
