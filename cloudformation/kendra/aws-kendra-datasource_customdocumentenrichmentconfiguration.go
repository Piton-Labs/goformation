package kendra

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// DataSource_CustomDocumentEnrichmentConfiguration AWS CloudFormation Resource (AWS::Kendra::DataSource.CustomDocumentEnrichmentConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.html
type DataSource_CustomDocumentEnrichmentConfiguration struct {

	// InlineConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.html#cfn-kendra-datasource-customdocumentenrichmentconfiguration-inlineconfigurations
	InlineConfigurations []DataSource_InlineCustomDocumentEnrichmentConfiguration `json:"InlineConfigurations,omitempty"`

	// PostExtractionHookConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.html#cfn-kendra-datasource-customdocumentenrichmentconfiguration-postextractionhookconfiguration
	PostExtractionHookConfiguration *DataSource_HookConfiguration `json:"PostExtractionHookConfiguration,omitempty"`

	// PreExtractionHookConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.html#cfn-kendra-datasource-customdocumentenrichmentconfiguration-preextractionhookconfiguration
	PreExtractionHookConfiguration *DataSource_HookConfiguration `json:"PreExtractionHookConfiguration,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.html#cfn-kendra-datasource-customdocumentenrichmentconfiguration-rolearn
	RoleArn string `json:"RoleArn,omitempty"`

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
func (r *DataSource_CustomDocumentEnrichmentConfiguration) AWSCloudFormationType() string {
	return "AWS::Kendra::DataSource.CustomDocumentEnrichmentConfiguration"
}
