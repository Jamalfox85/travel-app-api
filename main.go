// package main

// import (
// 	"fmt"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func helloWorld(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hello World")
// }

// func main() {
//     // Disable Extra Gin Logs
//     // gin.SetMode(gin.ReleaseMode)

//     // Create app instance
//     // app := api.NewApplication()

//     // Set server port and start
//     // listenAddr := "localhost:8080"
//     // server := api.NewServer(listenAddr)
//     // server.Start(app)

//     http.HandleFunc("/", helloWorld)
//     http.ListenAndServe(":5000", nil)

// }

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)
func main() {
 port := os.Getenv("PORT")
helloHandler := func(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Hello, world!\n")
 }
http.HandleFunc("/", helloHandler)
 log.Println("Listing for" + port)
 log.Fatal(http.ListenAndServe(":"+port, nil))
}