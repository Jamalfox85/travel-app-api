-- name: GetTripsByUser :many
SELECT * FROM Trips
WHERE userId = ?;