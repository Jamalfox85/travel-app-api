package data

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"travel-app-api/data/queries"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID			int
	FirstName	string
	LastName	string
	Email		string
	Username	string
	Preferences	json.RawMessage
}

type UserRepository struct {
	queries *queries.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	queries := queries.New(db)

	return &UserRepository{
		queries: 	queries,
	}
}

func (r *UserRepository) FindUser(ctx *gin.Context, userId int) (User, error) {
	formattedUserId := int32(userId)
	fmt.Println(formattedUserId)

	row, err := r.queries.GetUser(ctx, 1)
	if err != nil {
		return User{}, fmt.Errorf("error fetching user details", err)
	}


	user := User{
		ID:				int(row.Userid),
		FirstName:		row.Firstname.String,
		LastName:		row.Lastname.String,
		Email:			row.Email.String,
		Username:		row.Username.String,
		Preferences:	row.Preferences,
	}


	return user, nil
}