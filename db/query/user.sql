-- name: CreateUser :one
INSERT INTO users (
  role_id,
  email,
  phone_number,
  name
) VALUES (
  $1, $2, $3,$4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;
