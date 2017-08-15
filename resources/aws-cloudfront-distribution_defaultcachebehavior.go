package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.DefaultCacheBehavior AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html
type AWSCloudFrontDistribution_DefaultCacheBehavior struct {

	// AllowedMethods AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-allowedmethods

	AllowedMethods []string `json:"AllowedMethods"`

	// CachedMethods AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-cachedmethods

	CachedMethods []string `json:"CachedMethods"`

	// Compress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-compress

	Compress bool `json:"Compress"`

	// DefaultTTL AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-defaultttl

	DefaultTTL int64 `json:"DefaultTTL"`

	// ForwardedValues AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-forwardedvalues

	ForwardedValues AWSCloudFrontDistribution_ForwardedValues `json:"ForwardedValues"`

	// MaxTTL AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-maxttl

	MaxTTL int64 `json:"MaxTTL"`

	// MinTTL AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-minttl

	MinTTL int64 `json:"MinTTL"`

	// SmoothStreaming AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-smoothstreaming

	SmoothStreaming bool `json:"SmoothStreaming"`

	// TargetOriginId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-targetoriginid

	TargetOriginId string `json:"TargetOriginId"`

	// TrustedSigners AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-trustedsigners

	TrustedSigners []string `json:"TrustedSigners"`

	// ViewerProtocolPolicy AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-defaultcachebehavior.html#cfn-cloudfront-defaultcachebehavior-viewerprotocolpolicy

	ViewerProtocolPolicy string `json:"ViewerProtocolPolicy"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_DefaultCacheBehavior) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.DefaultCacheBehavior"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_DefaultCacheBehavior) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_DefaultCacheBehaviorResources retrieves all AWSCloudFrontDistribution_DefaultCacheBehavior items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_DefaultCacheBehavior(template *Template) map[string]*AWSCloudFrontDistribution_DefaultCacheBehavior {

	results := map[string]*AWSCloudFrontDistribution_DefaultCacheBehavior{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_DefaultCacheBehavior{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_DefaultCacheBehaviorWithName retrieves all AWSCloudFrontDistribution_DefaultCacheBehavior items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_DefaultCacheBehavior(name string, template *Template) (*AWSCloudFrontDistribution_DefaultCacheBehavior, error) {

	result := &AWSCloudFrontDistribution_DefaultCacheBehavior{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_DefaultCacheBehavior{}, errors.New("resource not found")

}