package cfschema

type Resource struct {
	AdditionalIdentifiers           []PropertyJsonPointers `json:"additionalIdentifiers,omitempty"`
	AdditionalProperties            *bool                  `json:"additionalProperties,omitempty"`
	ConditionalCreateOnlyProperties PropertyJsonPointers   `json:"conditionalCreateOnlyProperties,omitempty"`
	CreateOnlyProperties            PropertyJsonPointers   `json:"createOnlyProperties,omitempty"`
	Definitions                     map[string]*Property   `json:"definitions,omitempty"`
	DeprecatedProperties            PropertyJsonPointers   `json:"deprecatedProperties,omitempty"`
	Description                     *string                `json:"description,omitempty"`
	Handlers                        map[string]*Handler    `json:"handler,omitempty"`
	PrimaryIdentifier               PropertyJsonPointers   `json:"primaryIdentifier,omitempty"`
	Properties                      map[string]*Property   `json:"properties,omitempty"`
	ReadOnlyProperties              PropertyJsonPointers   `json:"readOnlyProperties,omitempty"`
	ReplacementStrategy             *string                `json:"replacementStrategy,omitempty"`
	Required                        []string               `json:"required,omitempty"`
	ResourceLink                    *ResourceLink          `json:"resourceLink,omitempty"`
	SourceURL                       *string                `json:"sourceUrl,omitempty"`
	Taggable                        *bool                  `json:"taggable,omitempty"`
	TypeName                        *string                `json:"typeName,omitempty"`
	WriteOnlyProperties             PropertyJsonPointers   `json:"writeOnlyProperties,omitempty"`
}

func (r *Resource) IsCreateOnlyPropertyPath(path string) bool {
	if r == nil {
		return false
	}

	for _, createOnlyProperty := range r.CreateOnlyProperties {
		if createOnlyProperty.EqualsStringPath(path) {
			return true
		}
	}

	return false
}

func (r *Resource) IsRequired(name string) bool {
	if r == nil {
		return false
	}

	for _, req := range r.Required {
		if req == name {
			return true
		}
	}

	return false
}

func (r *Resource) ResolveProperty(p *Property) *Property {
	if r == nil || p == nil {
		return nil
	}

	if p.Ref != nil {
		switch p.Ref.Type() {
		case ReferenceTypeDefinitions:
			return r.Definitions[p.Ref.Field()]
		case ReferenceTypeProperties:
			return r.Properties[p.Ref.Field()]
		}
	}

	return p
}
