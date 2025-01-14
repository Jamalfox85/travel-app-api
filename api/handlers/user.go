package handlers

import (
	"fmt"
	"net/http"
	"travel-app-api/data"

	"github.com/gin-gonic/gin"
)

type UserFinder interface {
	AuthorizeUser(*gin.Context, data.User) (data. User, error)
}

// func GetUser(users UserFinder) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		userId, _ := strconv.Atoi(ctx.Param("userId"))

// 		user, _ := users.FindUser(ctx, userId)

// 		ctx.IndentedJSON(http.StatusOK, user)
// 	}
// }

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

		fmt.Println("User:", activeUser)
		ctx.IndentedJSON(http.StatusOK, userDetails)
	}
}

