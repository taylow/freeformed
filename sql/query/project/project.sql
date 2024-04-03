-- name: CreateProject :one
-- CreateProject creates a new Project entry
INSERT INTO "project" (
    id,
    team_id,
    name,
    description,
    enabled
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: ReadProject :one
-- ReadProject fetches a single Project by ID
SELECT *
FROM "project"
WHERE id = $1
LIMIT 1;
-- name: ListProjects :many
-- ListProjects fetches a list of Projects by their IDs
SELECT *
FROM "project"
WHERE id = ANY(sqlc.arg(ids)::uuid []);
-- name: ListProjectsByTeam :many
-- ListProjectsByTeam fetches a list of Projects by team ID
SELECT *
FROM "project"
WHERE team_id = $1;
-- name: UpdateProject :one
-- UpdateProject updates the details of a Project
UPDATE "project"
SET name = $2,
  description = $3,
  updated_at = now()
WHERE id = $1
RETURNING *;
-- name: DeleteProject :one
-- DeleteProject deletes a Project
DELETE FROM "project"
WHERE id = $1
RETURNING *;
-- name: SoftDeleteProject :exec
-- SoftDeleteProject sets the deleted_at timestamp to now(), indicating that the Project is deleted
UPDATE "project"
SET deleted_at = now(),
  updated_at = now()
WHERE id = $1;
-- name: UnsoftDeleteProject :exec
-- UnsoftDeleteProject sets the deleted_at timestamp to NULL, indicating that the Project is not deleted
UPDATE "project"
SET soft_deleted_at = NULL,
  updated_at = now()
WHERE id = $1;
-- name: EnableProject :exec
-- EnableProject sets the enabled field to true
UPDATE "project"
SET enabled = true,
  updated_at = now()
WHERE id = $1;
-- name: DisableProject :exec
-- DisableProject sets the enabled field to false
UPDATE "project"
SET enabled = false,
  updated_at = now()
WHERE id = $1;
-- name: ChangeProjectTeam :exec
-- ChangeProjectTeam changes the team of a Project
UPDATE "project"
SET team_id = $2,
  updated_at = now()
WHERE id = $1
RETURNING *;
