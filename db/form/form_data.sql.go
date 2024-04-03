// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: form_data.sql

package postgres

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

const createFormData = `-- name: CreateFormData :one
INSERT INTO "form_data" (
    id,
    form_id,
    status,
    data
)
VALUES ($1, $2, $3, $4)
RETURNING id, form_id, status, data, created_at, updated_at, deleted_at
`

type CreateFormDataParams struct {
	ID     uuid.UUID       `db:"id" json:"id"`
	FormID string          `db:"form_id" json:"form_id"`
	Status string          `db:"status" json:"status"`
	Data   json.RawMessage `db:"data" json:"data"`
}

// CreateFormData creates a new Form entry
func (q *Queries) CreateFormData(ctx context.Context, arg CreateFormDataParams) (FormDatum, error) {
	row := q.queryRow(ctx, q.createFormDataStmt, createFormData,
		arg.ID,
		arg.FormID,
		arg.Status,
		arg.Data,
	)
	var i FormDatum
	err := row.Scan(
		&i.ID,
		&i.FormID,
		&i.Status,
		&i.Data,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteFormData = `-- name: DeleteFormData :one
DELETE FROM "form_data"
WHERE id = $1
RETURNING id, form_id, status, data, created_at, updated_at, deleted_at
`

// DeleteFormData deletes a FormData entry
func (q *Queries) DeleteFormData(ctx context.Context, id uuid.UUID) (FormDatum, error) {
	row := q.queryRow(ctx, q.deleteFormDataStmt, deleteFormData, id)
	var i FormDatum
	err := row.Scan(
		&i.ID,
		&i.FormID,
		&i.Status,
		&i.Data,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listFormDataByFormID = `-- name: ListFormDataByFormID :many
SELECT id, form_id, status, data, created_at, updated_at, deleted_at
FROM "form_data"
WHERE form_id = $1
`

// ListFormDataByFormID fetches a list of FormData by their form ID
func (q *Queries) ListFormDataByFormID(ctx context.Context, formID string) ([]FormDatum, error) {
	rows, err := q.query(ctx, q.listFormDataByFormIDStmt, listFormDataByFormID, formID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FormDatum
	for rows.Next() {
		var i FormDatum
		if err := rows.Scan(
			&i.ID,
			&i.FormID,
			&i.Status,
			&i.Data,
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

const readFormData = `-- name: ReadFormData :one
SELECT id, form_id, status, data, created_at, updated_at, deleted_at
FROM "form_data"
WHERE id = $1
LIMIT 1
`

// ReadFormData fetches a single FormData by ID
func (q *Queries) ReadFormData(ctx context.Context, id uuid.UUID) (FormDatum, error) {
	row := q.queryRow(ctx, q.readFormDataStmt, readFormData, id)
	var i FormDatum
	err := row.Scan(
		&i.ID,
		&i.FormID,
		&i.Status,
		&i.Data,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateStatus = `-- name: UpdateStatus :one
UPDATE "form_data"
SET status = $2,
  updated_at = now()
WHERE id = $1
RETURNING id, form_id, status, data, created_at, updated_at, deleted_at
`

type UpdateStatusParams struct {
	ID     uuid.UUID `db:"id" json:"id"`
	Status string    `db:"status" json:"status"`
}

// UpdateStatus updates the status of a Form
func (q *Queries) UpdateStatus(ctx context.Context, arg UpdateStatusParams) (FormDatum, error) {
	row := q.queryRow(ctx, q.updateStatusStmt, updateStatus, arg.ID, arg.Status)
	var i FormDatum
	err := row.Scan(
		&i.ID,
		&i.FormID,
		&i.Status,
		&i.Data,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
