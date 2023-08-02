-- name: CreateProduct :one
INSERT INTO products (
  name,
  brand,
  price,
  quantity
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProductForUpdate :one
SELECT * FROM products
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListProducts :many
SELECT * FROM products
WHERE store_id = $1
ORDER BY id;

-- name: UpdateProduct :one
UPDATE products
SET price = $2, quantity = $3
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;