package resources

import (
	"fmt"
	"strconv"

	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var period FuncType = func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, []reporting.Failure) {
	if result, errs := ValueString.Validate(property, value, self, context); result == reporting.ValidateAbort || errs != nil {
		return reporting.ValidateOK, errs
	}

	num, err := strconv.Atoi(value.(string))
	if err != nil {
		return reporting.ValidateOK, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Period is not a number: %s", value), context)}
	}

	if num == 0 || num%60 != 0 {
		return reporting.ValidateOK, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Period is not a multiple of 60: %s", value), context)}
	}

	return reporting.ValidateOK, nil
}
