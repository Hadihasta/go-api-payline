-- name: CreateUser :one
INSERT INTO users (
  email,
  phone_number,
  name
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;
