-- name: GetUserById :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ? LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (username, password, email) values (?,?,?);

-- name: UserExists :one
SELECT count(*) FROM users WHERE username = ? OR email = ?;