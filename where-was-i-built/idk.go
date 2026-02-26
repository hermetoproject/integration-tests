//go:build !unix && !windows

package where_was_i_built

import "github.com/cachito-testing/gomod-pandemonium/weird"

func useTheArbitraryDependency() {
	weird.Weird()
}

func WhatOS() string {
	return "Not gonna lie, I have no idea."
}
