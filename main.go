package main

import (
	"fmt"

	"github.com/dasuken/sandbox/pkg"
)

func main() {
	s := pkg.New()
	fmt.Println("hello world" + s)
}
