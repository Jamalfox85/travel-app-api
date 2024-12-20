// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package queries

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT userid, firstname, lastname, email, username, preferences FROM users
WHERE UserID = ?
`

func (q *Queries) GetUser(ctx context.Context, userid int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userid)
	var i User
	err := row.Scan(
		&i.Userid,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Username,
		&i.Preferences,
	)
	return i, err
}
