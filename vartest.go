package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("the opening system is：%s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("path is: %s\n", path)
}
