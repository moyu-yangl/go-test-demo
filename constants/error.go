package constants

type ErrorType struct {
	Code    int
	Message string
}

const (
	SystemErrorCode    = 901
	SystemErrorMessage = "system error please try again"

	TicketNotFoundCode    = 902
	TicketNotFoundMessage = "ticket not found"
)
