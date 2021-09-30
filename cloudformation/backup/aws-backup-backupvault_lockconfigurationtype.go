package backup

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// BackupVault_LockConfigurationType AWS CloudFormation Resource (AWS::Backup::BackupVault.LockConfigurationType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupvault-lockconfigurationtype.html
type BackupVault_LockConfigurationType struct {

	// changeableForDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupvault-lockconfigurationtype.html#cfn-backup-backupvault-lockconfigurationtype-changeablefordays
	changeableForDays float64 `json:"changeableForDays,omitempty"`

	// maxRetentionDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupvault-lockconfigurationtype.html#cfn-backup-backupvault-lockconfigurationtype-maxretentiondays
	maxRetentionDays float64 `json:"maxRetentionDays,omitempty"`

	// minRetentionDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupvault-lockconfigurationtype.html#cfn-backup-backupvault-lockconfigurationtype-minretentiondays
	minRetentionDays float64 `json:"minRetentionDays,omitempty"`

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
func (r *BackupVault_LockConfigurationType) AWSCloudFormationType() string {
	return "AWS::Backup::BackupVault.LockConfigurationType"
}
