package main_test

import (
	"bytes"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/cachito-testing/gomod-pandemonium"
)

func TestGomodPandemonium(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GomodPandemonium Suite")
}

var _ = Describe("GomodPandemonium", func() {
	It("works", func() {
		var buf bytes.Buffer
		main.RunMain(&buf)

		Expect(buf.String()).To(MatchRegexp(`Where am I running\? Not a terminal.\nWhere was I built\? .*(Windows!|Unix-like!|no idea.)\n`))
	})
})
