package iam

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
var policy = NestedResource{
	Description: "IAM Policies",

	Properties: Properties{
		"PolicyDocument": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},

		"PolicyName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html
func Group() Resource {
	return Resource{
		AwsType: "AWS::IAM::Group",

		Attributes: map[string]Schema{
			"Arn": Schema{
				Type: ValueString,
			},
		},

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"ManagedPolicyArns": Schema{
				Type:  ValueString,
				Array: true,
			},

			"Path": Schema{
				Type: ValueString,
			},

			"Policies": Schema{
				Type:  policy,
				Array: true,
			},
		},
	}
}
