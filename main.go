package main

import (
	"os"
	"travel-app-api/api"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    // router := gin.Default()
	// router.Use(CORSMiddleware())
    // Disable Extra Gin Logs
    gin.SetMode(gin.ReleaseMode)

    // Create app instance
    app := api.NewApplication()

    // Set server port and start
    listenAddr := "8080"
    server := api.NewServer(listenAddr)
    server.Start(app)


    port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// if err := router.Run(":" + port); err != nil {
	// 	log.Panicf("error: %s", err)
	// }
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}