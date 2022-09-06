-- name: CreateAccount :one
INSERT INTO account (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccountById :one
SELECT * FROM ACCOUNT WHERE ID=$1 LIMIT 1;

-- name: GetAccountByOwner :one
SELECT * FROM ACCOUNT WHERE OWNER=$1 LIMIT 1;

-- name: GetAccounts :many
SELECT * FROM ACCOUNT 
ORDER BY CREATED ASC
 LIMIT $1
 OFFSET $2;

 -- name: UpdateAccount :one
UPDATE ACCOUNT 
SET BALANCE = $2
WHERE ID = $1
RETURNING *
;

-- name: GenericSearch :many
SELECT * FROM ACCOUNT WHERE $1 = $2 
ORDER BY CREATED ASC
 LIMIT $1
 OFFSET $2;

-- name: DeleteAccount :exec
DELETE FROM ACCOUNT WHERE id = $1;