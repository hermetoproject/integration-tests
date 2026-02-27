package main

import (
	"fmt"
	"io"
	"os"

	"github.com/release-engineering/retrodep/v2/retrodep"

	"github.com/hermetoproject/integration-tests/terminaltor"
	where "github.com/hermetoproject/integration-tests/where-was-i-built"
)

var _ = retrodep.ErrorNoGo

func RunMain(out io.Writer) {
	msg, err := terminaltor.Terminalte()
	if err == nil {
		fmt.Fprintf(out, "Where am I running? %s\n", msg)
	}
	fmt.Fprintf(out, "Where was I built? %s\n", where.WhatOS())
}

func main() {
	RunMain(os.Stdout)
}
