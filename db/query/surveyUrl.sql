-- name: GetSurveyUrlById :one
SELECT *
FROM "surveyUrl"
WHERE "ocdB2cSurveyUrlId" = $1 LIMIT 1;

-- name: ListNewestSurveyUrls :many
SELECT *
FROM "surveyUrl"
ORDER BY "technicalCreationDate" DESC
LIMIT 10;


-- name: ListNewestSurveyUrlsByBrandAndCountry :many
SELECT *
FROM "surveyUrl"
WHERE "brand" = $1 AND "country" = $2
ORDER BY "technicalCreationDate" DESC
LIMIT 10;


-- name: CreateSurveyUrl :one
INSERT INTO "surveyUrl" (
  "brand",
  "country",
  "ocdB2cSurveyUrlId",
  "technicalCreationDate",
  "technicalLastUpdateDate",
  "sourceName",
  "sourceSurveyId",
  "surveyId",
  "scenario",
  "channel",
  "surveyLanguage",
  "ocdMasterId",
  "relatedObjectName",
  "relatedObjectId",
  "url"
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
)
RETURNING *;

-- name: UpdateSurveyUrl :one
UPDATE "surveyUrl"
SET
  "brand" = $1,
  "country" = $2,
  "technicalCreationDate" = $3,
  "technicalLastUpdateDate" = $4,
  "sourceName" = $5,
  "sourceSurveyId" = $6,
  "surveyId" = $7,
  "scenario" = $8,
  "channel" = $9,
  "surveyLanguage" = $10,
  "relatedObjectName" = $11,
  "relatedObjectId" = $12,
  "url" = $13
WHERE "ocdB2cSurveyUrlId" = $14
RETURNING *;


-- name: DeleteSurveyUrl :exec
DELETE FROM "surveyUrl"
WHERE "ocdB2cSurveyUrlId" = $1;

-- name: FindSurveyUrlsByBrandAndCountry :many
SELECT *
FROM "surveyUrl"
WHERE "brand" = $1 AND "country" = $2
ORDER BY "technicalCreationDate" DESC;


-- name: LinkSurveyUrlToRecipient :one
UPDATE "surveyUrl"
SET "ocdMasterId" = $1
WHERE "ocdB2cSurveyUrlId" = $2
RETURNING *;