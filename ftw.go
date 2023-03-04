package main

import "os"

func main() {
	data, err := GetFTWTestDoc().Encode()
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}
