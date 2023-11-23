package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sidhant.borade/go-gin/banking/dto"
	"sidhant.borade/go-gin/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(ctx *gin.Context) {

	var request dto.NewAccountRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error-Message": err.Error()})
	}

	request.Customerid = ctx.Param("customer_id")
	account, appError := h.service.NewAccount(request)
	if appError != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error-Message": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusCreated, account)
	}
}
