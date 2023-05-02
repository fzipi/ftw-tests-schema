// Copyright 2023 Felipe Zipitria
// SPDX-License-Identifier: Apache-2.0

//go:build mage
// +build mage

package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var addLicenseVersion = "v1.1.1" // https://github.com/google/addlicense
var golangCILintVer = "v1.52.2"  // https://github.com/golangci/golangci-lint/releases
var gosImportsVer = "v0.3.8"     // https://github.com/rinchsan/gosimports/releases/tag/v0.3.8

var errRunGoModTidy = errors.New("go.mod/sum not formatted, commit changes")
var errNoGitDir = errors.New("no .git directory found")
var errUpdateGeneratedFiles = errors.New("generated files need to be updated")

// Format formats code in this repository.
func Format() error {
	if err := sh.RunV("go", "generate", "./..."); err != nil {
		return err
	}

	if err := sh.RunV("go", "mod", "tidy"); err != nil {
		return err
	}

	// addlicense strangely logs skipped files to stderr despite not being erroneous, so use the long sh.Exec form to
	// discard stderr too.
	if _, err := sh.Exec(map[string]string{}, io.Discard, io.Discard, "go", "run", fmt.Sprintf("github.com/google/addlicense@%s", addLicenseVersion),
		"-c", "Felipe Zipitria",
		"-s=only",
		"-ignore", "**/*.yml",
		"-ignore", "**/*.yaml",
		"-ignore", "v?/**", "."); err != nil {
		return err
	}
	return sh.RunV("go", "run", fmt.Sprintf("github.com/rinchsan/gosimports/cmd/gosimports@%s", gosImportsVer),
		"-w",
		"-local",
		"github.com/fzipi/ftw-tests-yaml-schema",
		".")
}

// Lint verifies code quality.
func Lint() error {
	if err := sh.RunV("go", "generate", "./..."); err != nil {
		return err
	}

	if sh.Run("git", "diff", "--exit-code", "--", "'*.gen.go'") != nil {
		return errUpdateGeneratedFiles
	}

	if err := sh.RunV("go", "run", fmt.Sprintf("github.com/golangci/golangci-lint/cmd/golangci-lint@%s", golangCILintVer), "run"); err != nil {
		return err
	}

	if err := sh.RunV("go", "mod", "tidy"); err != nil {
		return err
	}

	if sh.Run("git", "diff", "--exit-code", "go.mod", "go.sum") != nil {
		return errRunGoModTidy
	}

	return nil
}

// Test runs all tests.
func Test() error {
	if err := sh.RunV("go", "test", "./..."); err != nil {
		return err
	}

	return nil
}

func Markdown() error {
	if err := sh.RunV("go", "build"); err != nil {
		return err
	}
	markdown, err := sh.Output("./ftw-tests-yaml-schema")
	if err != nil {
		return err
	}
	// Write markdown to file
	fmt.Println(markdown)

	return nil

}

// Check runs lint and tests.
func Check() {
	mg.SerialDeps(Lint, Test)
}
