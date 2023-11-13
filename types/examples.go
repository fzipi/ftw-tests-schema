// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package types

var (
	exampleTests = []Test{
		exampleTest,
	}
	exampleTest = Test{
		TestTitle:       "123456-1",
		TestDescription: "Unix RCE using `time`",
		Stages:          exampleStages,
	}
	exampleStage = Stage{
		Description: "Get cookie from server",
		Input:       exampleInput,
		Output:      exampleOutput,
	}
	exampleStages = []Stage{
		exampleStage,
	}
	exampleHeaders = map[string]string{
		"User-Agent": "CRS Tests",
		"Host":       "localhost",
		"Accept":     "*/*",
	}
	exampleInput = Input{
		DestAddr:            strPtr("192.168.0.1"),
		Port:                intPtr(8080),
		Protocol:            strPtr("http"),
		URI:                 strPtr("/test"),
		Version:             strPtr("HTTP/1.1"),
		Headers:             exampleHeaders,
		Method:              strPtr("REPORT"),
		Data:                nil,
		EncodedRequest:      "TXkgRGF0YQo=",
		SaveCookie:          boolPtr(false),
		StopMagic:           boolPtr(true),
		AutocompleteHeaders: boolPtr(false),
	}
	exampleOutput = Output{
		Status:           200,
		ResponseContains: "",
		LogContains:      "nothing",
		NoLogContains:    "",
		Log:              exampleLog,
		ExpectError:      boolPtr(true),
	}
	exampleLog = Log{
		ExpectId:     123456,
		NoExpectId:   123456,
		MatchRegex:   `id[:\s"]*123456`,
		NoMatchRegex: `id[:\s"]*123456`,
	}
	metaExample = FTWOverridesMeta{
		Engine:      "libmodsecurity3",
		Platform:    "nginx",
		Annotations: annotationsExample,
	}
	annotationsExample = map[string]string{
		"os":      "Debian Bullseye",
		"purpose": "L7ASR test suite",
	}
	reasonExample = "nginx returns 400 when `Content-Length` header is sent in a\n" +
		"`Transfer-Encoding: chunked` request."
	testOverridesExample = []TestOverride{
		{
			RuleId:        920100,
			TestIds:       []int{4, 6},
			Reason:        reasonExample,
			ExpectFailure: func() *bool { b := true; return &b }(),
			Output:        exampleOutput,
		},
	}
)
