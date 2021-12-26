// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsServicecatalogProductInvalidTypeRule checks the pattern is valid
type AwsServicecatalogProductInvalidTypeRule struct {
	resourceType  string
	attributeName string
	max           int
	enum          []string
}

// NewAwsServicecatalogProductInvalidTypeRule returns new rule with default attributes
func NewAwsServicecatalogProductInvalidTypeRule() *AwsServicecatalogProductInvalidTypeRule {
	return &AwsServicecatalogProductInvalidTypeRule{
		resourceType:  "aws_servicecatalog_product",
		attributeName: "type",
		max:           8191,
		enum: []string{
			"CLOUD_FORMATION_TEMPLATE",
			"MARKETPLACE",
		},
	}
}

// Name returns the rule name
func (r *AwsServicecatalogProductInvalidTypeRule) Name() string {
	return "aws_servicecatalog_product_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsServicecatalogProductInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsServicecatalogProductInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsServicecatalogProductInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsServicecatalogProductInvalidTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"type must be 8191 characters or less",
					attribute.Expr,
				)
			}
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}