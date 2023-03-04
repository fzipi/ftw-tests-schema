//go:generate docgen types.go types_doc.go FTWTest
package main

// Input represents the input request in a stage
// The fields `Version`, `Method` and `URI` we want to explicitly now when they are set to ""
type Input struct {
	// description: |
	// DestAddr is the IP of the destination host that the test will send the message to.
	// examples:
	//   - name: DestAddr
	//     value: "127.0.0.1"
	DestAddr *string `yaml:"dest_addr,omitempty" koanf:"dest_addr,omitempty"`
	// description: |
	// Port allows you to declare which port on the destination host the tests should connect to.
	// examples:
	//   - name: Port
	//     value: 80
	Port *int `yaml:"port,omitempty" koanf:"port,omitempty"`
	// description: |
	// Protocol allows you to declare which port on the destination host the tests should connect to.
	// examples:
	//   - name: Protocol
	//     value: 80
	Protocol *string `yaml:"protocol,omitempty" koanf:"protocol,omitempty"`
	// description: |
	// URI allows you to declare which port on the destination host the tests should connect to.
	// examples:
	//   - name: URI
	//     value: "/get?hello=world"
	URI *string `yaml:"uri,omitempty" koanf:"uri,omitempty"`
	// description: |
	// Version it the HTTP version used.
	// examples:
	//   - name: Version
	//     value: "1.1"
	Version *string `yaml:"version,omitempty" koanf:"version,omitempty"`
	// description: |
	// Headers is a map of headers to be sent.
	// examples:
	//   - name: Headers
	//     value: list
	//Headers ftwhttp.Header `yaml:"headers,omitempty" koanf:"headers,omitempty"`
	// description: |
	// Method allows you to declare which port on the destination host the tests should connect to.
	// examples:
	//   - name: Method
	//     value: "GET"
	Method *string `yaml:"method,omitempty" koanf:"method,omitempty"`
	// description: |
	// Data allows you to declare which port on the destination host the tests should connect to.
	// examples:
	//   - name: Data
	//     value: "Bibitti bopi"
	Data *string `yaml:"data,omitempty" koanf:"data,omitempty"`
	// description: |
	// SaveCookie allows you to declare which port on the destination host the tests should connect to.
	// examples:
	//   - name: SaveCookie
	//     value: 80
	SaveCookie bool `yaml:"save_cookie,omitempty" koanf:"save_cookie,omitempty"`
	// description: |
	// StopMagic allows you to declare which port on the destination host the tests should connect to.
	// examples:
	//   - name: StopMagic
	//     value: false
	StopMagic bool `yaml:"stop_magic" koanf:"stop_magic,omitempty"`
	// description: |
	// EncodedRequest allows you to declare which port on the destination host the tests should connect to.
	// examples:
	//   - name: EncodedRequest
	//     value: "a"
	EncodedRequest string `yaml:"encoded_request,omitempty" koanf:"encoded_request,omitempty"`
	// description: |
	// RAWRequest is deprecated.
	// examples:
	//   - name: RAWRequest
	//     value: asd
	RAWRequest string `yaml:"raw_request,omitempty" koanf:"raw_request,omitempty"`
}

// Output is the response expected from the test
type Output struct {
	// description: |
	// Status describes the HTTP status error code expected as response.
	// examples:
	//   - name: Status
	//     value: [200]
	Status []int `yaml:"status,flow,omitempty"`
	// description: |
	// ResponseContains describes the text that should be contained in the HTTP response.
	// examples:
	//   - name: ResponseContains
	//     value: "Hello, World"
	ResponseContains string `yaml:"response_contains,omitempty"`
	// description: |
	// LogContains describes the text that should be contained in the WAF logs.
	// examples:
	//   - name: LogContains
	//     value: `id "920100"`
	LogContains string `yaml:"log_contains,omitempty"`
	// description: |
	// NoLogContains describes the text that should be contained in the WAF logs.
	// examples:
	//   - name: NoLogContains
	//     value: `id "920100"
	NoLogContains string `yaml:"no_log_contains,omitempty"`
	// description: |
	// ExpectError describes the text that should be contained in the WAF logs.
	// examples:
	//   - name: ExpectError
	//     value: `id "920100"
	ExpectError bool `yaml:"expect_error,omitempty"`
}

// Stage is an individual test stage
type Stage struct {
	// description: |
	// Input is the data that is passed to the test
	// examples:
	//   - name: Input
	//     value: test
	Input Input `yaml:"input"`
	// description: |
	// Output is the data that is returned from the test
	// examples:
	//   - name: Output
	//     value: test
	Output Output `yaml:"output"`
}

// Stages is the list of all the stages to perform this test.
type Stages []Stage

// Test is an individual test
type Test struct {
	// description: |
	// TestTitle the title of this particular test. It is used for inclusion/exclusion of each run by the tool.
	// examples:
	//   - name: TestTitle
	//     value: 920100-1
	TestTitle string `yaml:"test_title"`
	// description: |
	// TestDescription is the description for this particular test. Should be used to describe the internals of
	// the specific things this test is targeting.
	// examples:
	//   - name: TestDescription
	//     value: This test targets something
	TestDescription string `yaml:"desc,omitempty"`
	// description: |
	Stages Stages `yaml:"stages"`
}

// Meta describes the metadata information of this yaml test file
type Meta struct {
	// description: |
	// Author is the list of authors that added content to this file
	// examples:
	//   - name: Author
	//     value: Felipe Zipitria
	Author string `yaml:"author,omitempty"`
	// description: |
	// Enabled indicates if the tests are enabled to be run by the engine or not.
	// examples:
	//   - name: Enabled
	//     value: false
	Enabled bool `yaml:"enabled,omitempty"`
	// description: |
	// Name is the name of the tests contained in this file.
	// examples:
	//   - name: Name
	//     value: test01
	Name string `yaml:"name,omitempty"`
	// description: |
	// Description is a textual description of the tests contained in this file.
	// examples:
	//   - name: Description
	//     value: The tests here target SQL injection.
	Description string `yaml:"description,omitempty"`
	// description: |
	// Version is the version of the YAML Schema.
	// examples:
	//   - name: Version
	//     value: v1
	Version string `yaml:"version,omitempty"`
}

// FTWTest is the base type used when unmarshaling YAML tests files
type FTWTest struct {
	// description: |
	// Meta describes the metadata information of this yaml test file
	Meta Meta `yaml:"meta"`

	// description: |
	// FileName is the name of the file where these tests are.
	// examples:
	//   - name: FileName
	//     value: test-1234.yaml
	FileName string `yaml:"filename"`
	// description: |
	// Tests is a list of FTW tests
	// examples:
	//   - name: Tests
	//     value: the tests
	Tests []Test `yaml:"tests"`
}
