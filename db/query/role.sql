-- format command harus  (-- name: <FunctionName> :<querytype>)

-- name: CreateRoles :one
INSERT INTO roles (
  role_name
) VALUES (
  $1
)
RETURNING *;

-- name: GetRoles :one
SELECT * FROM roles
WHERE id = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateRoles :one
UPDATE roles
  set role_name = $2
WHERE id = $1
RETURNING id, role_name;


-- name: DeleteRoles :exec
DELETE FROM roles
WHERE id = $1;