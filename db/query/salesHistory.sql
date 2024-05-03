-- name: GetSalesHistoryById :one
SELECT *
FROM "salesHistory"
WHERE "ocdTicketId" = $1
LIMIT 1;

-- name: ListNewestSalesHistory :many
SELECT *
FROM "salesHistory"
ORDER BY "technicalLastUpdateDate" DESC
LIMIT 10;

-- name: CreateSalesHistory :one
INSERT INTO "salesHistory" (
  "brand",
  "country",
  "ocdTicketId",
  "technicalCreationDate",
  "technicalLastUpdateDate",
  "source",
  "sourceName",
  "sourceChannel",
  "sourcePersonId",
  "sourceTicketNumber",
  "sourceStoreType",
  "sourceStatusOrder",
  "ocdContactMasterId",
  "ocdContactVersionId",
  "ocdStoreId"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING *;

-- name: UpdateSalesHistory :one
UPDATE "salesHistory"
SET
  "brand" = $1,
  "country" = $2,
  "technicalCreationDate" = $3,
  "technicalLastUpdateDate" = $4,
  "source" = $5,
  "sourceName" = $6,
  "sourceChannel" = $7,
  "sourcePersonId" = $8,
  "sourceTicketNumber" = $9,
  "sourceStoreType" = $10,
  "sourceStatusOrder" = $11,
  "ocdContactMasterId" = $12,
  "ocdContactVersionId" = $13,
  "ocdStoreId" = $14
WHERE "ocdTicketId" = $15
RETURNING *;

-- name: DeleteSalesHistory :exec
DELETE FROM "salesHistory"
WHERE "ocdTicketId" = $1;

-- name: FindSalesHistoryByBrand :many
SELECT *
FROM "salesHistory"
WHERE "brand" = $1 AND "country" = $2
ORDER BY "technicalLastUpdateDate" DESC
LIMIT 10;

-- name: FindSalesHistoryByCountry :many
SELECT *
FROM "salesHistory"
WHERE "country" = $1
ORDER BY "technicalLastUpdateDate" DESC
LIMIT 10;
