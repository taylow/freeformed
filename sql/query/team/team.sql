-- name: CreateTeam :one
-- CreateTeam creates a new Team entry
INSERT INTO "team" (
    id,
    name,
    description,
    enabled,
    owner_id
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: ReadTeam :one
-- ReadTeam fetches a single Team by ID
SELECT *
FROM "team"
WHERE id = $1
LIMIT 1;
-- name: ListTeams :many
-- ListTeams fetches a list of Teams by their IDs
SELECT *
FROM "team"
WHERE id = ANY(sqlc.arg(ids)::uuid []);
-- name: UpdateTeam :one
-- UpdateTeam updates the details of a Team
UPDATE "team"
SET name = $2,
  description = $3,
  updated_at = now()
WHERE id = $1
RETURNING *;
-- name: DeleteTeam :one
-- DeleteTeam deletes a Team
DELETE FROM "team"
WHERE id = $1
RETURNING *;
-- name: SoftDeleteTeam :exec
-- SoftDeleteTeam sets the deleted_at timestamp to now(), indicating that the Team is deleted
UPDATE "team"
SET deleted_at = now(),
  updated_at = now()
WHERE id = $1;
-- name: UnsoftDeleteTeam :exec
-- UnsoftDeleteTeam sets the deleted_at timestamp to NULL, indicating that the Team is not deleted
UPDATE "team"
SET soft_deleted_at = NULL,
  updated_at = now()
WHERE id = $1;
-- name: EnableTeam :exec
-- EnableTeam sets the enabled field to true
UPDATE "team"
SET enabled = true,
  updated_at = now()
WHERE id = $1;
-- name: DisableTeam :exec
-- DisableTeam sets the enabled field to false
UPDATE "team"
SET enabled = false,
  updated_at = now()
WHERE id = $1;
-- name: ChangeTeamOwner :exec
-- ChangeTeamOwner changes the owner of a Team
UPDATE "team"
SET owner_id = $2,
  updated_at = now()
WHERE id = $1
RETURNING *;
-- name: ReadTeamByOwner :one
-- ReadTeamByOwner fetches a single Team by owner_id
SELECT *
FROM "team"
WHERE owner_id = $1;