/*
Copyright © 2025 da99/diego <null@null>

*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("COP: %v\n", runtime.Compiler)
	fmt.Printf("ARC: %v\n", runtime.GOARCH)
	fmt.Printf("MAX: %v\n", runtime.GOMAXPROCS(0))
	fmt.Printf("VER: %v\n", runtime.Version())
}
