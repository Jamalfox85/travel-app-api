package main

import (
	"os"
	"travel-app-api/api"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

    // Disable Extra Gin Logs
    gin.SetMode(gin.ReleaseMode)

    // Create app instance
    app := api.NewApplication()

    // Set server port and start
    port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
    server := api.NewServer(port)
    server.Start(app)

}

