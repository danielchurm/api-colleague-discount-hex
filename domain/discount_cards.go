package domain

type Card struct {
	CardNumber  string `json:"cardNumber"`
	IssueNumber int    `json:"issueNumber"`
	Status      string `json:"status"` // Possible values: NEW, VERIFIED
}
