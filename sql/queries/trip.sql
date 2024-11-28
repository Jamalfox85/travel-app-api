-- name: GetTripsByUser :many
SELECT * FROM trips
WHERE userId = ?;

-- name: CreateTrip :exec
INSERT INTO trips (Title, Location, userId, start_date, end_date, place_id, photo_uri, latitude, longitude)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);