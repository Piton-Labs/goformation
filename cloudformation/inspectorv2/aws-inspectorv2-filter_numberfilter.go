package inspectorv2

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Filter_NumberFilter AWS CloudFormation Resource (AWS::InspectorV2::Filter.NumberFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-numberfilter.html
type Filter_NumberFilter struct {

	// LowerInclusive AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-numberfilter.html#cfn-inspectorv2-filter-numberfilter-lowerinclusive
	LowerInclusive float64 `json:"LowerInclusive,omitempty"`

	// UpperInclusive AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-numberfilter.html#cfn-inspectorv2-filter-numberfilter-upperinclusive
	UpperInclusive float64 `json:"UpperInclusive,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Filter_NumberFilter) AWSCloudFormationType() string {
	return "AWS::InspectorV2::Filter.NumberFilter"
}
