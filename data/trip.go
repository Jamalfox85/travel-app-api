package data

import (
	"database/sql"
	"fmt"

	"travel-app-api/data/queries"

	"github.com/gin-gonic/gin"
)

type Trip struct {
	ID			int
	Title		string
	Location	string
	UserID		int
	StartDate	string
	EndDate		string
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
		return nil, fmt.Errorf("error fetching trips for user", err)
	}

	var trips []Trip
	for _, row := range rows {
		trip := Trip{
			ID:			int(row.Tripid),
			Title:		row.Title.String,
			Location:	row.Location.String,
			UserID:		int(row.Userid.Int32),
			StartDate:	row.StartDate.Time.Format("YYYY-MM-DD"),
			EndDate:	row.EndDate.Time.Format("YYYY-MM-DD"),
		}
		trips = append(trips, trip)
	}

	return trips, nil
}