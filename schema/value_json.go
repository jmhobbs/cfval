package schema

import (
	"fmt"
	"strconv"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

var JSON FuncType

func validateJSON(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	switch t := value.(type) {
	case parse.IntrinsicFunction:
		jsonType := Schema{Type: JSON}
		return ValidateIntrinsicFunctions(t, NewPropertyContext(ctx, jsonType), SupportedFunctionsAll)
	case map[string]interface{}:
		return validateJSONMap(t, ctx)
	case []interface{}:
		return validateJSONArray(t, ctx)
	case string:
		stringItemContext := NewPropertyContext(ctx, Schema{Type: ValueString})
		return ValueString.Validate(t, stringItemContext)
	case float64:
		numberItemContext := NewPropertyContext(ctx, Schema{Type: ValueNumber})
		return ValueNumber.Validate(t, numberItemContext)
	case bool:
		boolItemContext := NewPropertyContext(ctx, Schema{Type: ValueBool})
		return ValueBool.Validate(t, boolItemContext)
	default:
		panic(fmt.Sprintf("Unexpected JSON type %T", t))
	}
}

func validateJSONMap(value map[string]interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)
	stringItemContext := NewPropertyContext(ctx, Schema{Type: ValueString})

	for k, v := range value {
		if _, errs := validateJSON(v, PropertyContextAdd(stringItemContext, k)); errs != nil {
			failures = append(failures, errs...)
		}
	}

	if len(failures) == 0 {
		return reporting.ValidateOK, nil
	}

	return reporting.ValidateOK, failures
}

func validateJSONArray(value []interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)

	for i, item := range value {
		if _, errs := validateJSON(item, PropertyContextAdd(ctx, strconv.Itoa(i))); errs != nil {
			failures = append(failures, errs...)
		}
	}

	if len(failures) == 0 {
		return reporting.ValidateOK, nil
	}

	return reporting.ValidateOK, failures
}

func coerceJSON(to PropertyType) Coercion {
	if ft, ok := to.(FuncType); ok && ft.Description == "JSON" {
		return CoercionAlways
	} else if to == ValueUnknown {
		return CoercionBegrudgingly
	}

	return CoercionNever
}

func init() {
	JSON = FuncType{
		Description: "JSON",

		Fn:          validateJSON,
		CoercibleFn: coerceJSON,
	}
}
