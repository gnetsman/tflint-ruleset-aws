// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferAccessInvalidExternalIDRule checks the pattern is valid
type AwsTransferAccessInvalidExternalIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsTransferAccessInvalidExternalIDRule returns new rule with default attributes
func NewAwsTransferAccessInvalidExternalIDRule() *AwsTransferAccessInvalidExternalIDRule {
	return &AwsTransferAccessInvalidExternalIDRule{
		resourceType:  "aws_transfer_access",
		attributeName: "external_id",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^S-1-[\d-]+$`),
	}
}

// Name returns the rule name
func (r *AwsTransferAccessInvalidExternalIDRule) Name() string {
	return "aws_transfer_access_invalid_external_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferAccessInvalidExternalIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferAccessInvalidExternalIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferAccessInvalidExternalIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferAccessInvalidExternalIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"external_id must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"external_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^S-1-[\d-]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}