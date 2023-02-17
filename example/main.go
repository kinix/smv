package main

import (
	smv "github.com/kinix/struct-memory-visualization"
)

type exampleStruct struct {
	bool1 bool
	num1  int64
	bool2 bool
	bool4 bool
	num2  int32
	bool5 bool
}

func main() {
	smv.DrawMemory(exampleStruct{}, "image.png")
}
