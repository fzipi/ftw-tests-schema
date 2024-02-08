// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"

	"github.com/coreruleset/ftw-tests-schema/internal/helpers"
)

var testYaml = `---
filename: "testYaml.yaml"
meta:
  author: "ftw-tests-schema"
  enabled: true
  name: "testYaml"
  description: "Simple YAML to test that the schema is working."
tests:
  - test_title: 1234-1
    rule_id: 1234
    test_id: 1
    desc: "Test that the schema is working."
    stages:
      - input:
          dest_addr: "192.168.0.1"
          port: 8080
          method: "REPORT"
          headers:
            User-Agent: "CRS Tests"
            Host: "localhost"
            Accept: "*/*"
          encoded_request: "TXkgRGF0YQo="
          uri: "/test"
          protocol: "http"
          autocomplete_headers: false
          stop_magic: true
          save_cookie: false
        output:
          status: 200
          response_contains: ""
          log_contains: "nothing"
          no_log_contains: ""
  - test_title: 1234-2
    stages:
      - input:
          dest_addr: "127.0.0.1"
          port: 80
          method: "OPTIONS"
          headers:
            User-Agent: "FTW Schema Tests"
            Host: "localhost"
        output:
          status: 200
`

var ftwTest = &FTWTest{
	FileName: "testYaml.yaml",
	Meta: FTWTestMeta{
		Author:      "ftw-tests-schema",
		Enabled:     helpers.BoolPtr(true),
		Name:        "testYaml",
		Description: "Simple YAML to test that the schema is working.",
	},
	Tests: []Test{
		{
			TestTitle:       "1234-1",
			RuleId:          1234,
			TestId:          1,
			TestDescription: "Test that the schema is working.",
			Stages: []Stage{
				{
					Description: ExampleStage.Description,
					Input:       ExampleInput,
					Output:      ExampleOutput,
				},
			},
		},
		{
			TestTitle: "1234-2",
			Stages: []Stage{
				{
					Input: Input{
						DestAddr: helpers.StrPtr("127.0.0.1"),
						Port:     helpers.IntPtr(80),
						Method:   helpers.StrPtr("OPTIONS"),
						Headers: map[string]string{
							"User-Agent": "FTW Schema Tests",
							"Host":       "localhost",
						},
					},
					Output: Output{
						Status: 200,
					},
				},
			},
		},
	},
}

func TestUnmarshalFTWTest(t *testing.T) {
	var ftw FTWTest
	assertions := assert.New(t)

	err := yaml.Unmarshal([]byte(testYaml), &ftw)
	assertions.NoError(err)

	assertions.Equal(ftwTest.FileName, ftw.FileName)
	assertions.Equal(ftwTest.Meta.Author, ftw.Meta.Author)
	assertions.Equal(ftwTest.Meta.Enabled, ftw.Meta.Enabled)
	assertions.Equal(ftwTest.Meta.Name, ftw.Meta.Name)
	assertions.Equal(ftwTest.Meta.Description, ftw.Meta.Description)
	assertions.Len(ftwTest.Tests, len(ftw.Tests))

	for i, test := range ftw.Tests {
		expectedTest := ftwTest.Tests[i]
		assertions.Equal(expectedTest.TestTitle, test.TestTitle)
		assertions.Equal(expectedTest.RuleId, test.RuleId)
		assertions.Equal(expectedTest.TestId, test.TestId)
		assertions.Len(test.Stages, len(expectedTest.Stages))

		for j, stage := range test.Stages {
			expectedStage := expectedTest.Stages[j]
			assertions.Equal(expectedStage.Input.DestAddr, stage.Input.DestAddr)
			assertions.Equal(expectedStage.Input.Port, stage.Input.Port)
			assertions.Equal(expectedStage.Input.Method, stage.Input.Method)
			assertions.Len(stage.Input.Headers, len(expectedStage.Input.Headers))

			for k, header := range stage.Input.Headers {
				expectedHeader := expectedStage.Input.Headers[k]
				assertions.Equal(expectedHeader, header)
			}

			assertions.Equal(expectedStage.Output.NoLogContains, stage.Output.NoLogContains)
			assertions.Equal(expectedStage.Output.Status, stage.Output.Status)
		}
	}
}
