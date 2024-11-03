package api

import "travel-app-api/data"

type Application struct {
	Users				*data.UserRepository
	Trips				*data.TripRepository
	ItineraryItems		*data.ItineraryItemRepository
}