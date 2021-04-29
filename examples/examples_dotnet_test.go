// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
// +build dotnet all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccRecordCSharp(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			RunUpdateTest: false, // this curerntly doesn't work in the released csharp version
			Dir:           filepath.Join(getCwd(t), "record", "csharp"),
		})

	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseCsharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.DNSimple",
		},
	})

	return baseCsharp
}
