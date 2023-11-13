// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

//go:generate dstdocgen -package types -path . -structure FTWTest -output ./test_doc.go
//go:generate dstdocgen -package types -path . -structure FTWOverrides -output ./overrides_doc.go

package types

// Welcome to the FTW YAMLFormat documentation.
// In this document we will explain all the possible options that can be used within the YAML format.
// Generally this is the preferred format for writing tests in as they don't require any programming skills
// in order to understand and change. If you find a bug in this format please open an issue.

// FTWTest is the base type used when unmarshaling YAML tests files
type FTWTest struct {
	// description: |
	//   Meta describes the metadata information of this yaml test file
	Meta FTWTestMeta `yaml:"meta"`

	// description: |
	//   FileName is the name of the file where these tests are.
	// examples:
	//   - name: FileName
	//     value: "\"test-1234.yaml\""
	FileName string

	// description: |
	//   Tests is a list of FTW tests
	// examples:
	//   - value: exampleTests
	Tests []Test `yaml:"tests"`
}

// FTWTestMeta describes the metadata information of this yaml test file
type FTWTestMeta struct {
	// description: |
	//   Author is the list of authors that added content to this file
	// examples:
	//   - name: Author
	//     value: "\"Felipe Zipitria\""
	Author string `yaml:"author,omitempty"`

	// description: |
	//   Enabled indicates if the tests are enabled to be run by the engine or not.
	// examples:
	//   - name: Enabled
	//     value: false
	//
	// Deprecated: ignored; use platform specific overrides instead
	Enabled *bool `yaml:"enabled,omitempty"`

	// description: |
	//   Name is the name of the tests contained in this file.
	// examples:
	//   - name: Name
	//     value: "\"test01\""
	Name string `yaml:"name,omitempty"`

	// description: |
	//   Description is a textual description of the tests contained in this file.
	// examples:
	//   - name: Description
	//     value: "\"The tests here target SQL injection.\""
	Description string `yaml:"description,omitempty"`

	// description: |
	//   Version is the version of the YAML Schema.
	// examples:
	//   - name: Version
	//     value: "\"v1\""
	Version string `yaml:"version,omitempty"`
}

// Test is an individual test. One test can have multiple stages.
type Test struct {
	// description: |
	//   TestTitle is the title of this particular test. It is used for inclusion/exclusion of each run by the tool.
	// examples:
	//   - value: exampleTest.TestTitle
	//
	// Deprecated: use `rule_id` and `test_id`
	TestTitle string `yaml:"test_title,omitempty"`

	// description: |
	//   RuleId is the ID of the rule this test targets
	// examples:
	//   - name: RuleId
	//     value: 123456
	RuleId int `yaml:"rule_id"`

	// description: |
	//   TestId is the ID of the test, in relation to `rule_id`
	// examples:
	//   - name: TestId
	//     value: 4
	TestId int `yaml:"test_id"`

	// description: |
	//   TestDescription is the description for this particular test. Should be used to describe the internals of
	//   the specific things this test is targeting.
	// examples:
	//   - value: exampleTest.TestDescription
	TestDescription string `yaml:"desc,omitempty"`

	// description: |
	//   Stages is the list of all the stages to perform this test.
	// examples:
	//   - value: exampleStages
	Stages []Stage `yaml:"stages"`
}

// Stage is a list of stages
type Stage struct {
	// description: |
	//   StageData is an individual test stage.
	//
	// Deprecated: use the other fields of `Stage`
	SD StageData `yaml:"stage,omitempty"`
	// description: |
	//   Describes the purpose of this stage.
	// examples:
	//   - value: exampleStage.Description
	Description string `yaml:"description,omitempty"`

	// description: |
	//   Input is the data that is passed to the test
	// examples:
	//   - name: Input
	//     value: exampleInput
	Input Input `yaml:"input"`

	// description: |
	//   Output is the data that is returned from the test
	// examples:
	//   - name: Output
	//     value: exampleOutput
	Output Output `yaml:"output"`
}

// StageData is the data that is passed to the test, and the data that is returned from the test
//
// Deprecated: use the other fields of `stage`
type StageData struct {
	// description: |
	//   Input is the data that is passed to the test
	// examples:
	//   - name: Input
	//     value: exampleInput
	Input Input `yaml:"input"`

	// description: |
	//   Output is the data that is returned from the test
	// examples:
	//   - name: Output
	//     value: exampleOutput
	Output Output `yaml:"output"`
}

