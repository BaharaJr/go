-- name: CreateEntry :one
INSERT INTO ENTRY (
  amount,
  account
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntryId :one
SELECT * FROM ENTRY WHERE ID=$1 LIMIT 1;

-- name: GetEntries :many
SELECT * FROM ENTRY 
ORDER BY CREATED ASC
 LIMIT $1
 OFFSET $2;

 -- name: UpdateEntry :one
UPDATE ENTRY 
SET AMOUNT = $2
WHERE ID = $1
RETURNING *
;

-- name: DeleteEntry :exec
DELETE FROM ENTRY WHERE id = $1;