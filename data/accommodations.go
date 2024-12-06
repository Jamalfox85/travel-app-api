package data

import (
	"database/sql"
	"fmt"
	"time"

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
	Url			string
	Phone		string
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

func (r *AccommodationRepository) CreateAccommodation(ctx *gin.Context, accommodation Accommodation) error {
	formattedStartDate, _ := time.Parse("2006-01-02", accommodation.StartDate)
	formattedEndDate, _ := time.Parse("2006-01-02", accommodation.EndDate)

	params := queries.CreateAccommodationParams{
		Tripid:		int32(accommodation.TripId),
		Title:		accommodation.Title,
		Address:	sql.NullString{String: accommodation.Address, Valid: true},
		StartDate:	sql.NullTime{Time: formattedStartDate, Valid: true},
		EndDate:	sql.NullTime{Time: formattedEndDate, Valid: true},
		Url:		sql.NullString{String: accommodation.Url, Valid: true},
		Phone:		sql.NullString{String: accommodation.Phone, Valid: true},
	}

	err := r.queries.CreateAccommodation(ctx, params)
	if err != nil {
		return fmt.Errorf("error creating accommodation", err)
	}
	return nil
}