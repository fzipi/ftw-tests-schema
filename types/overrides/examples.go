// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package overrides

import "github.com/coreruleset/ftw-tests-schema/types"

var (
	// imported
	AnnotationsExample = types.AnnotationsExample
	ReasonExample      = types.ReasonExample
	ExampleLog         = types.ExampleLog

	MetaExample = FTWOverridesMeta{
		Engine:      "libmodsecurity3",
		Platform:    "nginx",
		Annotations: types.AnnotationsExample,
	}
	TestOverridesExample = []TestOverride{
		{
			RuleId:        920100,
			TestIds:       []int{4, 6},
			Reason:        types.ReasonExample,
			ExpectFailure: func() *bool { b := true; return &b }(),
			Output:        types.ExampleOutput,
		},
	}
)
