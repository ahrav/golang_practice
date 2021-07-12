package main

import (
	"fmt"
	"runtime"
)

const ok = true

func main() {
	hey()
	fmt.Println("Hello Gopher!")
	bye()
	var hello = "Hello!"
	fmt.Println(hello, ok)
	fmt.Println(runtime.NumCPU())
}
