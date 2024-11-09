// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package queries

import (
	"database/sql"
	"encoding/json"
)

type ItineraryItem struct {
	Itemid   int32
	Tripid   sql.NullInt32
	Title    sql.NullString
	Location sql.NullString
	Date     sql.NullTime
}

type Trip struct {
	Tripid    int32
	Title     sql.NullString
	Location  sql.NullString
	Userid    sql.NullInt32
	StartDate sql.NullTime
	EndDate   sql.NullTime
}

type User struct {
	Userid      int32
	Firstname   sql.NullString
	Lastname    sql.NullString
	Email       sql.NullString
	Username    sql.NullString
	Preferences json.RawMessage
}
