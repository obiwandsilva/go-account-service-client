package controllers

import (
	"go-account-service-client/resources/gateways/accountgateway"
	"log"

	"github.com/gin-gonic/gin"
)

// GetAccountInfo is a handler for the route
func GetAccountInfo(ctx *gin.Context) {
	accountID := ctx.Param("accountId")
	result, err := accountgateway.GetAccountInfo(accountID)

	if err != nil {
		log.Printf("ERROR DETECTED: %v", err)
		ctx.Status(500)
		return
	}

	ctx.Header("content-type", "application/json")
	ctx.JSON(200, result)
}
