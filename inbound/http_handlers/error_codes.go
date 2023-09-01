package http_handlers

const (
	NotFoundError = 99000

	InternalServerError         = 39000
	InvalidUserIDInQueryParam   = 39001
	FailedToMarshalResponseBody = 39002

	OrchestratorFailedToGetUserDetails = 39100

	GetDiscountFailedWithBadRequest      = 39200
	GetDiscountFailedWithNotFound        = 39201
	GetDiscountFailedWithUnexpectedError = 39202
	GetDiscountFailedCardNumberTooShort  = 39203
)
