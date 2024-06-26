// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: form.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const checkFormExists = `-- name: CheckFormExists :one
SELECT EXISTS (
  SELECT 1
  FROM "form"
  WHERE id = $1
)
`

// CheckFormExists checks if a Form exists by ID
func (q *Queries) CheckFormExists(ctx context.Context, id string) (bool, error) {
	row := q.queryRow(ctx, q.checkFormExistsStmt, checkFormExists, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createForm = `-- name: CreateForm :one
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
RETURNING id, project_id, name, description, color, enabled, file_uploads_enabled, redirect_url, created_at, updated_at, deleted_at
`

type CreateFormParams struct {
	ID                 string    `db:"id" json:"id"`
	ProjectID          uuid.UUID `db:"project_id" json:"project_id"`
	Name               string    `db:"name" json:"name"`
	Description        string    `db:"description" json:"description"`
	Color              string    `db:"color" json:"color"`
	Enabled            bool      `db:"enabled" json:"enabled"`
	FileUploadsEnabled bool      `db:"file_uploads_enabled" json:"file_uploads_enabled"`
	RedirectUrl        string    `db:"redirect_url" json:"redirect_url"`
}

// CreateForm creates a new Form entry
func (q *Queries) CreateForm(ctx context.Context, arg CreateFormParams) (Form, error) {
	row := q.queryRow(ctx, q.createFormStmt, createForm,
		arg.ID,
		arg.ProjectID,
		arg.Name,
		arg.Description,
		arg.Color,
		arg.Enabled,
		arg.FileUploadsEnabled,
		arg.RedirectUrl,
	)
	var i Form
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Name,
		&i.Description,
		&i.Color,
		&i.Enabled,
		&i.FileUploadsEnabled,
		&i.RedirectUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteForm = `-- name: DeleteForm :one
DELETE FROM "form"
WHERE id = $1
RETURNING id, project_id, name, description, color, enabled, file_uploads_enabled, redirect_url, created_at, updated_at, deleted_at
`

// DeleteForm deletes a Form
func (q *Queries) DeleteForm(ctx context.Context, id string) (Form, error) {
	row := q.queryRow(ctx, q.deleteFormStmt, deleteForm, id)
	var i Form
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Name,
		&i.Description,
		&i.Color,
		&i.Enabled,
		&i.FileUploadsEnabled,
		&i.RedirectUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const disableFileUploads = `-- name: DisableFileUploads :exec
UPDATE "form"
SET file_uploads_enabled = false,
  updated_at = now()
WHERE id = $1
`

// DisableFileUploads sets the file_uploads_enabled field to false
func (q *Queries) DisableFileUploads(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.disableFileUploadsStmt, disableFileUploads, id)
	return err
}

const disableForm = `-- name: DisableForm :exec
UPDATE "form"
SET enabled = false,
  updated_at = now()
WHERE id = $1
`

// DisableForm sets the enabled field to false
func (q *Queries) DisableForm(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.disableFormStmt, disableForm, id)
	return err
}

const enableFileUploads = `-- name: EnableFileUploads :exec
UPDATE "form"
SET file_uploads_enabled = true,
  updated_at = now()
WHERE id = $1
`

// EnableFileUploads sets the file_uploads_enabled field to true
func (q *Queries) EnableFileUploads(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.enableFileUploadsStmt, enableFileUploads, id)
	return err
}

const enableForm = `-- name: EnableForm :exec
UPDATE "form"
SET enabled = true,
  updated_at = now()
WHERE id = $1
`

// EnableForm sets the enabled field to true
func (q *Queries) EnableForm(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.enableFormStmt, enableForm, id)
	return err
}

const form = `-- name: Form :one
SELECT id, project_id, name, description, color, enabled, file_uploads_enabled, redirect_url, created_at, updated_at, deleted_at
FROM "form"
WHERE id = $1
LIMIT 1
`

// Form fetches a single Form by ID
func (q *Queries) Form(ctx context.Context, id string) (Form, error) {
	row := q.queryRow(ctx, q.formStmt, form, id)
	var i Form
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Name,
		&i.Description,
		&i.Color,
		&i.Enabled,
		&i.FileUploadsEnabled,
		&i.RedirectUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listForms = `-- name: ListForms :many
SELECT id, project_id, name, description, color, enabled, file_uploads_enabled, redirect_url, created_at, updated_at, deleted_at
FROM "form"
WHERE id = ANY($1::uuid [])
`

// ListForms fetches a list of Forms by their IDs
func (q *Queries) ListForms(ctx context.Context, ids []uuid.UUID) ([]Form, error) {
	rows, err := q.query(ctx, q.listFormsStmt, listForms, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Form
	for rows.Next() {
		var i Form
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.Name,
			&i.Description,
			&i.Color,
			&i.Enabled,
			&i.FileUploadsEnabled,
			&i.RedirectUrl,
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

const listFormsByProject = `-- name: ListFormsByProject :many
SELECT id, project_id, name, description, color, enabled, file_uploads_enabled, redirect_url, created_at, updated_at, deleted_at
FROM "form"
WHERE project_id = $1
`

// ListFormsByProject fetches a list of Forms by their project ID
func (q *Queries) ListFormsByProject(ctx context.Context, projectID uuid.UUID) ([]Form, error) {
	rows, err := q.query(ctx, q.listFormsByProjectStmt, listFormsByProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Form
	for rows.Next() {
		var i Form
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.Name,
			&i.Description,
			&i.Color,
			&i.Enabled,
			&i.FileUploadsEnabled,
			&i.RedirectUrl,
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

const setFormColor = `-- name: SetFormColor :exec
UPDATE "form"
SET color = $2,
  updated_at = now()
