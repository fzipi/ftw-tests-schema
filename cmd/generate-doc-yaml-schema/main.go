// Copyright 2023 Felipe Zipitria
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/coreruleset/ftw-tests-schema/types"
)

func main() {
	data, err := types.GetFTWTestDoc().Encode()
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}
}
