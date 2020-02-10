package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWeatherMonster(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WeatherMonster Suite")
}
