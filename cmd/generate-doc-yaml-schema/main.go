// Copyright 2023 OWASP CRS
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	overrides "github.com/coreruleset/ftw-tests-schema/types/overrides"
	test "github.com/coreruleset/ftw-tests-schema/types/test"
)

func main() {
	data, err := test.GetFTWTestDoc().Encode()
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}

	data, err = overrides.GetFTWOverridesDoc().Encode()
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}
}
