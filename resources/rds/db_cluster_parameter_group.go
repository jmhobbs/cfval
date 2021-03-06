package rds

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html
var DBClusterParameterGroup = Resource{
	AwsType: "AWS::RDS::DBClusterParameterGroup",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"Description": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Family": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Parameters": Schema{
			Type:     JSON,
			Required: constraints.Always,
		},

		"Tags": Schema{
			Type: Multiple(common.ResourceTag),
		},
	},
}
