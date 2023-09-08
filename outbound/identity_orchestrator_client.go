package outbound

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/JSainsburyPLC/go-logrus-wrapper"
)

type idOrchRespBody struct {
	Email string `json:"email"`
}

type IdentityOrchestratorClient struct {
	host       string
	apikey     string
	httpClient HttpClient
}

func NewIdentityOrchestratorClient(host, apikey string, client HttpClient) IdentityOrchestratorClient {
	return IdentityOrchestratorClient{
		host:       host,
		apikey:     apikey,
		httpClient: client,
	}
}

func (c IdentityOrchestratorClient) GetEmail(ctx context.Context, userId int) (string, error) {
	url := fmt.Sprintf("http://%s/api/v1/users/%d", c.host, userId)
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("X-Request-ID", ctx.Value(log.ContextKeyRequestID).(string))
	req.Header.Add("X-User", ctx.Value(log.ContextKeyUserID).(string))
	req.Header.Add("X-API-KEY", c.apikey)

	resp, _ := c.httpClient.Do(req)

	var body idOrchRespBody
	_ = json.NewDecoder(resp.Body).Decode(&body)

	return body.Email, nil
}
