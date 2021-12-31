// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGuarddutyPublishingDestinationInvalidDetectorIDRule checks the pattern is valid
type AwsGuarddutyPublishingDestinationInvalidDetectorIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsGuarddutyPublishingDestinationInvalidDetectorIDRule returns new rule with default attributes
func NewAwsGuarddutyPublishingDestinationInvalidDetectorIDRule() *AwsGuarddutyPublishingDestinationInvalidDetectorIDRule {
	return &AwsGuarddutyPublishingDestinationInvalidDetectorIDRule{
		resourceType:  "aws_guardduty_publishing_destination",
		attributeName: "detector_id",
		max:           300,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsGuarddutyPublishingDestinationInvalidDetectorIDRule) Name() string {
	return "aws_guardduty_publishing_destination_invalid_detector_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyPublishingDestinationInvalidDetectorIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyPublishingDestinationInvalidDetectorIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyPublishingDestinationInvalidDetectorIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyPublishingDestinationInvalidDetectorIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"detector_id must be 300 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"detector_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
