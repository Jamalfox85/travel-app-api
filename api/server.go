package api

import (
	"fmt"
	"travel-app-api/api/handlers"
	"travel-app-api/data"

	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
	Users *data.UserRepository
	Trips *data.TripRepository
	ItineraryItems *data.ItineraryItemRepository
}

func NewServer(	listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start(app *Application) {
	router := gin.Default()

	// Users
	router.GET("/users/:userId", handlers.GetUser(app.Users))
	
	// Trips
	router.GET("/trips/:userId", handlers.GetTrips(app.Trips))

	// Itinerary Items
	router.GET("/itinerary/:tripId", handlers.GetItineraryItems(app.ItineraryItems))


	
	fmt.Println("Server Running on", s.listenAddr);
    router.Run(s.listenAddr)
}

