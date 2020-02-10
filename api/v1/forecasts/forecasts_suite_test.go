package forecasts_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestForecasts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Forecasts Suite")
}
