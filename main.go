package main

import (
	"fmt"
	"github.com/cachito-testing/some-module"
	"github.com/cachito-testing/some-module/some-package"
)


func main() {
	fmt.Println("Hello, local dependencies.")
	some_module.Hi()
	some_package.Hi()
}
