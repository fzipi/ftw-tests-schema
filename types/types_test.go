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
    desc: "Test that the schema is working."
    stages:
      - stage:
          input:
            dest_addr: "127.0.0.1"
            port: 80
            headers:
              User-Agent: "FTW Schema Tests"
              Host: "localhost"
          output:
            no_log_contains: 'id "1234"'
  - test_title: 1234-2
    stages:
      - stage:
          input:
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
			TestDescription: "Test that the schema is working.",
			Stages: []Stage{
				{
					SD: StageData{
						Input: Input{
							DestAddr: strPtr("127.0.0.1"),
							Port:     intPtr(80),
							Headers: map[string]string{
								"User-Agent": "FTW Schema Tests",
								"Host":       "localhost",
							},
						},
						Output: Output{
							NoLogContains: "id \"1234\"",
						},
					},
				},
			},
		},
		{
			TestTitle: "1234-2",
			Stages: []Stage{
				{
					SD: StageData{
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
		assert.Len(test.Stages, len(expectedTest.Stages))

		for j, stage := range test.Stages {
			expectedStage := expectedTest.Stages[j]
			assert.Equal(expectedStage.SD.Input.DestAddr, stage.SD.Input.DestAddr)
			assert.Equal(expectedStage.SD.Input.Port, stage.SD.Input.Port)
			assert.Equal(expectedStage.SD.Input.Method, stage.SD.Input.Method)
			assert.Len(stage.SD.Input.Headers, len(expectedStage.SD.Input.Headers))

			for k, header := range stage.SD.Input.Headers {
				expectedHeader := expectedStage.SD.Input.Headers[k]
				assert.Equal(expectedHeader, header)
			}

			assert.Equal(expectedStage.SD.Output.NoLogContains, stage.SD.Output.NoLogContains)
			assert.Equal(expectedStage.SD.Output.Status, stage.SD.Output.Status)
		}
	}
}

// overrides

var overridesYaml = `---
version: "v0.0.0"
meta:
  engine: "libmodsecurity3"
  platform: "nginx"
  annotations:
    os: "Debian Bullseye"
    purpose: "L7ASR test suite"
test_overrides:
  - rule_id: 920100
    test_ids: [4, 6]
    reason: |-
      nginx returns 400 when ` + "`" + "Content-Length" + "`" + ` header is sent in a
      ` + "`" + `Transfer-Encoding: chunked` + "`" + ` request.
    expect_failure: true
    output:
      status: 200
      log_contains: "nothing"
      log:
        match_regex: 'id[:\s"]*123456'
        no_match_regex: 'id[:\s"]*123456'
      expect_error: true
`

var ftwOverrides = &FTWOverrides{
	Version:       "v0.0.0",
	Meta:          metaExample,
	TestOverrides: testOverridesExample,
}

func TestUnmarshalFTWOverrides(t *testing.T) {
	var overrides FTWOverrides

	err := yaml.Unmarshal([]byte(overridesYaml), &overrides)

	if err != nil {
		t.Errorf("Unmarshal: %v", err)
	}

	assert.Equal(t, ftwOverrides.Version, overrides.Version)
}

func TestUnmarshalMeta(t *testing.T) {
	var overrides FTWOverrides
	assert := assert.New(t)

	err := yaml.Unmarshal([]byte(overridesYaml), &overrides)
	if err != nil {
		t.Errorf("Unmarshal: %v", err)
	}

	meta := overrides.Meta
	expectedMeta := ftwOverrides.Meta
	assert.Equal(expectedMeta.Engine, meta.Engine)
	assert.Equal(expectedMeta.Platform, meta.Platform)
	assert.NotNil(meta.Annotations)
	assert.Len(meta.Annotations, len(expectedMeta.Annotations))

	for key, value := range meta.Annotations {
		expectedValue := expectedMeta.Annotations[key]
		assert.Equal(expectedValue, value)
	}
}

func TestUnmarmarshalTestOverrides(t *testing.T) {
	var overrides FTWOverrides
	asserter := assert.New(t)

	err := yaml.Unmarshal([]byte(overridesYaml), &overrides)
	if err != nil {
		t.Errorf("Unmarshal: %v", err)
	}

	testOverrides := overrides.TestOverrides
	expectedTestOverrides := ftwOverrides.TestOverrides
	asserter.NotNil(testOverrides)
	asserter.Len(testOverrides, len(expectedTestOverrides))

	for index, testOverride := range testOverrides {
		expectedTestOverride := expectedTestOverrides[index]
		asserter.Equal(expectedTestOverride.RuleId, testOverride.RuleId)
		asserter.ElementsMatch(expectedTestOverride.TestIds, testOverride.TestIds)
		asserter.Equal(expectedTestOverride.Reason, testOverride.Reason)
		asserter.Equal(expectedTestOverride.ExpectFailure, testOverride.ExpectFailure)
		if !assert.ObjectsAreEqual(expectedTestOverride.Output, testOverride.Output) {
			asserter.Failf("Output:", "%v != %v", testOverride.Output, expectedTestOverride.Output)
		}
	}
}
