package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"travel-app-api/data"

	"github.com/gin-gonic/gin"
)

type AccommodationFinder interface {
	FindAccommodations(*gin.Context, int) ([]data.Accommodation, error)
	CreateAccommodation(*gin.Context, data.Accommodation) (error)
}

func GetAccommodations(accommodations AccommodationFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tripId, _ := strconv.Atoi(ctx.Param("tripId"))
		accommodationArray, _ := accommodations.FindAccommodations(ctx, tripId)

		ctx.IndentedJSON(http.StatusOK, accommodationArray)
	}
}

func CreateAccommodation(accommodations AccommodationFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newAccommodation data.Accommodation

		if err := ctx.ShouldBindJSON(&newAccommodation);
		err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := accommodations.CreateAccommodation(ctx, newAccommodation)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("New Accommodation", newAccommodation);
		ctx.JSON(http.StatusCreated, gin.H{"message": "New Accommodation Created Successfully!"})
	}
}