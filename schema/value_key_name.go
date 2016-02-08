package schema

import (
	"github.com/jagregory/cfval/reporting"
)

var KeyName FuncType = func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, []reporting.Failure) {
	if result, errs := ValueString.Validate(property, value, self, context); result == reporting.ValidateAbort || errs != nil {
		return reporting.ValidateOK, errs
	}

	// TODO: KeyName validation
	return reporting.ValidateOK, nil
}
