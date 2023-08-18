package e2e_test

import (
	"testing"

	"github.com/caarlos0/env/v6"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type E2eConfig struct {
	ApiColleagueDiscountHost string `env:"API_COLLEAGUE_DISCOUNT_HOST" envDefault:"http://api-colleague-discount-ecs.app.internal:8080"`
}

var cfg E2eConfig

func TestConfiguration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Test Suite")
}

var _ = BeforeSuite(func() {
	Expect(env.Parse(&cfg)).To(Succeed())
})
