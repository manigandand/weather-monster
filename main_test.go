package main_test

import (
	"flag"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

// Can't test the main since it has infinite loop
var _ = Describe("Main Test Suite", func() {
	BeforeEach(func() {
	})
	AfterEach(func() {
	}, 0.1)
})
