package schema

import (
	"fmt"
	"strconv"

	"github.com/jagregory/cfval/reporting"
)

func ValidateBuiltinFns(s Schema, value map[string]interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	if ref, ok := value["Ref"]; ok {
		return NewRef(s, ref.(string)).Validate(self.Template(), append(context, "Ref"))
	}

	if join, ok := value["Fn::Join"]; ok {
		return validateJoin(join, self, append(context, "Fn::Join"))
	}

	if getatt, ok := value["Fn::GetAtt"]; ok {
		return validateGetAtt(getatt, self, append(context, "Fn::GetAtt"))
	}

	if find, ok := value["Fn::FindInMap"]; ok {
		return validateFindInMap(find, self, append(context, "Fn::FindInMap"))
	}

	if base64, ok := value["Fn::Base64"]; ok {
		return validateBase64(base64, self, append(context, "Fn::Base64"))
	}

	// not a builtin, but this isn't necessarily bad so we don't return an error here
	return reporting.ValidateOK, nil
}

// TODO: Supported functions within a function
func validateFindInMap(value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	find, ok := value.([]interface{})
	if !ok {
		return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure("Options need to be an array", context)}
	}

	if len(find) != 3 {
		return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("Options has wrong number of items, expected: 3, actual: %d", len(find)), context)}
	}

	mapName := find[0]
	_, mapNameIsString := mapName.(string)
	if _, errs := ValueString.Validate(Schema{Type: ValueString}, mapName, self, append(context, "0")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString {
		// map name is a string, so we can do some further interrogation
		// TODO: lookup whether MapName is a valid Map
	}

	topLevelKey := find[1]
	_, topLevelKeyIsString := topLevelKey.(string)
	if _, errs := ValueString.Validate(Schema{Type: ValueString}, topLevelKey, self, append(context, "1")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString && topLevelKeyIsString {
		// TODO: lookup whether topLevelKey is in mapName
	}

	secondLevelKey := find[2]
	_, secondLevelKeyIsString := secondLevelKey.(string)
	if _, errs := ValueString.Validate(Schema{Type: ValueString}, secondLevelKey, self, append(context, "2")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString && topLevelKeyIsString && secondLevelKeyIsString {
		// TODO: lookup whether secondLevelKeyIsString is in topLevelKey
	}

	return reporting.ValidateAbort, nil
}

func validateBase64(value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	_, errs := ValueString.Validate(Schema{Type: ValueString}, value, self, context)
	return reporting.ValidateAbort, errs
}

func validateJoin(value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	if items, ok := value.([]interface{}); ok {
		if len(items) != 2 {
			return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("Join has incorrect number of arguments (expected: 2, actual: %d)", len(items)), context)}
		}

		_, ok := items[0].(string)
		if !ok {
			return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("Join '%s' is not a valid delimiter", items[0]), context)}
		}

		parts, ok := items[1].([]interface{})
		if !ok {
			return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("Join items are not valid: %s", items[1]), context)}
		}

		failures := make(reporting.Failures, 0, len(parts))
		for i, part := range parts {
			if _, errs := ValueString.Validate(Schema{Type: ValueString}, part, self, append(context, "1", strconv.Itoa(i))); errs != nil {
				failures = append(failures, errs...)
			}
		}

		if len(failures) == 0 {
			return reporting.ValidateAbort, nil
		}

		return reporting.ValidateAbort, failures
	}

	return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("GetAtt has invalid value '%s'", value), context)}
}

func validateGetAtt(value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	if items, ok := value.([]interface{}); ok {
		if len(items) != 2 {
			return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("GetAtt has incorrect number of arguments (expected: 2, actual: %d)", len(items)), context)}
		}

		if resourceID, ok := items[0].(string); ok {
			if _, ok := self.Template().Resources[resourceID]; ok {
				if _, ok := items[1].(string); ok {
					// TODO: Check attr is actually a valid attribute for the resource type
					return reporting.ValidateAbort, nil
				}
			}
			// resource not found
			return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("GetAtt '%s' is not a resource", resourceID), context)}
		}

		// resource not a string
		return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("GetAtt '%s' is not a valid resource name", items[0]), context)}
	}

	return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("GetAtt has invalid value '%s'", value), context)}
}