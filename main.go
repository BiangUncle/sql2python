package main

import (
	"os"
	"sql2python/parse"
)

func main() {
	println("get arg: ", os.Args[1])
	parse.Parse(os.Args[1])
}
