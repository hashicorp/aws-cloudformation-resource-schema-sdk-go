package cfschema

import (
	"fmt"
)

// Expand replaces all Definition and Property JSON Pointer references with their content.
// This functionality removes the need for recursive logic when accessingn Definition and Property.
func (r *Resource) Expand() error {
	if r == nil {
		return nil
	}

	err := r.ResolveProperties(r.Definitions)

	if err != nil {
		return fmt.Errorf("error expanding Resource (%s) Definitions: %w", *r.TypeName, err)
	}

	err = r.ResolveProperties(r.Properties)

	if err != nil {
		return fmt.Errorf("error expanding Resource (%s) Properties: %w", *r.TypeName, err)
	}

	return nil
}

// ResolveProperties resolves all References in a top-level name-to-property map.
func (r *Resource) ResolveProperties(properties map[string]*Property) error {
	for propertyName, property := range properties {
		resolved, err := r.ResolveProperty(property)

		if err != nil {
			return fmt.Errorf("error resolving %s: %w", propertyName, err)
		}

		if resolved {
			continue
		}

		switch property.Type.String() {
		case PropertyTypeArray:
			_, err = r.ResolveProperty(property.Items)

			if err != nil {
				return fmt.Errorf("error resolving %s Items: %w", propertyName, err)
			}
		case PropertyTypeObject:
			for objPropertyName, objProperty := range property.Properties {
				resolved, err := r.ResolveProperty(objProperty)

				if err != nil {
					return fmt.Errorf("error resolving %s Property (%s): %w", propertyName, objPropertyName, err)
				}

				if resolved {
					continue
				}

				switch objProperty.Type.String() {
				case PropertyTypeArray:
					_, err = r.ResolveProperty(objProperty.Items)

					if err != nil {
						return fmt.Errorf("error resolving %s Property (%s) Items: %w", propertyName, objPropertyName, err)
					}

				case PropertyTypeObject:
					for pattern, patternProperty := range objProperty.PatternProperties {
						_, err = r.ResolveProperty(patternProperty)

						if err != nil {
							return fmt.Errorf("error resolving %s Property (%s) Pattern(%s): %w", propertyName, objPropertyName, pattern, err)
						}
					}
				}
			}

			for patternName, objProperty := range property.PatternProperties {
				resolved, err := r.ResolveProperty(objProperty)

				if err != nil {
					return fmt.Errorf("error resolving %s pattern Property (%s): %w", propertyName, patternName, err)
				}

				if resolved {
					continue
				}

				switch objProperty.Type.String() {
				case PropertyTypeArray:
					_, err = r.ResolveProperty(objProperty.Items)

					if err != nil {
						return fmt.Errorf("error resolving %s Property (%s) Items: %w", propertyName, patternName, err)
					}
				}
			}
		}
	}

	return nil
}

// ResolveProperty resolves any Reference (JSON Pointer) in a Property.
// Returns whether a Reference was resolved.
func (r *Resource) ResolveProperty(property *Property) (bool, error) {
	if property != nil && property.Ref != nil {
		ref := property.Ref
		resolution, err := r.ResolveReference(*ref)

		if err != nil {
			return false, err
		}

		*property = *resolution

		return true, nil
	}

	return false, nil
}
