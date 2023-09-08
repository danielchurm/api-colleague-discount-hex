package outbound

import "net/http"

//go:generate mockgen -destination=../mocks/outbound/http_client.go -source=http_client.go

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
