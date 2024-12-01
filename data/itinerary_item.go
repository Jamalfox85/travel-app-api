package data

import (
	"database/sql"
	"fmt"
	"time"

	"travel-app-api/data/queries"

	"github.com/gin-gonic/gin"
)

type ItineraryItem struct {
	Id			int
	TripId		int
	Title		string
	Date		string
	Url			string
	Phone		string
	Address		string
	PoiId		string
	Rating		int
	Price		int
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
			Id:			int(row.Itemid),
			TripId:		int(row.Tripid.Int32),
			Title:		row.Title.String,
			Date:		row.Date.Time.Format("2006-01-02"),
			Url:		row.Url.String,
			Phone:		row.Phone.String,
			Address:	row.Address.String,
			PoiId:		row.Poiid.String,
			Rating: 	int(row.Rating.Int32),
			Price:		int(row.Price.Int32),

		}
		itineraryItems = append(itineraryItems, item)
	}

	return itineraryItems, nil
}

func (r *ItineraryItemRepository) CreateItineraryItem(ctx *gin.Context, item ItineraryItem) error {
	formattedDate, _ := time.Parse("2006-01-02", item.Date)

	params := queries.CreateItineraryItemParams{
		Tripid:		sql.NullInt32{Int32: int32(item.TripId), Valid: item.TripId != 0},
		Title:		sql.NullString{String: item.Title, Valid: item.Title != ""},
		Date:		sql.NullTime{Time: formattedDate, Valid: true},
		Url:		sql.NullString{String: item.Url, Valid: item.Url != ""},
		Phone:		sql.NullString{String: item.Phone, Valid: item.Phone != ""},
		Address:	sql.NullString{String: item.Address, Valid: item.Address != ""},
		Poiid:		sql.NullString{String: item.PoiId, Valid: item.PoiId != ""},
		Rating:		sql.NullInt32{Int32: int32(item.Rating), Valid: item.Rating != 0},
		Price:		sql.NullInt32{Int32: int32(item.Price), Valid: item.Price != 0},

	}
	
	err := r.queries.CreateItineraryItem(ctx, params)
	if err != nil {
		return fmt.Errorf("error creating itinerary item", err)
	}

	return nil
}