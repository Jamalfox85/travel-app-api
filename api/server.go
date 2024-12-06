package api

import (
	"fmt"
	"log"
	"travel-app-api/api/handlers"
	"travel-app-api/data"

	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
	Users *data.UserRepository
	Trips *data.TripRepository
	ItineraryItems *data.ItineraryItemRepository
	Accommodations *data.AccommodationRepository
}

func NewServer(	listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start(app *Application) {
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Users
	router.GET("/users/:userId", handlers.GetUser(app.Users))
	
	// Trips
	router.GET("/trips/:userId", handlers.GetTrips(app.Trips))
	router.POST("/trips", handlers.CreateTrip(app.Trips))

	// Itinerary Items
	router.GET("/itinerary/:tripId", handlers.GetItineraryItems(app.ItineraryItems))
	router.POST("/itinerary", handlers.CreateItineraryItem(app.ItineraryItems))

	// Accommodations
	router.GET("/accommodations/:tripId", handlers.GetAccommodations(app.Accommodations))
	router.POST("/accommodations", handlers.CreateAccommodation(app.Accommodations))


	fmt.Println("Server Running on", s.listenAddr);
	if err := router.Run(":" + s.listenAddr); err != nil {
		log.Panicf("error: %s", err)
	}
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