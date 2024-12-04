package handlers

import (
	"net/http"
	"strconv"
	"travel-app-api/data"

	"github.com/gin-gonic/gin"
)

type AccommodationFinder interface {
	FindAccommodations(*gin.Context, int) ([]data.Accommodation, error)
}

func GetAccommodations(accommodations AccommodationFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tripId, _ := strconv.Atoi(ctx.Param("tripId"))
		accommodationArray, _ := accommodations.FindAccommodations(ctx, tripId)

		ctx.IndentedJSON(http.StatusOK, accommodationArray)
	}
}