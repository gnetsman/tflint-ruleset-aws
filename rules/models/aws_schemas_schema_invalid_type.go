// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSchemasSchemaInvalidTypeRule checks the pattern is valid
type AwsSchemasSchemaInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSchemasSchemaInvalidTypeRule returns new rule with default attributes
func NewAwsSchemasSchemaInvalidTypeRule() *AwsSchemasSchemaInvalidTypeRule {
	return &AwsSchemasSchemaInvalidTypeRule{
		resourceType:  "aws_schemas_schema",
		attributeName: "type",
		enum: []string{
			"OpenApi3",
		},
	}
}

// Name returns the rule name
func (r *AwsSchemasSchemaInvalidTypeRule) Name() string {
	return "aws_schemas_schema_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSchemasSchemaInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSchemasSchemaInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSchemasSchemaInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSchemasSchemaInvalidTypeRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
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