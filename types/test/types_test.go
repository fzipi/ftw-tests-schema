// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
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
		Enabled:     boolPtr(true),
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
						DestAddr: strPtr("127.0.0.1"),
						Port:     intPtr(80),
						Method:   strPtr("OPTIONS"),
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
	assert := assert.New(t)

	err := yaml.Unmarshal([]byte(testYaml), &ftw)
	assert.NoError(err)

	assert.Equal(ftwTest.FileName, ftw.FileName)
	assert.Equal(ftwTest.Meta.Author, ftw.Meta.Author)
	assert.Equal(ftwTest.Meta.Enabled, ftw.Meta.Enabled)
	assert.Equal(ftwTest.Meta.Name, ftw.Meta.Name)
	assert.Equal(ftwTest.Meta.Description, ftw.Meta.Description)
	assert.Len(ftwTest.Tests, len(ftw.Tests))

	for i, test := range ftw.Tests {
		expectedTest := ftwTest.Tests[i]
		assert.Equal(expectedTest.TestTitle, test.TestTitle)
		assert.Equal(expectedTest.RuleId, test.RuleId)
		assert.Equal(expectedTest.TestId, test.TestId)
		assert.Len(test.Stages, len(expectedTest.Stages))

		for j, stage := range test.Stages {
			expectedStage := expectedTest.Stages[j]
			assert.Equal(expectedStage.Input.DestAddr, stage.Input.DestAddr)
			assert.Equal(expectedStage.Input.Port, stage.Input.Port)
			assert.Equal(expectedStage.Input.Method, stage.Input.Method)
			assert.Len(stage.Input.Headers, len(expectedStage.Input.Headers))

			for k, header := range stage.Input.Headers {
				expectedHeader := expectedStage.Input.Headers[k]
				assert.Equal(expectedHeader, header)
			}

			assert.Equal(expectedStage.Output.NoLogContains, stage.Output.NoLogContains)
			assert.Equal(expectedStage.Output.Status, stage.Output.Status)
		}
	}
}
