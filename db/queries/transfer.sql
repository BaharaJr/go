-- name: CreateTransfer :one
INSERT INTO TRANSFER (
  sender,
  receiver,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM TRANSFER
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM TRANSFER
WHERE 
    sender = $1 OR
    receiver = $2
ORDER BY id
LIMIT $3
OFFSET $4;