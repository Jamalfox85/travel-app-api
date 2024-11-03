package handlers

import (
	"net/http"
	"strconv"
	"travel-app-api/data"

	"github.com/gin-gonic/gin"
)

type UserFinder interface {
	FindUser(*gin.Context, int) (data.User, error)
}

func GetUser(users UserFinder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, _ := strconv.Atoi(ctx.Param("userId"))

		user, _ := users.FindUser(ctx, userId)

		ctx.IndentedJSON(http.StatusOK, user)
	}
}

