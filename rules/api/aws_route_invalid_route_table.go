// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsRouteInvalidRouteTableRule checks whether attribute value actually exists
type AwsRouteInvalidRouteTableRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsRouteInvalidRouteTableRule returns new rule with default attributes
func NewAwsRouteInvalidRouteTableRule() *AwsRouteInvalidRouteTableRule {
	return &AwsRouteInvalidRouteTableRule{
		resourceType:  "aws_route",
		attributeName: "route_table_id",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsRouteInvalidRouteTableRule) Name() string {
	return "aws_route_invalid_route_table"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRouteInvalidRouteTableRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRouteInvalidRouteTableRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRouteInvalidRouteTableRule) Link() string {
	return ""
}

// Metadata returns the metadata about deep checking
func (r *AwsRouteInvalidRouteTableRule) Metadata() interface{} {
	return map[string]bool{"deep": true}
}

// Check checks whether the attributes are included in the list retrieved by DescribeRouteTables
func (r *AwsRouteInvalidRouteTableRule) Check(rr tflint.Runner) error {
	runner := rr.(*aws.Runner)

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
			{Name: "provider"},
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

		if !r.dataPrepared {
			awsClient, err := runner.AwsClient(resource.Body.Attributes)
			if err != nil {
				return err
			}
			logger.Debug("invoking DescribeRouteTables")
			r.data, err = awsClient.DescribeRouteTables()
			if err != nil {
				err := fmt.Errorf("An error occurred while invoking DescribeRouteTables; %w", err)
				logger.Error("%s", err)
				return err
			}
			r.dataPrepared = true
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid route table ID.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	}

	return nil
}
