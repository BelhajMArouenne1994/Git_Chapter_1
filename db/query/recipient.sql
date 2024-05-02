-- name: GetRecipient :one
SELECT * FROM "recipient"
WHERE "ocdMasterId" = $1 LIMIT 1;

-- name: ListRecipients :many
SELECT * FROM "recipient"
ORDER BY "username";

-- name: CreateRecipient :one
INSERT INTO "recipient" (
  "ocdMasterId", "username", "role", "created_at"
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateRecipient :exec
UPDATE "recipient"
SET "username" = $2,
    role = $3,
    "created_at" = $4
WHERE "ocdMasterId" = $1
RETURNING *;

-- name: DeleteRecipient :exec
DELETE FROM "recipient"
WHERE "ocdMasterId" = $1;