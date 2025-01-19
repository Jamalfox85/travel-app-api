package data

import (
	"database/sql"
	"fmt"

	"travel-app-api/data/queries"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserID		int
	FirstName	string
	LastName	string
	Email		string
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

func (r *UserRepository) AuthorizeUser(ctx *gin.Context, activeUser User) (User, error) {

	email := sql.NullString{String: activeUser.Email, Valid: activeUser.Email != ""}
	row, _ := r.queries.GetUser(ctx, email)


	if row.UserID != 0 {
		userDetails := User{
			UserID:				int(row.UserID),
			FirstName:		row.FirstName.String,
			LastName:		row.LastName.String,
			Email:			row.Email.String,
		}
		return userDetails, nil		
		
	} else {
		err := r.CreateUser(ctx, activeUser)
		if err != nil {
			return User{}, fmt.Errorf("error creating new user", err)
		}

		id, _ := r.queries.GetLastInsertId(ctx)
		activeUser.UserID = int(id)

		return activeUser, nil
	}

}

func (r *UserRepository) CreateUser(ctx *gin.Context, newUser User) (error) {

	params := queries.CreateUserParams{
		FirstName: sql.NullString{String: newUser.FirstName, Valid: newUser.FirstName != ""},
		LastName: sql.NullString{String: newUser.LastName, Valid: newUser.LastName != ""},
		Email: sql.NullString{String: newUser.Email, Valid: newUser.Email != ""},
	}

	err := r.queries.CreateUser(ctx, params);
	if err != nil {
		fmt.Println(err);
		return fmt.Errorf("error creating new user");
	}
	return nil
}