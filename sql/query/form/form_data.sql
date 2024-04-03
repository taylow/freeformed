-- name: CreateFormData :one
-- CreateFormData creates a new Form entry
INSERT INTO "form_data" (
    id,
    form_id,
    status,
    data
)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: ReadFormData :one
-- ReadFormData fetches a single FormData by ID
SELECT *
FROM "form_data"
WHERE id = $1
LIMIT 1;
-- name: ListFormDataByFormID :many
-- ListFormDataByFormID fetches a list of FormData by their form ID
SELECT *
FROM "form_data"
WHERE form_id = $1;
-- name: UpdateStatus :one
-- UpdateStatus updates the status of a Form
UPDATE "form_data"
SET status = $2,
  updated_at = now()
WHERE id = $1
RETURNING *;
-- name: DeleteFormData :one
-- DeleteFormData deletes a FormData entry
DELETE FROM "form_data"
WHERE id = $1
RETURNING *;
