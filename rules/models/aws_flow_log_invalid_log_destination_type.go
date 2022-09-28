// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFlowLogInvalidLogDestinationTypeRule checks the pattern is valid
type AwsFlowLogInvalidLogDestinationTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsFlowLogInvalidLogDestinationTypeRule returns new rule with default attributes
func NewAwsFlowLogInvalidLogDestinationTypeRule() *AwsFlowLogInvalidLogDestinationTypeRule {
	return &AwsFlowLogInvalidLogDestinationTypeRule{
		resourceType:  "aws_flow_log",
		attributeName: "log_destination_type",
		enum: []string{
			"cloud-watch-logs",
			"s3",
			"kinesis-data-firehose",
		},
	}
}

// Name returns the rule name
func (r *AwsFlowLogInvalidLogDestinationTypeRule) Name() string {
	return "aws_flow_log_invalid_log_destination_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFlowLogInvalidLogDestinationTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFlowLogInvalidLogDestinationTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFlowLogInvalidLogDestinationTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFlowLogInvalidLogDestinationTypeRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

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
					fmt.Sprintf(`"%s" is an invalid value as log_destination_type`, truncateLongMessage(val)),
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
