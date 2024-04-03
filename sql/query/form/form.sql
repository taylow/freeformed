-- name: CreateForm :one
-- CreateForm creates a new Form entry
INSERT INTO "form" (
    id,
    project_id,
    name,
    description,
    color,
    enabled,
    file_uploads_enabled,
    redirect_url
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
-- name: Form :one
-- Form fetches a single Form by ID
SELECT *
FROM "form"
WHERE id = $1
LIMIT 1;
-- name: ListForms :many
-- ListForms fetches a list of Forms by their IDs
SELECT *
FROM "form"
WHERE id = ANY(sqlc.arg(ids)::uuid []);
-- name: ListFormsByProject :many
-- ListFormsByProject fetches a list of Forms by their project ID
SELECT *
FROM "form"
WHERE project_id = $1;
-- name: UpdateForm :one
-- UpdateForm updates the details of a Form
UPDATE "form"
SET name = $2,
  description = $3,
  color = $4,
  enabled = $5,
  file_uploads_enabled = $6,
  redirect_url = $7,
  updated_at = now()
WHERE id = $1
RETURNING *;
-- name: DeleteForm :one
-- DeleteForm deletes a Form
DELETE FROM "form"
WHERE id = $1
RETURNING *;
-- name: SoftDeleteForm :exec
-- SoftDeleteForm sets the deleted_at timestamp to now(), indicating that the Form is deleted
UPDATE "form"
SET deleted_at = now(),
  updated_at = now()
WHERE id = $1;
-- name: UnsoftDeleteForm :exec
-- UnsoftDeleteForm sets the deleted_at timestamp to NULL, indicating that the Form is not deleted
UPDATE "form"
SET soft_deleted_at = NULL,
  updated_at = now()
WHERE id = $1;
-- name: EnableForm :exec
-- EnableForm sets the enabled field to true
UPDATE "form"
SET enabled = true,
  updated_at = now()
WHERE id = $1;
-- name: DisableForm :exec
-- DisableForm sets the enabled field to false
UPDATE "form"
SET enabled = false,
  updated_at = now()
WHERE id = $1;
-- name: EnableFileUploads :exec
-- EnableFileUploads sets the file_uploads_enabled field to true
UPDATE "form"
SET file_uploads_enabled = true,
  updated_at = now()
WHERE id = $1;
-- name: DisableFileUploads :exec
-- DisableFileUploads sets the file_uploads_enabled field to false
UPDATE "form"
SET file_uploads_enabled = false,
  updated_at = now()
WHERE id = $1;
-- name: SetFormName :exec
-- SetFormName sets the name field to the given value
UPDATE "form"
SET name = $2,
  updated_at = now()
WHERE id = $1;
-- name: SetFormDescription :exec
-- SetFormDescription sets the description field to the given value
UPDATE "form"
SET description = $2,
  updated_at = now()
WHERE id = $1;
-- name: SetFormColor :exec
-- SetFormColor sets the color field to the given value
UPDATE "form"
SET color = $2,
  updated_at = now()
WHERE id = $1;
-- name: SetFormRedirectUrl :exec
-- SetFormRedirectUrl sets the redirect_url field to the given value
UPDATE "form"
SET redirect_url = $2,
  updated_at = now()
WHERE id = $1;
-- name: CheckFormExists :one
-- CheckFormExists checks if a Form exists by ID
SELECT EXISTS (
  SELECT 1
  FROM "form"
  WHERE id = $1
);