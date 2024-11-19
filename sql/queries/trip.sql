-- name: GetTripsByUser :many
SELECT * FROM Trips
WHERE userId = ?;

-- name: CreateTrip :exec
INSERT INTO Trips (Title, Location, userId, start_date, end_date, place_id, photo_uri, latitude, longitude)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);