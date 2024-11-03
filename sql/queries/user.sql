-- name: GetUser :one
SELECT * FROM Users
WHERE UserID = ?;