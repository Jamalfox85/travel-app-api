package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"travel-app-api/data"

	"github.com/gin-gonic/gin"
)

type ItineraryItemFinder interface {
	FindItineraryItems(*gin.Context, int) ([]data.ItineraryItem, error)
}

func GetItineraryItems(itineraryItems ItineraryItemFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tripId, _ := strconv.Atoi(ctx.Param("tripId"))
		fmt.Println("tripId", tripId);
		itemArray, _ := itineraryItems.FindItineraryItems(ctx, tripId)

		ctx.IndentedJSON(http.StatusOK, itemArray)
	}
}