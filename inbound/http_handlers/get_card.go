package http_handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	apiproblem "github.com/JSainsburyPLC/go-api-problem"
	"github.com/JSainsburyPLC/smartshop-api-colleague-discount/domain"
)

//go:generate mockgen -destination=../../mocks/inbound/http_handlers/get_card.go -source=get_card.go
type ColleagueDiscountCardRetriever interface {
	GetCardForUser(userId int) (domain.Card, error)
}

type getCardResponse struct {
	CardNumber  string `json:"card_number"`
	IssueNumber string `json:"issue_number"`
	Status      string `json:"status"`
}

type GetCardHandler struct {
	apiProblemFactory              apiproblem.Factory
	colleagueDiscountCardRetriever ColleagueDiscountCardRetriever
}

func NewGetCardHandler(apiProblemFactory apiproblem.Factory, colleagueDiscountCardRetriever ColleagueDiscountCardRetriever) GetCardHandler {
	return GetCardHandler{
		apiProblemFactory:              apiProblemFactory,
		colleagueDiscountCardRetriever: colleagueDiscountCardRetriever,
	}
}

func (g GetCardHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userIdParam := request.URL.Query().Get("user_id")
	userId, err := readIntGreaterThanZero(userIdParam)
	if err != nil {
		problem := g.apiProblemFactory.NewApiProblem(
			http.StatusBadRequest,
			InvalidUserIDInQueryParam,
			"user id must be a positive integer",
		)
		problem.LogAndWriteDetailed(request.Context(), writer)
		return
	}

	cardForUser, _ := g.colleagueDiscountCardRetriever.GetCardForUser(userId)

	resp := getCardResponse{
		CardNumber:  cardForUser.CardNumber,
		IssueNumber: strconv.Itoa(cardForUser.IssueNumber),
		Status:      cardForUser.Status,
	}

	err = json.NewEncoder(writer).Encode(resp)
	if err != nil {
		// TODO
		return
	}
}

func readIntGreaterThanZero(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	if i <= 0 {
		return 0, errors.New(s + " is less than zero")
	}

	return i, nil
}
