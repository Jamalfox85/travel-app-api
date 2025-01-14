// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package queries

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (first_name, last_name, email)
VALUES (?, ?, ?)
`

type CreateUserParams struct {
	FirstName sql.NullString
	LastName  sql.NullString
	Email     sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.FirstName, arg.LastName, arg.Email)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_id, first_name, last_name, email FROM users
WHERE email = ?
`

func (q *Queries) GetUser(ctx context.Context, email sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
	)
	return i, err
}
