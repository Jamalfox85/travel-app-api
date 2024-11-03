package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"travel-app-api/data"

	"github.com/joho/godotenv"
)

func NewApplication() *Application {
	// Load env file
	godotenv.Load()

	// Create DB Instance and table repositories
	db := newDB()
	users := data.NewUserRepository(db)
	trips := data.NewTripRepository(db)
	itineraryItems := data.NewItineraryItemRepository(db)

	return &Application{
		Users: users,
		Trips: trips,
		ItineraryItems: itineraryItems,
	}
}

func newDB() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("AWS_DBS_CONNECTION_STRING"))
	if err != nil {
		log.Fatal("Error initializing database");
	}
	fmt.Println("Connected to DBS!")
	return db;
}