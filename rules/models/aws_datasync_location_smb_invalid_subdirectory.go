// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationSmbInvalidSubdirectoryRule checks the pattern is valid
type AwsDatasyncLocationSmbInvalidSubdirectoryRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationSmbInvalidSubdirectoryRule returns new rule with default attributes
func NewAwsDatasyncLocationSmbInvalidSubdirectoryRule() *AwsDatasyncLocationSmbInvalidSubdirectoryRule {
	return &AwsDatasyncLocationSmbInvalidSubdirectoryRule{
		resourceType:  "aws_datasync_location_smb",
		attributeName: "subdirectory",
		max:           4096,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationSmbInvalidSubdirectoryRule) Name() string {
	return "aws_datasync_location_smb_invalid_subdirectory"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationSmbInvalidSubdirectoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationSmbInvalidSubdirectoryRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationSmbInvalidSubdirectoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationSmbInvalidSubdirectoryRule) Check(runner tflint.Runner) error {
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
					"subdirectory must be 4096 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`),
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