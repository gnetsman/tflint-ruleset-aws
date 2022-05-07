// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsInstanceInvalidSubnetRule checks whether attribute value actually exists
type AwsInstanceInvalidSubnetRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsInstanceInvalidSubnetRule returns new rule with default attributes
func NewAwsInstanceInvalidSubnetRule() *AwsInstanceInvalidSubnetRule {
	return &AwsInstanceInvalidSubnetRule{
		resourceType:  "aws_instance",
		attributeName: "subnet_id",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsInstanceInvalidSubnetRule) Name() string {
	return "aws_instance_invalid_subnet"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInstanceInvalidSubnetRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsInstanceInvalidSubnetRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsInstanceInvalidSubnetRule) Link() string {
	return ""
}

// Metadata returns the metadata about deep checking
func (r *AwsInstanceInvalidSubnetRule) Metadata() interface{} {
	return map[string]bool{"deep": true}
}

// Check checks whether the attributes are included in the list retrieved by DescribeSubnets
func (r *AwsInstanceInvalidSubnetRule) Check(rr tflint.Runner) error {
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
			logger.Debug("invoking DescribeSubnets")
			r.data, err = awsClient.DescribeSubnets()
			if err != nil {
				err := fmt.Errorf("An error occurred while invoking DescribeSubnets; %w", err)
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
					fmt.Sprintf(`"%s" is invalid subnet ID.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	}

	return nil
}
