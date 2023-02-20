package main

import (
	"os"
)

func main() {
	println("get arg: ", os.Args[1])

	err := CreateDir("model")
	if err != nil {
		panic(err)
	}

	err = Parse(os.Args[1])
	if err != nil {
		panic(err)
	}
}
