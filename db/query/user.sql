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

-- name: GetListUsersWithRole :many
SELECT 
  u.id,
  u.role_id,
  r.role_name,
  u.email,
  u.phone_number,
  u.name,
  u.is_active,
  u.created_at,
  u.updated_at
FROM users u
JOIN roles r ON u.role_id = r.id
ORDER BY u.id ASC;