package http_handlers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHttpHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "http handlers suite")
}
