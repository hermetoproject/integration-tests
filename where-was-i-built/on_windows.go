//go:build windows

package where_was_i_built

import "github.com/Microsoft/go-winio"

func useTheArbitraryDependency() {
	winio.DialPipe("\\some\\path", nil)
}

func WhatOS() string {
	return "Windows!"
}
