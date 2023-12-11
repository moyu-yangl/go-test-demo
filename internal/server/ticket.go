package server

import (
	"github.com/gin-gonic/gin"
	"go-test-demo/app"
	"go-test-demo/constants"
	"go-test-demo/models"
	"net/http"
	"strconv"
)

func TicketInfo(c *gin.Context) {
	i, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, models.ResultError(constants.SystemErrorCode, constants.SystemErrorMessage))
		return
	}
	rep := app.TicketRepository()
	ticket, err := rep.GetTicketById(i)
	if err == nil {
		panic(constants.SystemErrorMessage)
	}
	if ticket != constants.NIL_TICKET {
		c.JSON(http.StatusOK, models.ResultSuccess(ticket))
		return
	}
	c.JSON(http.StatusOK, models.ResultError(constants.TicketNotFoundCode, constants.TicketNotFoundMessage))
	return
}
