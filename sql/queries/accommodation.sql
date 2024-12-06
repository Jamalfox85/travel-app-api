-- name: GetAccommodations :many
SELECT * FROM accommodations
WHERE tripId = ?;

-- name: CreateAccommodation :exec
INSERT INTO accommodations (tripId, title, address, start_date, end_date, url, phone)
VALUES (?, ?, ?, ?, ?, ?, ?);