package data

import (
	"database/sql"
	"fmt"
	"log"

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
			Start_date:	row.StartDate.Format("2006-01-02"),
			End_date:	row.EndDate.Format("2006-01-02"),
		}
		trips = append(trips, trip)
	}

	return trips, nil
}