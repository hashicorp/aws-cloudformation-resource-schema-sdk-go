package cfschema

import (
	"fmt"
)

// Expand replaces all Definition and Property JSON Pointer references with their content.
//
// This functionality removes the need for recursive logic when accessing
// Definition and Property.
func (r *Resource) Expand() error {
	if r == nil {
		return nil
	}

	for definitionName, definition := range r.Definitions {
		resolved, err := r.Resolve(definition)

		if err != nil {
			return fmt.Errorf("error resolving Definition (%s): %w", definitionName, err)
		}

		if resolved {
			continue
		}

		switch definition.Type.String() {
		case PropertyTypeArray:
			_, err = r.Resolve(definition.Items)

			if err != nil {
				return fmt.Errorf("error resolving Definition (%s) Items: %w", definitionName, err)
			}
		case PropertyTypeObject:
			for objPropertyName, objProperty := range definition.Properties {
				resolved, err := r.Resolve(objProperty)

				if err != nil {
					return fmt.Errorf("error resolving Definition (%s) Property (%s): %w", definitionName, objPropertyName, err)
				}

				if resolved {
					continue
				}

				switch objProperty.Type.String() {
				case PropertyTypeArray:
					_, err = r.Resolve(objProperty.Items)

					if err != nil {
						return fmt.Errorf("error resolving Definition (%s) Property (%s) Items: %w", definitionName, objPropertyName, err)
					}
				}
			}
		}
	}

	for propertyName, property := range r.Properties {
		resolved, err := r.Resolve(property)

		if err != nil {
			return fmt.Errorf("error resolving Property (%s): %w", propertyName, err)
		}

		if resolved {
			continue
		}

		switch property.Type.String() {
		case PropertyTypeArray:
			_, err = r.Resolve(property.Items)

			if err != nil {
				return fmt.Errorf("error resolving Property (%s) Items: %w", propertyName, err)
			}
		case PropertyTypeObject:
			for objPropertyName, objProperty := range property.Properties {
				resolved, err := r.Resolve(objProperty)

				if err != nil {
					return fmt.Errorf("error resolving Property (%s) Property (%s): %w", propertyName, objPropertyName, err)
				}

				if resolved {
					continue
				}

				switch objProperty.Type.String() {
				case PropertyTypeArray:
					_, err = r.Resolve(objProperty.Items)

					if err != nil {
						return fmt.Errorf("error resolving Property (%s) Property (%s) Items: %w", propertyName, objPropertyName, err)
					}
				}
			}
		}
	}

	return nil
}

// Resolve resolves any Reference (JSON Pointer) in a Property.
// Returns whether a Reference was resolved.
func (r *Resource) Resolve(property *Property) (bool, error) {
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
