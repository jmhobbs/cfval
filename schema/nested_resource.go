package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type NestedResource struct {
	Description string
	Properties  Properties
}

func (res NestedResource) Describe() string {
	return res.Description
}

func (NestedResource) CoercibleTo(PropertyType) Coercion {
	return CoercionNever
}

func (res NestedResource) Validate(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	if values, ok := value.(map[string]interface{}); ok {
		tnr := TemplateNestedResource{
			template:       self.Template(),
			NestedResource: res,
			Properties:     values,
		}
		failures, visited := res.Properties.Validate(tnr, values, context)

		// Reject any properties we weren't expecting
		for key := range res.Properties {
			if !visited[key] {
				failures = append(failures, reporting.NewFailure(fmt.Sprintf("Unknown property '%s' for nested %s", key, res.Description), append(context, key)))
			}
		}

		if len(failures) == 0 {
			return reporting.ValidateOK, nil
		}

		return reporting.ValidateOK, failures
	}

	return reporting.ValidateOK, reporting.Failures{reporting.NewFailure(fmt.Sprintf("Invalid type %T for nested resource %s", value, res.Description), context)}
}