package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"travel-app-api/data/queries"

	"github.com/gin-gonic/gin"
)

type Trip struct {
	ID				int
	Title			string
	Location		string
	UserID			int
	Start_date		string
	End_date		string
	Place_id		string
	Photo_uri		string
	Latitude		string
	Longitude		string
}

type TripRepository struct {
	queries *queries.Queries
}

func NewTripRepository(db *sql.DB) *TripRepository {
	queries := queries.New(db)

	return &TripRepository{
		queries: 	queries,
	}
}

func (r *TripRepository) FindTrips(ctx *gin.Context, userId int) ([]Trip, error) {
	formattedUserId := NewNullInt32(int32(userId))

	rows, err := r.queries.GetTripsByUser(ctx, formattedUserId)
	if err != nil {
		log.Fatalf("sqlc query error: %v", err)  // Use log.Fatalf for immediate feedback
	}

	var trips []Trip
	for _, row := range rows {
		fmt.Println(row);
		trip := Trip{
			ID:			int(row.Tripid),
			Title:		row.Title.String,
			Location:	row.Location.String,
			UserID:		int(row.Userid.Int32),
			Start_date:	row.StartDate.Time.Format("2006-01-02"),
			End_date:	row.EndDate.Time.Format("2006-01-02"),
			Place_id: 	row.PlaceID.String,
			Photo_uri: 	row.PhotoUri.String,
			Latitude:	row.Latitude.String,
			Longitude:	row.Longitude.String,
		}
		trips = append(trips, trip)
	}

	return trips, nil
}

func (r * TripRepository) CreateTrip(ctx *gin.Context, newTrip Trip) (error) {
	formattedStartDate, _ := time.Parse("2006-01-02", newTrip.Start_date)
	formattedEndDate, _ := time.Parse("2006-01-02", newTrip.End_date)

	params := queries.CreateTripParams{
		Title: sql.NullString{String: newTrip.Title, Valid: newTrip.Title != ""},
		Location: sql.NullString{String: newTrip.Location, Valid: newTrip.Location != ""},
		Userid: sql.NullInt32{Int32: int32(newTrip.UserID), Valid: true},
		StartDate: sql.NullTime{Time: formattedStartDate, Valid: true},
		EndDate:  sql.NullTime{Time: formattedEndDate, Valid: true},
		PlaceID: sql.NullString{String: newTrip.Place_id, Valid: newTrip.Place_id != ""},
		PhotoUri: sql.NullString{String: newTrip.Photo_uri, Valid: newTrip.Photo_uri != ""},
		Latitude:	sql.NullString{String: newTrip.Latitude, Valid: newTrip.Latitude != ""},
		Longitude:	sql.NullString{String: newTrip.Longitude, Valid: newTrip.Longitude != ""},
	}

	err := r.queries.CreateTrip(ctx, params);
	if err != nil {
		return fmt.Errorf("error creating new trip")
	}
	return nil
}