package main

import (
	"k8s.io/kubectl/pkg/util/templates"

	"github.com/kcp-dev/kcp-tests/pkg/test/ginkgo"
	_ "github.com/kcp-dev/kcp-tests/test/extended"
)

// staticSuites are all known test suites this binary should run
var staticSuites = []*ginkgo.TestSuite{
	{
		Name: "all",
		Description: templates.LongDesc(`
		Run all tests.
		`),
		Matches: func(name string) bool { return true },
	},
}
