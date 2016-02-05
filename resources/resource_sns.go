package resources

import . "github.com/jagregory/cfval/schema"

func Topic() Resource {
	return Resource{
		AwsType: "AWS::SNS::Topic",

		// Topic ARN
		ReturnValue: Schema{
			Type: TypeString,
		},

		Properties: map[string]Schema{
			"DisplayName": Schema{
				Type: TypeString,
			},

			// "Subscription": ArrayOf(Schema{
			// 	Type: Resource{...}
			// }),

			"TopicName": Schema{
				Type: TypeString,
			},
		},
	}
}
