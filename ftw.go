// Copyright 2023 Felipe Zipitria
// SPDX-License-Identifier: Apache-2.0

package main

import "os"

func main() {
	data, err := GetFTWTestDoc().Encode()
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}
