-- name: GetUser :one
SELECT * FROM users
WHERE email = ?;

-- name: CreateUser :exec
INSERT INTO users (first_name, last_name, email)
VALUES (?, ?, ?);

-- name: GetLastInsertId :one
SELECT LAST_INSERT_ID() AS user_id;
