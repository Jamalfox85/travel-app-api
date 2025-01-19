package handlers

import (
	"net/http"
	"travel-app-api/data"

	"github.com/gin-gonic/gin"
)

type UserFinder interface {
	AuthorizeUser(*gin.Context, data.User) (data. User, error)
}

func AuthorizeUser(users UserFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var activeUser data.User

		if err := ctx.ShouldBindJSON(&activeUser);
		err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userDetails, err := users.AuthorizeUser(ctx, activeUser)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, userDetails)
	}
}

