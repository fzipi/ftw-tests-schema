// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package types

import "github.com/coreruleset/ftw-tests-schema/internal/helpers"

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
		DestAddr:            helpers.StrPtr("192.168.0.1"),
		Port:                helpers.IntPtr(8080),
		Protocol:            helpers.StrPtr("http"),
		URI:                 helpers.StrPtr("/test"),
		Version:             helpers.StrPtr("HTTP/1.1"),
		Headers:             ExampleHeaders,
		Method:              helpers.StrPtr("REPORT"),
		Data:                nil,
		EncodedRequest:      "TXkgRGF0YQo=",
		SaveCookie:          helpers.BoolPtr(false),
		StopMagic:           helpers.BoolPtr(true),
		AutocompleteHeaders: helpers.BoolPtr(false),
	}
	ExampleOutput = Output{
		Status:           200,
		ResponseContains: "",
		LogContains:      "nothing",
		NoLogContains:    "",
		Log:              ExampleLog,
		ExpectError:      helpers.BoolPtr(true),
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
