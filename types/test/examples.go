// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package types

var (
	ExampleTests = []Test{
		ExampleTest,
	}
	ExampleTest = Test{
		TestTitle:       "123456-1",
		TestDescription: "Unix RCE using `time`",
		Stages:          ExampleStages,
	}
	ExampleStage = Stage{
		Description: "Get cookie from server",
		Input:       ExampleInput,
		Output:      ExampleOutput,
	}
	ExampleStages = []Stage{
		ExampleStage,
	}
	ExampleHeaders = map[string]string{
		"User-Agent": "CRS Tests",
		"Host":       "localhost",
		"Accept":     "*/*",
	}
	ExampleInput = Input{
		DestAddr:            strPtr("192.168.0.1"),
		Port:                intPtr(8080),
		Protocol:            strPtr("http"),
		URI:                 strPtr("/test"),
		Version:             strPtr("HTTP/1.1"),
		Headers:             ExampleHeaders,
		Method:              strPtr("REPORT"),
		Data:                nil,
		EncodedRequest:      "TXkgRGF0YQo=",
		SaveCookie:          boolPtr(false),
		StopMagic:           boolPtr(true),
		AutocompleteHeaders: boolPtr(false),
	}
	ExampleOutput = Output{
		Status:           200,
		ResponseContains: "",
		LogContains:      "nothing",
		NoLogContains:    "",
		Log:              ExampleLog,
		ExpectError:      boolPtr(true),
	}
	ExampleLog = Log{
		ExpectId:     123456,
		NoExpectId:   123456,
		MatchRegex:   `id[:\s"]*123456`,
		NoMatchRegex: `id[:\s"]*123456`,
	}
	AnnotationsExample = map[string]string{
		"os":      "Debian Bullseye",
		"purpose": "L7ASR test suite",
	}
	ReasonExample = "nginx returns 400 when `Content-Length` header is sent in a\n" +
		"`Transfer-Encoding: chunked` request."
)
