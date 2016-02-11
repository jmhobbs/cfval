package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

// TODO: better name for this. It's either a TemplateResource or a "NestedTemplateResource"
type SelfRepresentation interface {
	Template() *Template
	Property(name string) (interface{}, bool)
}

type Properties map[string]Schema

func (p Properties) Validate(self SelfRepresentation, values map[string]interface{}, context []string) (reporting.Failures, map[string]bool) {
	failures := make(reporting.Failures, 0, len(p)*2)
	visited := make(map[string]bool)

	for key, schema := range p {
		visited[key] = true
		value, _ := values[key]

		// Validate conflicting properties
		if value != nil && schema.Conflicts != nil && schema.Conflicts.Pass(values) {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Conflict: %s", schema.Conflicts.Describe(values)), append(context, key)))
		}

		// Validate Required
		if value == nil && schema.Required != nil && schema.Required.Pass(values) {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Property is required: %s", schema.Required.Describe(values)), append(context, key)))
		}

		// assuming the above either failed and logged some failures, or passed and
		// we can safely skip over a nil property
		if value == nil {
			continue
		}

		if _, errs := schema.Validate(value, self, append(context, key)); errs != nil {
			failures = append(failures, errs...)
		}
	}

	return failures, visited
}

func collectKeys(m1 map[string]Schema, m2 map[string]interface{}) []string {
	set := make(map[string]bool)
	for key := range m1 {
		set[key] = true
	}
	for key := range m2 {
		set[key] = true
	}

	keys := make([]string, len(set))

	i := 0
	for k := range set {
		keys[i] = k
		i++
	}

	return keys
}