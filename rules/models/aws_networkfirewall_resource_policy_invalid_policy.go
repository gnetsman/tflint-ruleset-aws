// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsNetworkfirewallResourcePolicyInvalidPolicyRule checks the pattern is valid
type AwsNetworkfirewallResourcePolicyInvalidPolicyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsNetworkfirewallResourcePolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsNetworkfirewallResourcePolicyInvalidPolicyRule() *AwsNetworkfirewallResourcePolicyInvalidPolicyRule {
	return &AwsNetworkfirewallResourcePolicyInvalidPolicyRule{
		resourceType:  "aws_networkfirewall_resource_policy",
		attributeName: "policy",
		max:           395000,
		min:           1,
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsNetworkfirewallResourcePolicyInvalidPolicyRule) Name() string {
	return "aws_networkfirewall_resource_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsNetworkfirewallResourcePolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsNetworkfirewallResourcePolicyInvalidPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsNetworkfirewallResourcePolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsNetworkfirewallResourcePolicyInvalidPolicyRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"policy must be 395000 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"policy must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}