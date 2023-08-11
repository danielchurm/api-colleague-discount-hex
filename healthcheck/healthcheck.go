package healthcheck

//go:generate mockgen -destination=../mocks/healthcheck/healthcheck.go -source=healthcheck.go
type HealthChecker interface {
	GetHealth() (Health, error)
}

type Health struct {
	Status string  `json:"status"`
	Errors []error `json:"errors"`
}

type HealthCheckService struct {
}

func NewHealthCheckService() *HealthCheckService {
	return &HealthCheckService{}
}

func (hcs HealthCheckService) GetHealth() (Health, error) {
	return Health{"OK", []error{}}, nil
}
