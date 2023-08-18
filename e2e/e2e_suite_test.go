package e2e_test

import (
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/churmd/smockerclient"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type E2eConfig struct {
	ApiColleagueDiscountHost       string `env:"API_COLLEAGUE_DISCOUNT_HOST" envDefault:"http://api-colleague-discount-ecs.app.internal:8080"`
	MockSainsColleagueDiscountHost string `env:"COLLEAGUE_DISCOUNT_MOCK_ADMIN_HOST" envDefault:"http://sainsburys-colleague-discount-mock-server.app.internal:2080"`
	MockIdentityOrchestratorHost   string `env:"SMARTSHOP_ORCHESTRATOR_MOCK_HOST" envDefault:"http://smartshop-api-identity-orchestrator-mock-server.app.internal:2081"`
}

var (
	cfg                        E2eConfig
	mockSainsColleagueDiscount smockerclient.Instance
	mockIdentityOrchestrator   smockerclient.Instance
)

func TestConfiguration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Test Suite")
}

var _ = BeforeSuite(func() {
	Expect(env.Parse(&cfg)).To(Succeed())

	mockSainsColleagueDiscount = smockerclient.Instance{
		Url: cfg.MockSainsColleagueDiscountHost,
	}

	mockIdentityOrchestrator = smockerclient.Instance{
		Url: cfg.MockIdentityOrchestratorHost,
	}
})
