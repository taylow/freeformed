// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: project.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const changeProjectTeam = `-- name: ChangeProjectTeam :exec
UPDATE "project"
SET team_id = $2,
  updated_at = now()
WHERE id = $1
RETURNING id, team_id, name, description, enabled, created_at, updated_at, deleted_at
`

type ChangeProjectTeamParams struct {
	ID     uuid.UUID `db:"id" json:"id"`
	TeamID uuid.UUID `db:"team_id" json:"team_id"`
}

// ChangeProjectTeam changes the team of a Project
func (q *Queries) ChangeProjectTeam(ctx context.Context, arg ChangeProjectTeamParams) error {
	_, err := q.exec(ctx, q.changeProjectTeamStmt, changeProjectTeam, arg.ID, arg.TeamID)
	return err
}

const createProject = `-- name: CreateProject :one
INSERT INTO "project" (
    id,
    team_id,
    name,
    description,
    enabled
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, team_id, name, description, enabled, created_at, updated_at, deleted_at
`

type CreateProjectParams struct {
	ID          uuid.UUID `db:"id" json:"id"`
	TeamID      uuid.UUID `db:"team_id" json:"team_id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Enabled     bool      `db:"enabled" json:"enabled"`
}

// CreateProject creates a new Project entry
func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.queryRow(ctx, q.createProjectStmt, createProject,
		arg.ID,
		arg.TeamID,
		arg.Name,
		arg.Description,
		arg.Enabled,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.Name,
		&i.Description,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteProject = `-- name: DeleteProject :one
DELETE FROM "project"
WHERE id = $1
RETURNING id, team_id, name, description, enabled, created_at, updated_at, deleted_at
`

// DeleteProject deletes a Project
func (q *Queries) DeleteProject(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.queryRow(ctx, q.deleteProjectStmt, deleteProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.Name,
		&i.Description,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const disableProject = `-- name: DisableProject :exec
UPDATE "project"
SET enabled = false,
  updated_at = now()
WHERE id = $1
`

// DisableProject sets the enabled field to false
func (q *Queries) DisableProject(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.disableProjectStmt, disableProject, id)
	return err
}

const enableProject = `-- name: EnableProject :exec
UPDATE "project"
SET enabled = true,
  updated_at = now()
WHERE id = $1
`

// EnableProject sets the enabled field to true
func (q *Queries) EnableProject(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.enableProjectStmt, enableProject, id)
	return err
}

const listProjects = `-- name: ListProjects :many
SELECT id, team_id, name, description, enabled, created_at, updated_at, deleted_at
FROM "project"
WHERE id = ANY($1::uuid [])
`

// ListProjects fetches a list of Projects by their IDs
func (q *Queries) ListProjects(ctx context.Context, ids []uuid.UUID) ([]Project, error) {
	rows, err := q.query(ctx, q.listProjectsStmt, listProjects, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.TeamID,
			&i.Name,
			&i.Description,
			&i.Enabled,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProjectsByTeam = `-- name: ListProjectsByTeam :many
SELECT id, team_id, name, description, enabled, created_at, updated_at, deleted_at
FROM "project"
WHERE team_id = $1
`

// ListProjectsByTeam fetches a list of Projects by team ID
func (q *Queries) ListProjectsByTeam(ctx context.Context, teamID uuid.UUID) ([]Project, error) {
	rows, err := q.query(ctx, q.listProjectsByTeamStmt, listProjectsByTeam, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.TeamID,
			&i.Name,
			&i.Description,
			&i.Enabled,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readProject = `-- name: ReadProject :one
SELECT id, team_id, name, description, enabled, created_at, updated_at, deleted_at
FROM "project"
WHERE id = $1
LIMIT 1
`

// ReadProject fetches a single Project by ID
func (q *Queries) ReadProject(ctx context.Context, id uuid.UUID) (Project, error) {
	row := q.queryRow(ctx, q.readProjectStmt, readProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.Name,
		&i.Description,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const softDeleteProject = `-- name: SoftDeleteProject :exec
UPDATE "project"
SET deleted_at = now(),
  updated_at = now()
WHERE id = $1
`

// SoftDeleteProject sets the deleted_at timestamp to now(), indicating that the Project is deleted
func (q *Queries) SoftDeleteProject(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.softDeleteProjectStmt, softDeleteProject, id)
	return err
}

const unsoftDeleteProject = `-- name: UnsoftDeleteProject :exec
UPDATE "project"
SET soft_deleted_at = NULL,
  updated_at = now()
WHERE id = $1
`

// UnsoftDeleteProject sets the deleted_at timestamp to NULL, indicating that the Project is not deleted
func (q *Queries) UnsoftDeleteProject(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.unsoftDeleteProjectStmt, unsoftDeleteProject, id)
	return err
}

const updateProject = `-- name: UpdateProject :one
UPDATE "project"
SET name = $2,
  description = $3,
  updated_at = now()
WHERE id = $1
RETURNING id, team_id, name, description, enabled, created_at, updated_at, deleted_at
`

type UpdateProjectParams struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
}

// UpdateProject updates the details of a Project
func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (Project, error) {
	row := q.queryRow(ctx, q.updateProjectStmt, updateProject, arg.ID, arg.Name, arg.Description)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.Name,
		&i.Description,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
