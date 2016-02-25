package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getavailabilityzones.html
func validateGetAZs(builtin parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	value, found := builtin.UnderlyingMap["Fn::GetAZs"]
	if !found || value == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::GetAZs\" key")}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, "Fn::GetAZs"))}
	}

	switch t := value.(type) {
	case string:
		return reporting.ValidateAbort, nil
	case parse.IntrinsicFunction:
		if t.Key == parse.FnRef {
			return validateRef(t, PropertyContextAdd(ctx, string(parse.FnRef)))
		}
	}

	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Fn::GetAZs\" key: %s", value)}
}