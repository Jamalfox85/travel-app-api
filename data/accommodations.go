package data

import (
	"database/sql"
	"fmt"

	"travel-app-api/data/queries"

	"github.com/gin-gonic/gin"
)

type Accommodation struct {
	Id			int
	TripId		int
	Title		string
	Address		string
	StartDate	string
	EndDate		string
}

type AccommodationRepository struct {
	queries *queries.Queries
}

func NewAccommodationRepository(db *sql.DB) *AccommodationRepository {
	queries := queries.New(db)

	return &AccommodationRepository{
		queries: 	queries,
	}
}

func (r *AccommodationRepository) FindAccommodations(ctx *gin.Context, tripId int) ([]Accommodation, error) {
	formattedTripId := int32(tripId)
	rows, err := r.queries.GetAccommodations(ctx, formattedTripId)
	if err != nil {
		return nil, fmt.Errorf("error fetching accommodations for user", err)
	}

	var accommodations []Accommodation
	for _, row := range rows {
		item := Accommodation{
			Id:			int(row.ID),
			TripId:		int(row.Tripid),
			Title:		row.Title,
			Address:	row.Address.String,
			StartDate:	row.StartDate.Time.Format("2006-01-02"),
			EndDate:	row.EndDate.Time.Format("2006-01-02"),

		}
		accommodations = append(accommodations, item)
	}

	return accommodations, nil
}