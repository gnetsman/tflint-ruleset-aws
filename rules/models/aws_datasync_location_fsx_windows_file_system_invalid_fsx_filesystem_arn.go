// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule checks the pattern is valid
type AwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule returns new rule with default attributes
func NewAwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule() *AwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule {
	return &AwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule{
		resourceType:  "aws_datasync_location_fsx_windows_file_system",
		attributeName: "fsx_filesystem_arn",
		max:           128,
		pattern:       regexp.MustCompile(`^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):fsx:[a-z\-0-9]*:[0-9]{12}:file-system/fs-.*$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule) Name() string {
	return "aws_datasync_location_fsx_windows_file_system_invalid_fsx_filesystem_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationFsxWindowsFileSystemInvalidFsxFilesystemArnRule) Check(runner tflint.Runner) error {
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
					"fsx_filesystem_arn must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):fsx:[a-z\-0-9]*:[0-9]{12}:file-system/fs-.*$`),
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