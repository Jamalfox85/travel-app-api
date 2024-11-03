package handlers

import (
	"net/http"
	"strconv"
	"travel-app-api/data"

	"github.com/gin-gonic/gin"
)

type TripFinder interface {
	FindTrips(*gin.Context, int) ([]data.Trip, error)
}

func GetTrips(trips TripFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, _ := strconv.Atoi(ctx.Param("userId"))
		tripArray, _ := trips.FindTrips(ctx, userId)

		ctx.IndentedJSON(http.StatusOK, tripArray)
	}
}