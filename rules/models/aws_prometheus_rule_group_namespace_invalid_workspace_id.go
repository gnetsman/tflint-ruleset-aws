// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule checks the pattern is valid
type AwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule returns new rule with default attributes
func NewAwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule() *AwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule {
	return &AwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule{
		resourceType:  "aws_prometheus_rule_group_namespace",
		attributeName: "workspace_id",
		max:           64,
		min:           1,
		pattern:       regexp.MustCompile(`^[0-9A-Za-z][-.0-9A-Z_a-z]*$`),
	}
}

// Name returns the rule name
func (r *AwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule) Name() string {
	return "aws_prometheus_rule_group_namespace_invalid_workspace_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsPrometheusRuleGroupNamespaceInvalidWorkspaceIDRule) Check(runner tflint.Runner) error {
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
					"workspace_id must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"workspace_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9A-Za-z][-.0-9A-Z_a-z]*$`),
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