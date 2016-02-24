package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/pseudo-parameter-reference.html
var pseudoParameters = map[string]Schema{
	"AWS::AccountId": Schema{
		Type: ValueString,
	},

	"AWS::NotificationARNs": Schema{
		Type: Multiple(ValueString), // TODO: ARN loop here
	},

	"AWS::NoValue": Schema{
		Type: ValueString,
	},

	"AWS::Region": Schema{
		Type: ValueString,
	},

	"AWS::StackId": Schema{
		Type: ValueString,
	},

	"AWS::StackName": Schema{
		Type: ValueString,
	},
}

type RefTarget interface {
	TargetType() PropertyType
}

type Ref struct {
	target string
}

func NewRef(target string) Ref {
	return Ref{target}
}

func (ref Ref) Validate(ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if ctx.Template == nil {
		panic("Template is nil")
	}

	target := ref.resolveTarget(ctx.Definitions(), ctx.Template())
	if target == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Ref '%s' is not a resource, parameter, or pseudo-parameter", ref.target)}
	}

	targetType := target.TargetType()
	if targetType == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "%s cannot be used in a Ref", ref.target)}
	}

	switch targetType.CoercibleTo(ctx.Property().Type) {
	case CoercionNever:
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Ref value of '%s' is %s but is being assigned to a %s property", ref.target, targetType.Describe(), ctx.Property().Type.Describe())}
	case CoercionBegrudgingly:
		return reporting.ValidateAbort, reporting.Reports{reporting.NewWarning(ctx, "Ref value of '%s' is %s but is being dangerously coerced to a %s property", ref.target, targetType.Describe(), ctx.Property().Type.Describe())}
	}

	// fmt.Printf("Ref value of '%s' is %s and is coerced to a %s property", ref.target, targetType.Describe(), ctx.Property().Type.Describe())
	return reporting.ValidateAbort, nil
}

func (ref Ref) resolveTarget(definitions ResourceDefinitions, template *parse.Template) RefTarget {
	if resource, ok := template.Resources[ref.target]; ok {
		return definitions.Lookup(resource.Type)
	} else if parameter, ok := template.Parameters[ref.target]; ok {
		return definitions.LookupParameter(parameter.Type)
	} else if pseudoParameters, ok := pseudoParameters[ref.target]; ok {
		return pseudoParameters
	}

	return nil
}
