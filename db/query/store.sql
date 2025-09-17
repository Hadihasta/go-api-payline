-- name: CreateStore :one
INSERT INTO stores (
    store_access_id,
    name
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetStore :one
SELECT * FROM stores
WHERE name = $1
LIMIT 1;