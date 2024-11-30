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
	CreateItineraryItem(*gin.Context, data.ItineraryItem) (error)
}

func GetItineraryItems(itineraryItems ItineraryItemFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tripId, _ := strconv.Atoi(ctx.Param("tripId"))
		itemArray, _ := itineraryItems.FindItineraryItems(ctx, tripId)

		fmt.Println("itemArray", itemArray)
		ctx.IndentedJSON(http.StatusOK, itemArray)
	}
}

func CreateItineraryItem(itineraryItems ItineraryItemFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newItem data.ItineraryItem
		
		if err := ctx.ShouldBindJSON(&newItem); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := itineraryItems.CreateItineraryItem(ctx, newItem)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("New Itinerary Item:", newItem)
		ctx.JSON(http.StatusCreated, gin.H{"message": "New Itinerary Item Created Successfully!"})

	}
}