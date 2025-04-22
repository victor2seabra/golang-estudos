package main

import (
	"fmt"
	"strings"
)

func hello(name string) {
	fmt.Printf("Hello, %s\n", name)
	strings.Join([]string{"eu", "victor"}, "-")
}

func main() {
	hello("Victor")
}