// Input represents the input request in a stage
// The fields `Version`, `Method` and `URI` we want to explicitly now when they are set to ""
type Input struct {
	// description: |
	//   DestAddr is the IP of the destination host that the test will send the message to.
	// examples:
	//   - name: DestAddr
	//     value: "\"127.0.0.1\""
	DestAddr *string `yaml:"dest_addr,omitempty" koanf:"dest_addr,omitempty"`

	// description: |
	//   Port allows you to declare which port on the destination host the test should connect to.
	// examples:
	//   - name: Port
	//     value: 80
	Port *int `yaml:"port,omitempty" koanf:"port,omitempty"`

	// description: |
	//   Protocol allows you to declare which protocol the test should use when sending the request.
	// examples:
	//   - name: Protocol
	//     value: "\"http\""
	Protocol *string `yaml:"protocol,omitempty" koanf:"protocol,omitempty"`

	// description: |
	//   URI allows you to declare the URI the test should use as part of the request line.
	// examples:
	//   - name: URI
	//     value: "\"/get?hello=world\""
	URI *string `yaml:"uri,omitempty" koanf:"uri,omitempty"`

	// description: |
	//   Version allows you to declare the HTTP version the test should use as part of the request line.
	// examples:
	//   - name: Version
	//     value: "\"1.1\""
	Version *string `yaml:"version,omitempty" koanf:"version,omitempty"`

	// description: |
	//   Method allows you to declare the HTTP method the test should use as part of the request line.
	// examples:
	//   - name: Method
	//     value: "\"GET\""
	Method *string `yaml:"method,omitempty" koanf:"method,omitempty"`

	// description: |
	//   Method allows you to declare headers that the test should send.
	// examples:
	//   - name: Headers
	//     value: exampleHeaders
	Headers map[string]string `yaml:"headers,omitempty" koanf:"headers,omitempty"`

	// description: |
	//   Data allows you to declare the payload that the test should in the request body.
	// examples:
	//   - name: Data
	//     value: "\"Bibitti bopi\""
	Data *string `yaml:"data,omitempty" koanf:"data,omitempty"`

	// description: |
	//   SaveCookie allows you to automatically provide cookies if there are multiple stages and save cookie is set
	// examples:
	//   - name: SaveCookie
	//     value: 80
	SaveCookie *bool `yaml:"save_cookie,omitempty" koanf:"save_cookie,omitempty"`

	// description: |
	//   StopMagic is deprecated.
	// examples:
	//   - name: StopMagic
	//     value: false
	//
	// Deprecated: use AutocompleteHeaders instead
	StopMagic *bool `yaml:"stop_magic" koanf:"stop_magic,omitempty"`

	// description: |
	//   AutocompleteHeaders allows the test framework to automatically fill the request with Content-Type and Connection headers.
	//   Defaults to true.
	// examples:
	//   - name: StopMagic
	//     value: false
	AutocompleteHeaders *bool `yaml:"autocomplete_headers" koanf:"autocomplete_headers,omitempty"`

	// description: |
	//   EncodedRequest will take a base64 encoded string that will be decoded and sent through as the request.
	//   It will override all other settings
	// examples:
	//   - name: EncodedRequest
	//     value: "\"a\""
	EncodedRequest string `yaml:"encoded_request,omitempty" koanf:"encoded_request,omitempty"`

	// description: |
	//   RAWRequest is deprecated.
	// examples:
	//   - name: RAWRequest
	//     value: "\"TXkgRGF0YQo=\""
	//
	// Deprecated: use `encoded_request`
	RAWRequest string `yaml:"raw_request,omitempty" koanf:"raw_request,omitempty"`
}

// Output is the response expected from the test
type Output struct {
	// description: |
	//   Status describes the HTTP status code expected in the response.
	// examples:
	//   - name: Status
	//     value: 200
	Status int `yaml:"status,omitempty"`

	// description: |
	//   ResponseContains describes the text that should be contained in the HTTP response.
	// examples:
	//   - name: ResponseContains
	//     value: "\"Hello, World\""
	ResponseContains string `yaml:"response_contains,omitempty"`

	// description: |
	//   LogContains describes the text that should be contained in the WAF logs.
	// examples:
	//   - name: LogContains
	//     value: "\"id 920100\""
	//
	// Deprecated: use Log instead
	LogContains string `yaml:"log_contains,omitempty"`

	// description: |
	//   NoLogContains describes the text that should not be contained in the WAF logs.
	// examples:
	//   - name: NoLogContains
	//     value: "\"id 920100\""
	//
	// Deprecated: use Log instead
	NoLogContains string `yaml:"no_log_contains,omitempty"`

	// description: |
	//   Log is used to configure expectations about the log contents.
	// examples:
	//   - value: exampleLog
	Log Log `yaml:"log,omitempty"`

	// description: |
	//   When `ExpectError` is true, we don't expect an answer from the WAF, just an error.
	// examples:
	//   - name: ExpectError
	//     value: false
	ExpectError *bool `yaml:"expect_error,omitempty"`
}

// Log is used to configure expectations about the log contents.
type Log struct {
	// description: |
	//   Expect the given ID to be contained in the log output.
	// examples:
	//   - value: exampleLog.ExpectId
	ExpectId int `yaml:"expect_id,omitempty"`

	// description: |
	//   Expect the given ID _not_ to be contained in the log output.
	// examples:
	//   - value: exampleLog.NoExpectId
	NoExpectId int `yaml:"no_expect_id,omitempty"`

	// description: |
	//   Expect the regular expression to match log content for the current test.
	// examples:
	//   - value: exampleLog.MatchRegex
	MatchRegex string `yaml:"match_regex,omitempty"`

	// description: |
	//   Expect the regular expression to _not_ match log content for the current test.
	// examples:
	//   - value: exampleLog.NoMatchRegex
	NoMatchRegex string `yaml:"no_match_regex,omitempty"`
}

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
	//   - value: metaExample
	Meta FTWOverridesMeta `yaml:"meta"`

	// description: |
	//   List of test override specifications
	// examples:
	//   - value: testOverridesExample
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
	//   - value: annotationsExample
	Annotations map[string]string `yaml:"annotations"`
}

// TestOverride describes overrides for a single test
type TestOverride struct {
	// description: |
	//   ID of the rule this test targets.
	// examples:
	//   - value: testOverridesExample[0].RuleId
	RuleId int `yaml:"rule_id"`

	// description: |
	//   IDs of the tests for rule_id that overrides should be applied to.
	//   If this field is not set, the overrides will be applied to all tests of rule_id.
	// examples:
	//   - value: testOverridesExample[0].TestIds
	TestIds []int `yaml:"test_ids,flow,omitempty"`

	// description: |
	//   Describes why this override is necessary.
	// examples:
	//   - value: reasonExample
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
	Output Output `yaml:"output"`
}
