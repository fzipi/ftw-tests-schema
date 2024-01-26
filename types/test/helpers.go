// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package types

func intPtr(i int) *int {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}

func strPtr(s string) *string {
	return &s
}
