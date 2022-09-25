-- name: CreateUser :exec
INSERT INTO users (id,
                   full_name,
                   email,
                   password)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUser :one
SELECT id, full_name, email, created_at
FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdatePassword :one
UPDATE users
set password = $2
WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;
