// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule checks the pattern is valid
type AwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule returns new rule with default attributes
func NewAwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule() *AwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule {
	return &AwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule{
		resourceType:  "aws_dynamodb_kinesis_streaming_destination",
		attributeName: "stream_arn",
		max:           1024,
		min:           37,
	}
}

// Name returns the rule name
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule) Name() string {
	return "aws_dynamodb_kinesis_streaming_destination_invalid_stream_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDynamoDBKinesisStreamingDestinationInvalidStreamArnRule) Check(runner tflint.Runner) error {
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
					"stream_arn must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"stream_arn must be 37 characters or higher",
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