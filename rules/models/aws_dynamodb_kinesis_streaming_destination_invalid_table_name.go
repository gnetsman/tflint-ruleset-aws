// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule checks the pattern is valid
type AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule returns new rule with default attributes
func NewAwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule() *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule {
	return &AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule{
		resourceType:  "aws_dynamodb_kinesis_streaming_destination",
		attributeName: "table_name",
		max:           255,
		min:           3,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`),
	}
}

// Name returns the rule name
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule) Name() string {
	return "aws_dynamodb_kinesis_streaming_destination_invalid_table_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidTableNameRule) Check(runner tflint.Runner) error {
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
					"table_name must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"table_name must be 3 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_.-]+$`),
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