-- name: CreateStore :one
INSERT INTO stores (
    name
) VALUES (
    $1
)
RETURNING *;

-- name: GetStore :one
SELECT * FROM stores
WHERE name = $1
LIMIT 1;