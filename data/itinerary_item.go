package data

import (
	"database/sql"
	"fmt"

	"travel-app-api/data/queries"

	"github.com/gin-gonic/gin"
)

type ItineraryItem struct {
	ID			int
	TripID		int
	Title		string
	Location	string
	Date		string
}

type ItineraryItemRepository struct {
	queries *queries.Queries
}

func NewItineraryItemRepository(db *sql.DB) *ItineraryItemRepository {
	queries := queries.New(db)

	return &ItineraryItemRepository{
		queries: 	queries,
	}
}

func (r *ItineraryItemRepository) FindItineraryItems(ctx *gin.Context, tripId int) ([]ItineraryItem, error) {
	formattedTripId := NewNullInt32(int32(tripId))
	fmt.Println("formattedTripId", formattedTripId)
	rows, err := r.queries.GetItineraryItems(ctx, formattedTripId)
	if err != nil {
		return nil, fmt.Errorf("error fetching itinerary items for user", err)
	}

	var itineraryItems []ItineraryItem
	for _, row := range rows {
		item := ItineraryItem{
			ID:			int(row.Itemid),
			TripID:		int(row.Tripid.Int32),
			Title:		row.Title.String,
			Location:	row.Location.String,
			Date:		row.Date.Time.Format("MM-DD-YY"),

		}
		itineraryItems = append(itineraryItems, item)
	}

	return itineraryItems, nil
}