WHERE id = $1
`

type SetFormColorParams struct {
	ID    string `db:"id" json:"id"`
	Color string `db:"color" json:"color"`
}

// SetFormColor sets the color field to the given value
func (q *Queries) SetFormColor(ctx context.Context, arg SetFormColorParams) error {
	_, err := q.exec(ctx, q.setFormColorStmt, setFormColor, arg.ID, arg.Color)
	return err
}

const setFormDescription = `-- name: SetFormDescription :exec
UPDATE "form"
SET description = $2,
  updated_at = now()
WHERE id = $1
`

type SetFormDescriptionParams struct {
	ID          string `db:"id" json:"id"`
	Description string `db:"description" json:"description"`
}

// SetFormDescription sets the description field to the given value
func (q *Queries) SetFormDescription(ctx context.Context, arg SetFormDescriptionParams) error {
	_, err := q.exec(ctx, q.setFormDescriptionStmt, setFormDescription, arg.ID, arg.Description)
	return err
}

const setFormName = `-- name: SetFormName :exec
UPDATE "form"
SET name = $2,
  updated_at = now()
WHERE id = $1
`

type SetFormNameParams struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// SetFormName sets the name field to the given value
func (q *Queries) SetFormName(ctx context.Context, arg SetFormNameParams) error {
	_, err := q.exec(ctx, q.setFormNameStmt, setFormName, arg.ID, arg.Name)
	return err
}

const setFormRedirectUrl = `-- name: SetFormRedirectUrl :exec
UPDATE "form"
SET redirect_url = $2,
  updated_at = now()
WHERE id = $1
`

type SetFormRedirectUrlParams struct {
	ID          string `db:"id" json:"id"`
	RedirectUrl string `db:"redirect_url" json:"redirect_url"`
}

// SetFormRedirectUrl sets the redirect_url field to the given value
func (q *Queries) SetFormRedirectUrl(ctx context.Context, arg SetFormRedirectUrlParams) error {
	_, err := q.exec(ctx, q.setFormRedirectUrlStmt, setFormRedirectUrl, arg.ID, arg.RedirectUrl)
	return err
}

const softDeleteForm = `-- name: SoftDeleteForm :exec
UPDATE "form"
SET deleted_at = now(),
  updated_at = now()
WHERE id = $1
`

// SoftDeleteForm sets the deleted_at timestamp to now(), indicating that the Form is deleted
func (q *Queries) SoftDeleteForm(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.softDeleteFormStmt, softDeleteForm, id)
	return err
}

const unsoftDeleteForm = `-- name: UnsoftDeleteForm :exec
UPDATE "form"
SET soft_deleted_at = NULL,
  updated_at = now()
WHERE id = $1
`

// UnsoftDeleteForm sets the deleted_at timestamp to NULL, indicating that the Form is not deleted
func (q *Queries) UnsoftDeleteForm(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.unsoftDeleteFormStmt, unsoftDeleteForm, id)
	return err
}

const updateForm = `-- name: UpdateForm :one
UPDATE "form"
SET name = $2,
  description = $3,
  color = $4,
  enabled = $5,
  file_uploads_enabled = $6,
  redirect_url = $7,
  updated_at = now()
WHERE id = $1
RETURNING id, project_id, name, description, color, enabled, file_uploads_enabled, redirect_url, created_at, updated_at, deleted_at
`

type UpdateFormParams struct {
	ID                 string `db:"id" json:"id"`
	Name               string `db:"name" json:"name"`
	Description        string `db:"description" json:"description"`
	Color              string `db:"color" json:"color"`
	Enabled            bool   `db:"enabled" json:"enabled"`
	FileUploadsEnabled bool   `db:"file_uploads_enabled" json:"file_uploads_enabled"`
	RedirectUrl        string `db:"redirect_url" json:"redirect_url"`
}

// UpdateForm updates the details of a Form
func (q *Queries) UpdateForm(ctx context.Context, arg UpdateFormParams) (Form, error) {
	row := q.queryRow(ctx, q.updateFormStmt, updateForm,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Color,
		arg.Enabled,
		arg.FileUploadsEnabled,
		arg.RedirectUrl,
	)
	var i Form
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Name,
		&i.Description,
		&i.Color,
		&i.Enabled,
		&i.FileUploadsEnabled,
		&i.RedirectUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
