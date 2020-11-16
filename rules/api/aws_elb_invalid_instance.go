// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsELBInvalidInstanceRule checks whether attribute value actually exists
type AwsELBInvalidInstanceRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsELBInvalidInstanceRule returns new rule with default attributes
func NewAwsELBInvalidInstanceRule() *AwsELBInvalidInstanceRule {
	return &AwsELBInvalidInstanceRule{
		resourceType:  "aws_elb",
		attributeName: "instances",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsELBInvalidInstanceRule) Name() string {
	return "aws_elb_invalid_instance"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsELBInvalidInstanceRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsELBInvalidInstanceRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsELBInvalidInstanceRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by DescribeInstances
func (r *AwsELBInvalidInstanceRule) Check(rr tflint.Runner) error {
    runner := rr.(*aws.Runner)

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeInstances")
			var err error
			r.data, err = runner.AwsClient.DescribeInstances()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeInstances",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		return runner.EachStringSliceExprs(attribute.Expr, func(val string, expr hcl.Expression) {
			if !r.data[val] {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is invalid instance.`, val),
					expr,
				)
			}
		})
	})
}
