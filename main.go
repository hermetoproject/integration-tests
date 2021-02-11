package main

import (
	"fmt"
	"github.com/cachito-testing/some-module"
)


func main() {
	fmt.Println("Hello, local dependencies.")
	some_module.Hi()
}
