package overrides

import "github.com/coreruleset/ftw-tests-schema/test"

// Copyright 2023 Felipe Zipitria
// SPDX-License-Identifier: Apache-2.0

//go:generate dstdocgen -package overrides -path . -structure FTWOverrides -output ./overrides_doc.go

var (
	metaExample = Meta{
		Engine:      "libmodsecurity3",
		Platform:    "nginx",
		Annotations: annotationsExample,
	}
	annotationsExample = map[string]string{
		"os":      "Debian Bullseye",
		"purpose": "Bullseye integration",
	}
	reasonExample = "nginx returns 400 when `Content-Length` header is sent in a\n" +
		"`Transfer-Encoding: chunked` request."

	testOverridesExample = []TestOverride{
		TestOverride{
			RuleId:        920100,
			TestId:        4,
			Reason:        reasonExample,
			ExpectFailure: true,
			Output:        test.ExampleOutput,
		},
	}
)

// TODO: Welcome to the FTW YAMLFormat documentation.
// In this document we will explain all the possible options that can be used within the YAML format.
// Generally this is the preferred format for writing tests in as they don't require any programming skills
// in order to understand and change. If you find a bug in this format please open an issue.

// TODO: FTWTest is the base type used when unmarshaling YAML tests files
type FTWOverrides struct {
	// description: |
	//    The version field designates the version of the schema that validates this file
	// examples:
	//     - value: "\"v0.1.0\""
	Version string `yaml:"version"`

	// description: |
	//    Meta describes the metadata information
	// examples:
	//     - value: metaExample
	Meta Meta `yaml:"meta"`

	// description: |
	//    List of test override specifications
	// examples:
	//     - value: testOverridesExample
	TestOverrides []TestOverride `yaml:"test_overrides"`
}

// Meta describes the metadata information of this yaml file
type Meta struct {
	// description: |
	//    The name of the WAF engine the tests are expected to run against
	// examples:
	//    - value: "\"coraza\""
	Engine string `yaml:"engine"`

	// description: |
	//    The name of the platform (e.g., web server) the tests are expected to run against
	// examples:
	//    - value: "\"nginx\""
	Platform string `yaml:"platform"`

	// description: |
	//     Custom annotations; can be used to add additional meta information
	// examples:
	//     - value: annotationsExample
	Annotations map[string]string `yaml:"annotations"`
}

// TestOverride describes overrides for a single test
type TestOverride struct {
	// description: |
	//     ID of the rule this test targets.
	//     If this field is not empty, `test_id` must also be set.
	// examples:
	//     - value: "\"920100\""
	RuleId int `yaml:"rule_id,omitempty"`

	// description: |
	//     ID of the test this override applies to.
	//     If this field is not empty, `rule_id` must also be set.
	// examples:
	//     - value: 5
	TestId int `yaml:"test_id,omitempty"`

	// description: |
	//     Regular expression matching test names (can match multiple tests).
	//     If this field is empty, `rule_id` and `test_id` must be set.
	// examples:
	//     - value: "\"^910.*$\""
	NameRegex string `yaml:"name_regex,omitempty"`

	// description: |
	//    Describes why this override is necessary.
	// examples:
	//     - value: reasonExample
	Reason string `yaml:"reason"`

	// description: |
	//     Whether this test is expected to fail for this particular configuration.
	//     Default: false
	// examples:
	//     - value: true
	ExpectFailure bool `yaml:"expect_failure,omitempty"`

	// description: |
	//     Specifies overrides on the test output
	// examples:
	//     - value: 400
	Output test.Output `yaml:"output"`
}
