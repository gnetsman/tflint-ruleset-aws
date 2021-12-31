// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGuarddutyFilterInvalidDescriptionRule checks the pattern is valid
type AwsGuarddutyFilterInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsGuarddutyFilterInvalidDescriptionRule returns new rule with default attributes
func NewAwsGuarddutyFilterInvalidDescriptionRule() *AwsGuarddutyFilterInvalidDescriptionRule {
	return &AwsGuarddutyFilterInvalidDescriptionRule{
		resourceType:  "aws_guardduty_filter",
		attributeName: "description",
		max:           512,
	}
}

// Name returns the rule name
func (r *AwsGuarddutyFilterInvalidDescriptionRule) Name() string {
	return "aws_guardduty_filter_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyFilterInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyFilterInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyFilterInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyFilterInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 512 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
