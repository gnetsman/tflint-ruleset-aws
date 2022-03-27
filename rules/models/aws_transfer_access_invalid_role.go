// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferAccessInvalidRoleRule checks the pattern is valid
type AwsTransferAccessInvalidRoleRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsTransferAccessInvalidRoleRule returns new rule with default attributes
func NewAwsTransferAccessInvalidRoleRule() *AwsTransferAccessInvalidRoleRule {
	return &AwsTransferAccessInvalidRoleRule{
		resourceType:  "aws_transfer_access",
		attributeName: "role",
		max:           2048,
		min:           20,
		pattern:       regexp.MustCompile(`^arn:.*role/.*$`),
	}
}

// Name returns the rule name
func (r *AwsTransferAccessInvalidRoleRule) Name() string {
	return "aws_transfer_access_invalid_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferAccessInvalidRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferAccessInvalidRoleRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferAccessInvalidRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferAccessInvalidRoleRule) Check(runner tflint.Runner) error {
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
					"role must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role must be 20 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:.*role/.*$`),
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