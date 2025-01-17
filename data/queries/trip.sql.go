// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: trip.sql

package queries

import (
	"context"
	"database/sql"
)

const createTrip = `-- name: CreateTrip :exec
INSERT INTO trips (Title, Location, userId, start_date, end_date, place_id, photo_uri, latitude, longitude)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateTripParams struct {
	Title     sql.NullString
	Location  sql.NullString
	Userid    sql.NullInt32
	StartDate sql.NullTime
	EndDate   sql.NullTime
	PlaceID   sql.NullString
	PhotoUri  sql.NullString
	Latitude  sql.NullFloat64
	Longitude sql.NullFloat64
}

func (q *Queries) CreateTrip(ctx context.Context, arg CreateTripParams) error {
	_, err := q.db.ExecContext(ctx, createTrip,
		arg.Title,
		arg.Location,
		arg.Userid,
		arg.StartDate,
		arg.EndDate,
		arg.PlaceID,
		arg.PhotoUri,
		arg.Latitude,
		arg.Longitude,
	)
	return err
}

const getTripsByUser = `-- name: GetTripsByUser :many
SELECT tripid, title, location, userid, start_date, end_date, place_id, photo_uri, latitude, longitude FROM trips
WHERE userId = ?
`

func (q *Queries) GetTripsByUser(ctx context.Context, userid sql.NullInt32) ([]Trip, error) {
	rows, err := q.db.QueryContext(ctx, getTripsByUser, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Trip
	for rows.Next() {
		var i Trip
		if err := rows.Scan(
			&i.Tripid,
			&i.Title,
			&i.Location,
			&i.Userid,
			&i.StartDate,
			&i.EndDate,
			&i.PlaceID,
			&i.PhotoUri,
			&i.Latitude,
			&i.Longitude,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
