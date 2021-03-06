package ec2

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-mapping.html
var ec2BlockDeviceMapping = NestedResource{
	Description: "EC2 Block Device Mapping",
	Properties: Properties{
		"DeviceName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Ebs": Schema{
			Type:      ec2EbsBlockDevice,
			Required:  constraints.PropertyNotExists("VirtualName"),
			Conflicts: constraints.PropertyExists("VirtualName"),
		},

		"NoDevice": Schema{
			Type: JSON, // TODO: This should actually always be an empty map
		},

		"VirtualName": Schema{
			Type:      ValueString,
			Required:  constraints.PropertyNotExists("Ebs"),
			Conflicts: constraints.PropertyExists("Ebs"),
			ValidateFunc: RegexpValidate(
				"^ephemeral\\d+$",
				"The name must be in the form ephemeralX where X is a number starting from zero (0), for example, ephemeral0",
			),
		},
	},
}
