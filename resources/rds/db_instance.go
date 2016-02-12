package rds

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

func DBInstance() Resource {
	return Resource{
		AwsType: "AWS::RDS::DBInstance",

		// ID
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"AllocatedStorage": Schema{
				Type: ValueString,
			},

			"AllowMajorVersionUpgrade": Schema{
				Type: ValueBool,
			},

			"AutoMinorVersionUpgrade": Schema{
				Type: ValueBool,
			},

			"AvailabilityZone": Schema{
				Type: ValueString,
			},

			"BackupRetentionPeriod": Schema{
				Type: ValueString,
			},

			"CharacterSetName": Schema{
				Type: ValueString,
			},

			"DBClusterIdentifier": Schema{
				Type: ValueString,
			},

			"DBInstanceClass": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"DBInstanceIdentifier": Schema{
				Type: ValueString,
			},

			"DBName": Schema{
				Type: ValueString,
			},

			"DBParameterGroupName": Schema{
				Type: ValueString,
			},

			"DBSecurityGroups": Schema{
				Type:  ValueString,
				Array: true,
			},

			"DBSnapshotIdentifier": Schema{
				Type: ValueString,
			},

			"DBSubnetGroupName": Schema{
				Type: ValueString,
			},

			"Engine": Schema{
				Type: ValueString,
			},

			"EngineVersion": Schema{
				Type: ValueString,
			},

			"Iops": Schema{
				Type: ValueNumber,
			},

			"KmsKeyId": Schema{
				Type: ValueString,
			},

			"LicenseModel": Schema{
				Type: ValueString,
			},

			"MasterUsername": Schema{
				Type: ValueString,
			},

			"MasterUserPassword": Schema{
				Type: ValueString,
			},

			"MultiAZ": Schema{
				Type: ValueBool,
			},

			"OptionGroupName": Schema{
				Type: ValueString,
			},

			"Port": Schema{
				Type: ValueString,
			},

			"PreferredBackupWindow": Schema{
				Type: ValueString,
			},

			"PreferredMaintenanceWindow": Schema{
				Type: ValueString,
			},

			"PubliclyAccessible": Schema{
				Type: ValueBool,
			},

			"SourceDBInstanceIdentifier": Schema{
				Type: ValueString,
			},

			"StorageEncrypted": Schema{
				Type: ValueBool,
			},

			"StorageType": Schema{
				Type: ValueString,
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},

			"VPCSecurityGroups": Schema{
				Type:  ValueString,
				Array: true,
			},
		},
	}
}