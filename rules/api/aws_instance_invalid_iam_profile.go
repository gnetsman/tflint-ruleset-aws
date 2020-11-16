// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsInstanceInvalidIAMProfileRule checks whether attribute value actually exists
type AwsInstanceInvalidIAMProfileRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsInstanceInvalidIAMProfileRule returns new rule with default attributes
func NewAwsInstanceInvalidIAMProfileRule() *AwsInstanceInvalidIAMProfileRule {
	return &AwsInstanceInvalidIAMProfileRule{
		resourceType:  "aws_instance",
		attributeName: "iam_instance_profile",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsInstanceInvalidIAMProfileRule) Name() string {
	return "aws_instance_invalid_iam_profile"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInstanceInvalidIAMProfileRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsInstanceInvalidIAMProfileRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsInstanceInvalidIAMProfileRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by ListInstanceProfiles
func (r *AwsInstanceInvalidIAMProfileRule) Check(rr tflint.Runner) error {
    runner := rr.(*aws.Runner)

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking ListInstanceProfiles")
			var err error
			r.data, err = runner.AwsClient.ListInstanceProfiles()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking ListInstanceProfiles",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.data[val] {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is invalid IAM profile name.`, val),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
