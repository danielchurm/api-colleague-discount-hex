package outbound

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/domain"
)

type colleagueDiscountRespBody struct {
	CardNumber  string `json:"cardNumber"`
	IssueNumber int    `json:"issueNumber"`
	Status      string `json:"status"` // Possible values: NEW, VERIFIED
}

type CheckoutsColleagueDiscountClient struct {
	scheme     string
	host       string
	httpClient HttpClient
}

func NewCheckoutsColleagueDiscountClient(scheme, host string, httpClient HttpClient) CheckoutsColleagueDiscountClient {
	return CheckoutsColleagueDiscountClient{
		scheme:     scheme,
		host:       host,
		httpClient: httpClient,
	}
}

func (c CheckoutsColleagueDiscountClient) GetDiscountCard(ctx context.Context, email string) (domain.Card, error) {
	url := fmt.Sprintf("%s://%s/discount-card?email=%s", c.scheme, c.host, email)

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response, _ := c.httpClient.Do(request)

	var body colleagueDiscountRespBody
	_ = json.NewDecoder(response.Body).Decode(&body)

	card := domain.Card{
		CardNumber:  body.CardNumber,
		IssueNumber: body.IssueNumber,
		Status:      body.Status,
	}
	return card, nil
}
