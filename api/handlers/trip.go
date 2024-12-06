package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"travel-app-api/data"

	"github.com/gin-gonic/gin"
)

type TripFinder interface {
	FindTrips(*gin.Context, int) ([]data.Trip, error)
	CreateTrip(*gin.Context, data.Trip) (error)
}

func GetTrips(trips TripFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, _ := strconv.Atoi(ctx.Param("userId"))
		tripArray, _ := trips.FindTrips(ctx, userId)

		ctx.IndentedJSON(http.StatusOK, tripArray)
	}
}

func CreateTrip(trips TripFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newTrip data.Trip
		
		if err := ctx.ShouldBindJSON(&newTrip); 
		err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := trips.CreateTrip(ctx, newTrip)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

    	fmt.Println("New Trip:", newTrip)
    	ctx.JSON(http.StatusCreated, gin.H{"message": "New Trip Created Successfully!"})

	}
}