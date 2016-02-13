package elasti_cache

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

func azModeValidate(prop Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	if str, ok := value.(string); ok {
		if availabilityZones, ok := self.Property("PreferredAvailabilityZones"); ok {
			if str == "cross-az" && len(availabilityZones.([]interface{})) < 2 {
				return reporting.ValidateOK, reporting.Failures{reporting.NewFailure("Cross-AZ clusters must have multiple preferred availability zones", context)}
			}
		}
	}

	return reporting.ValidateOK, nil
}

func numCacheNodesValidate(prop Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
	if engine, ok := self.Property("Engine"); !ok || engine.(string) == "memcached" {
		return IntegerRangeValidate(1, 20)(prop, value, self, context)
	}

	return SingleValueValidate(float64(1))(prop, value, self, context)
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html
func CacheCluster() Resource {
	return Resource{
		AwsType: "AWS::ElastiCache::CacheCluster",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"AutoMinorVersionUpgrade": Schema{
				Type: ValueBool,
			},

			"AZMode": Schema{
				Type:         azMode,
				ValidateFunc: azModeValidate,
				Default:      "single-az",
			},

			"CacheNodeType": Schema{
				Type:     cacheNodeType,
				Required: constraints.Always,
			},

			"CacheParameterGroupName": Schema{
				Type: ValueString,
			},

			"CacheSecurityGroupNames": Schema{
				Type:  cacheSecurityGroupName,
				Array: true,
				Conflicts: constraints.Any{
					constraints.PropertyExists("CacheSubnetGroupName"),
					constraints.PropertyExists("VpcSecurityGroupIds"),
				},
			},

			"CacheSubnetGroupName": Schema{
				Type: cacheSecurityGroupName,
				Conflicts: constraints.Any{
					constraints.PropertyExists("CacheSecurityGroupNames"),
					constraints.PropertyExists("VpcSecurityGroupIds"),
				},
			},

			"ClusterName": Schema{
				Type: ValueString,
			},

			"Engine": Schema{
				Type:     engine,
				Required: constraints.Always,
			},

			"EngineVersion": Schema{
				Type: ValueString,
			},

			"NotificationTopicArn": Schema{
				Type: ValueString,
			},

			"NumCacheNodes": Schema{
				Type:         ValueNumber,
				Required:     constraints.Always,
				ValidateFunc: numCacheNodesValidate,
			},

			"Port": Schema{
				Type: ValueNumber,
			},

			"PreferredAvailabilityZone": Schema{
				Type: AvailabilityZone,
			},

			"PreferredAvailabilityZones": Schema{
				Type:     AvailabilityZone,
				Array:    true,
				Required: constraints.PropertyIs("AZMode", "cross-az"),
			},

			"PreferredMaintenanceWindow": Schema{
				Type: ValueString,
			},

			"SnapshotArns": Schema{
				Type:  ValueString,
				Array: true,
			},

			"SnapshotName": Schema{
				Type: ValueString,
			},

			"SnapshotRetentionLimit": Schema{
				Type: ValueNumber,
			},

			"SnapshotWindow": Schema{
				Type: ValueString,
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},

			"VpcSecurityGroupIds": Schema{
				Type:      SecurityGroupID,
				Array:     true,
				Conflicts: constraints.PropertyExists("CacheSecurityGroupNames"),
			},
		},
	}
}