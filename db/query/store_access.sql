-- name: CreateStoreAccess :one
INSERT INTO stores_access (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetStoreAccess :one
SELECT * FROM stores_access
WHERE id = $1 LIMIT 1;
