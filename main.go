package main

import (
	"fmt"
	"travel-app-api/api"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
}

func main() {
    // Disable Extra Gin Logs
    gin.SetMode(gin.ReleaseMode)

    // Create app instance
    app := api.NewApplication()
    
    // Set server port and start
    listenAddr := "localhost:8080"
    server := api.NewServer(listenAddr)
    server.Start(app)

    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":5000", nil)
    
}