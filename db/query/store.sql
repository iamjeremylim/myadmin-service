-- name: CreateStore :one
INSERT INTO stores (
  owner,
  name
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetStore :one
SELECT * FROM stores
WHERE id = $1 LIMIT 1;

-- name: GetStoreForUpdate :one
SELECT * FROM stores
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListStores :many
SELECT * FROM stores
WHERE owner = $1
ORDER BY id;

-- name: UpdateStore :one
UPDATE stores
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1;