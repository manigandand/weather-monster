package config_test

import (
	"os"
	. "weather-monster/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Context("Test MustEnv", func() {
		BeforeEach(func() {
			os.Setenv("ENV", EnvDev)
		})
		AfterEach(func() {
			// flush all the env
			os.Clearenv()
		})

		It("Should read env values and assign to config variables", func() {
			GetAllEnv()
			Expect(Env).To(Equal(EnvDev))
		})
		It("Should update the env", func() {
			os.Setenv("ENV", EnvStaging)
			GetAllEnv()
			Expect(Env).To(Equal(EnvStaging))
		})
	})
})
