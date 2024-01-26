// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package types

import test "github.com/coreruleset/ftw-tests-schema/types/test"

var (
	// imported
	AnnotationsExample = test.AnnotationsExample
	ReasonExample      = test.ReasonExample
	ExampleLog         = test.ExampleLog

	MetaExample = FTWOverridesMeta{
		Engine:      "libmodsecurity3",
		Platform:    "nginx",
		Annotations: test.AnnotationsExample,
	}
	TestOverridesExample = []TestOverride{
		{
			RuleId:        920100,
			TestIds:       []int{4, 6},
			Reason:        test.ReasonExample,
			ExpectFailure: func() *bool { b := true; return &b }(),
			Output:        test.ExampleOutput,
		},
	}
)
