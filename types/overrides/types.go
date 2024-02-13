// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

//go:generate dstdocgen -package overrides -path . -structure FTWOverrides -output ./overrides_doc.go

package overrides

import "github.com/coreruleset/ftw-tests-schema/types"

// FTWOverrides describes platform specific overrides for tests
type FTWOverrides struct {
	// description: |
	//   The version field designates the version of the schema that validates this file
	// examples:
	//   - value: "\"v0.1.0\""
	Version string `yaml:"version"`

	// description: |
	//   Meta describes the metadata information
	// examples:
	//   - value: MetaExample
	Meta FTWOverridesMeta `yaml:"meta"`

	// description: |
	//   List of test override specifications
	// examples:
	//   - value: TestOverridesExample
	TestOverrides []TestOverride `yaml:"test_overrides"`
}

// FTWOverridesMeta describes the metadata information of this yaml file
type FTWOverridesMeta struct {
	// description: |
	//   The name of the WAF engine the tests are expected to run against
	// examples:
	//   - value: "\"coraza\""
	Engine string `yaml:"engine"`

	// description: |
	//   The name of the platform (e.g., web server) the tests are expected to run against
	// examples:
	//   - value: "\"nginx\""
	Platform string `yaml:"platform"`

	// description: |
	//   Custom annotations; can be used to add additional meta information
	// examples:
	//   - value: AnnotationsExample
	Annotations map[string]string `yaml:"annotations"`
}

// TestOverride describes overrides for a single test
type TestOverride struct {
	// description: |
	//   ID of the rule this test targets.
	// examples:
	//   - value: TestOverridesExample[0].RuleId
	RuleId int `yaml:"rule_id"`

	// description: |
	//   IDs of the tests for rule_id that overrides should be applied to.
	//   If this field is not set, the overrides will be applied to all tests of rule_id.
	// examples:
	//   - value: TestOverridesExample[0].TestIds
	TestIds []int `yaml:"test_ids,flow,omitempty"`

	// description: |
	//   Describes why this override is necessary.
	// examples:
	//   - value: ReasonExample
	Reason string `yaml:"reason"`

	// description: |
	//   Whether this test is expected to fail for this particular configuration.
	//   Default: false
	// examples:
	//   - value: true
	ExpectFailure *bool `yaml:"expect_failure,omitempty"`

	// description: |
	//   Whether a stage should be retried once in case of failure.
	//   This option is primarily a workaround for a race condition in phase 5,
	//   where the log entry of a rule may be flushed after the test end marker.
	// examples:
	//   - value: true
	RetryOnce *bool `yaml:"retry_once,omitempty"`

	// description: |
	//   Specifies overrides on the test output
	// examples:
	//   - value: 400
	Output types.Output `yaml:"output"`
}
