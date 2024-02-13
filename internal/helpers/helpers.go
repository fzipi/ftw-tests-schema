// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package helpers

func IntPtr(i int) *int {
	return &i
}

func BoolPtr(b bool) *bool {
	return &b
}

func StrPtr(s string) *string {
	return &s
}
