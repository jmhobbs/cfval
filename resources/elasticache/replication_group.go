package elasticache

import (
	"strconv"
	"strings"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

func automaticFailoverEnabledValidation(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	if version, found := self.Property("EngineVersion"); found {
		if versionNumber, err := strconv.ParseFloat(version.(string), 64); err == nil {
			if versionNumber < 2.8 {
				return reporting.ValidateOK, reporting.Failures{reporting.NewFailure("EngineVersion must be 2.8 or higher for Automatic Failover", context)}
			}
		}
	}

	if nodeType, found := self.Property("CacheNodeType"); found {
		split := strings.Split(nodeType.(string), ".")
		if split[1] == "t1" || split[1] == "t2" {
			return reporting.ValidateOK, reporting.Failures{reporting.NewFailure("CacheNodeType must not be T1 or T2 Automatic Failover", context)}
		}
	}

	return reporting.ValidateOK, nil
}

func ReplicationGroup() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::ReplicationGroup",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"AutomaticFailoverEnabled": Schema{
				Type:    ValueBool,
				Default: true,

				// You cannot enable automatic failover for Redis versions earlier than 2.8.6 or for T1 and T2 cache node types.
				ValidateFunc: automaticFailoverEnabledValidation,
			},

			// Currently, this property isn't used by ElastiCache.
			"AutoMinorVersionUpgrade": Schema{
				Type: ValueBool,
			},

			"CacheNodeType": Schema{
				Type:     cacheNodeType,
				Required: constraints.Always,
			},

			"CacheParameterGroupName": Schema{
				Type: ValueString,
			},

			"CacheSecurityGroupNames": Schema{
				Type:      ValueString,
				Array:     true,
				Conflicts: constraints.PropertyExists("SecurityGroupIds"),
			},

			"CacheSubnetGroupName": Schema{
				Type: ValueString,
			},

			"Engine": Schema{
				Type:     engine,
				Required: constraints.Always,
				// ValidateFunc: SingleValue("redis"),
			},

			"EngineVersion": Schema{
				Type: ValueString,
			},

			"NotificationTopicArn": Schema{
				Type: ValueString,
			},

			// If automatic failover is enabled, you must specify a value greater than 1.
			"NumCacheClusters": Schema{
				Type:     ValueNumber,
				Required: constraints.Always,
			},

			"Port": Schema{
				Type: ValueNumber,
			},

			"PreferredCacheClusterAZs": Schema{
				Type:  AvailabilityZone,
				Array: true,
			},

			// Use the following format to specify a time range: ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). For example, you can specify sun:22:00-sun:23:30 for Sunday from 10 PM to 11:30 PM.
			"PreferredMaintenanceWindow": Schema{
				Type: ValueString,
			},

			"ReplicationGroupDescription": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"SecurityGroupIds": Schema{
				Type:      ValueString,
				Array:     true,
				Conflicts: constraints.PropertyExists("CacheSecurityGroupNames"),
			},

			// A single-element string list that specifies an ARN of a Redis .rdb snapshot file that is stored in Amazon Simple Storage Service (Amazon S3). The snapshot file populates the node group. The Amazon S3 object name in the ARN cannot contain commas. For example, you can specify arn:aws:s3:::my_bucket/snapshot1.rdb.
			"SnapshotArns": Schema{
				Type:  ValueString,
				Array: true,
			},

			"SnapshotRetentionLimit": Schema{
				Type: ValueNumber,
			},

			// The time range (in UTC) when ElastiCache takes a daily snapshot of your node group. For example, you can specify 05:00-09:00.
			"SnapshotWindow": Schema{
				Type: ValueString,
			},
		},
	}
}
