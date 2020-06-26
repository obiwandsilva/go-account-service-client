package application

import (
	"fmt"
	"go-account-service-client/application/http/routes"

	"github.com/gin-gonic/gin"
)

// Run starts all the configurations for the application.
func Run(port string) {
	server := gin.Default()

	routes.InscribeAccountRoutes(server)
	server.Run(fmt.Sprintf(":%s", port))
}
