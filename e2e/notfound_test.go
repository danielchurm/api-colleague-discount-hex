package e2e_test

import (
	"io"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Not Found", func() {
	Context("When requesting a route that does not exist", func() {
		It("It returns an API Problem", func() {
			client := http.Client{}
			req, _ := http.NewRequest(http.MethodGet, "http://api-go-template-ecs.app.internal:8080/never/going/to/exist", nil)
			req.Header.Add("Accept", "application/json")

			resp, err := client.Do(req)

			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))

			actBody, _ := io.ReadAll(resp.Body)
			expBody := `{
				"type": "https://github.com/JSainsburyPLC/smartshop-api-go-template/blob/develop/README_TEMPLATE.md#Error-Codes",
				"status": 404,
				"title" : "Not Found",
				"detail": "Unable to locate resource /never/going/to/exist",
				"code"  : 99000
			}`

			Expect(string(actBody)).To(MatchJSON(expBody))
		})
	})
})
