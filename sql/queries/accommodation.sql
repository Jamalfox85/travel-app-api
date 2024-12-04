-- name: GetAccommodations :many
SELECT * FROM accommodations
WHERE tripId = ?;