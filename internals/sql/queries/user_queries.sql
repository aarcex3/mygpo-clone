-- name: GetUser :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (username, password, email) values (?,?,?);