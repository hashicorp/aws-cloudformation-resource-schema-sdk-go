package cfschema

// Expand replaces all Definition and Property JSON Pointer references with their content.
//
// This functionality removes the need for recursive logic when accessing
// Definition and Property.
func (r *Resource) Expand() error {
	if r == nil {
		return nil
	}

	for _, definition := range r.Definitions {
		if ref := definition.Ref; ref != nil {
			*definition = *r.ResolveProperty(definition)
			definition.Ref = ref

			continue
		}

		switch definition.Type.String() {
		case PropertyTypeArray:
			if definition.Items == nil || definition.Items.Ref == nil {
				continue
			}

			ref := definition.Items.Ref
			*definition.Items = *r.ResolveProperty(definition.Items)
			definition.Items.Ref = ref
		case PropertyTypeObject:
			for _, objProperty := range definition.Properties {
				if ref := objProperty.Ref; ref != nil {
					*objProperty = *r.ResolveProperty(objProperty)
					objProperty.Ref = ref

					continue
				}

				switch objProperty.Type.String() {
				case PropertyTypeArray:
					if objProperty.Items == nil || objProperty.Items.Ref == nil {
						continue
					}

					ref := objProperty.Items.Ref
					*objProperty.Items = *r.ResolveProperty(objProperty.Items)
					objProperty.Items.Ref = ref
				}
			}
		}
	}

	for _, property := range r.Properties {
		if ref := property.Ref; ref != nil {
			*property = *r.ResolveProperty(property)
			property.Ref = ref

			continue
		}

		switch property.Type.String() {
		case PropertyTypeArray:
			if property.Items == nil || property.Items.Ref == nil {
				continue
			}

			ref := property.Items.Ref
			*property.Items = *r.ResolveProperty(property.Items)
			property.Items.Ref = ref
		case PropertyTypeObject:
			for _, objProperty := range property.Properties {
				if ref := objProperty.Ref; ref != nil {
					*objProperty = *r.ResolveProperty(objProperty)
					objProperty.Ref = ref

					continue
				}

				switch objProperty.Type.String() {
				case PropertyTypeArray:
					if objProperty.Items == nil || objProperty.Items.Ref == nil {
						continue
					}

					ref := objProperty.Items.Ref
					*objProperty.Items = *r.ResolveProperty(objProperty.Items)
					objProperty.Items.Ref = ref
				}
			}
		}
	}

	return nil
}
