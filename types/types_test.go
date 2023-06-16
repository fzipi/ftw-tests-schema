package types

import (
	"github.com/goccy/go-yaml"
	"testing"
)

var testInput = `---
dest_addr: "127.0.0.1"
port: 80
headers:
  User-Agent: "FTW Schema Tests"
  Host: "localhost"
`

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
            status: [200, 204]
`

var inputTest = &Input{
	DestAddr: strPtr("127.0.0.1"),
	Port:     intPtr(80),
	Headers: map[string]string{
		"User-Agent": "FTW Schema Tests",
		"Host":       "localhost",
	},
}

var ftwTest = &FTWTest{
	FileName: "testYaml.yaml",
	Meta: Meta{
		Author:      "ftw-tests-schema",
		Enabled:     true,
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
							Status: []int{200, 204},
						},
					},
				},
			},
		},
	},
}

func TestUnmarshalFTWTest(t *testing.T) {
	var ftw FTWTest

	err := yaml.Unmarshal([]byte(testYaml), &ftw)

	if err != nil {
		t.Errorf("Unmarshal: %v", err)
	}

	if ftw.FileName != ftwTest.FileName {
		t.Errorf("FileName: %v != %v", ftw.FileName, ftwTest.FileName)
	}
	if ftw.Meta.Author != ftwTest.Meta.Author {
		t.Errorf("Author: %v != %v", ftw.Meta.Author, ftwTest.Meta.Author)
	}
	if ftw.Meta.Enabled != ftwTest.Meta.Enabled {
		t.Errorf("Enabled: %v != %v", ftw.Meta.Enabled, ftwTest.Meta.Enabled)
	}
	if ftw.Meta.Name != ftwTest.Meta.Name {
		t.Errorf("Name: %v != %v", ftw.Meta.Name, ftwTest.Meta.Name)
	}
	if ftw.Meta.Description != ftwTest.Meta.Description {
		t.Errorf("Description: %v != %v", ftw.Meta.Description, ftwTest.Meta.Description)
	}
	if len(ftw.Tests) != len(ftwTest.Tests) {
		t.Errorf("Tests: %v != %v", len(ftw.Tests), len(ftwTest.Tests))
	}
	for i, test := range ftw.Tests {
		if test.TestTitle != ftwTest.Tests[i].TestTitle {
			t.Errorf("TestTitle: %v != %v", test.TestTitle, ftwTest.Tests[i].TestTitle)
		}
		if len(test.Stages) != len(ftwTest.Tests[i].Stages) {
			t.Errorf("Stages: %v != %v", len(test.Stages), len(ftwTest.Tests[i].Stages))
		}
		for j, stage := range test.Stages {
			if *stage.SD.Input.DestAddr != *ftwTest.Tests[i].Stages[j].SD.Input.DestAddr {
				t.Errorf("DestAddr: %v != %v", *stage.SD.Input.DestAddr, *ftwTest.Tests[i].Stages[j].SD.Input.DestAddr)
			}
			if *stage.SD.Input.Port != *ftwTest.Tests[i].Stages[j].SD.Input.Port {
				t.Errorf("Port: %v != %v", *stage.SD.Input.Port, *ftwTest.Tests[i].Stages[j].SD.Input.Port)
			}
			if stage.SD.Input.Method != nil && *stage.SD.Input.Method != *ftwTest.Tests[i].Stages[j].SD.Input.Method {
				t.Errorf("Method: %v != %v", stage.SD.Input.Method, ftwTest.Tests[i].Stages[j].SD.Input.Method)
			}
			if len(stage.SD.Input.Headers) != len(ftwTest.Tests[i].Stages[j].SD.Input.Headers) {
				t.Errorf("Headers: %v != %v", len(stage.SD.Input.Headers), len(ftwTest.Tests[i].Stages[j].SD.Input.Headers))
			}
			for k, header := range stage.SD.Input.Headers {
				if header != ftwTest.Tests[i].Stages[j].SD.Input.Headers[k] {
					t.Errorf("Header: %v != %v", header, ftwTest.Tests[i].Stages[j].SD.Input.Headers[k])
				}
			}
			if stage.SD.Output.NoLogContains != ftwTest.Tests[i].Stages[j].SD.Output.NoLogContains {
				t.Errorf("NoLogContains: %v != %v", stage.SD.Output.NoLogContains, ftwTest.Tests[i].Stages[j].SD.Output.NoLogContains)
			}
			if len(stage.SD.Output.Status) != len(ftwTest.Tests[i].Stages[j].SD.Output.Status) {
				t.Errorf("Status: %v != %v", len(stage.SD.Output.Status), len(ftwTest.Tests[i].Stages[j].SD.Output.Status))
			}
		}
	}
}

func TestUnmarshalInput(t *testing.T) {
	var input Input

	err := yaml.Unmarshal([]byte(testInput), &input)
	if err != nil {
		t.Errorf("Unmarshal: %v", err)
	}

	if input.DestAddr != nil && *input.DestAddr != *inputTest.DestAddr {
		t.Errorf("DestAddr: %v != %v", *input.DestAddr, *inputTest.DestAddr)
	}
	if input.Port != nil && *input.Port != *inputTest.Port {
		t.Errorf("Port: %v != %v", *input.Port, *inputTest.Port)
	}
	if input.Method != nil && *input.Method != *inputTest.Method {
		t.Errorf("Method: %v != %v", *input.Method, *inputTest.Method)
	}
}
