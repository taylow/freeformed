-- name: CreateUser :one
-- CreateUser creates a new User entry
INSERT INTO "user" (
    id,
    first_name,
    last_name,
    email,
    enabled
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: ReadUser :one
-- ReadUser fetches a single User by ID
SELECT *
FROM "user"
WHERE id = $1
LIMIT 1;
-- name: UpdateUser :one
-- UpdateUser updates the details of a User
UPDATE "user"
SET first_name = $2,
  last_name = $3,
  email = $4,
  updated_at = now()
WHERE id = $1
RETURNING *;
-- name: DeleteUser :one
-- DeleteUser deletes a User
DELETE FROM "user"
WHERE id = $1
RETURNING *;
-- name: SoftDeleteUser :exec
-- SoftDeleteUser sets the deleted_at timestamp to now(), indicating that the User is deleted
UPDATE "user"
SET deleted_at = now(),
  updated_at = now()
WHERE id = $1;
-- name: UnsoftDeleteUser :exec
-- UnsoftDeleteUser sets the deleted_at timestamp to NULL, indicating that the User is not deleted
UPDATE "user"
SET soft_deleted_at = NULL,
  updated_at = now()
WHERE id = $1;
-- name: EnableUser :exec
-- EnableUser sets the enabled field to true
UPDATE "user"
SET enabled = true,
  updated_at = now()
WHERE id = $1;
-- name: DisableUser :exec
-- DisableUser sets the enabled field to false
UPDATE "user"
SET enabled = false,
  updated_at = now()
WHERE id = $1;
