// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package postgres

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Form struct {
	ID                 string       `db:"id" json:"id"`
	ProjectID          uuid.UUID    `db:"project_id" json:"project_id"`
	Name               string       `db:"name" json:"name"`
	Description        string       `db:"description" json:"description"`
	Color              string       `db:"color" json:"color"`
	Enabled            bool         `db:"enabled" json:"enabled"`
	FileUploadsEnabled bool         `db:"file_uploads_enabled" json:"file_uploads_enabled"`
	RedirectUrl        string       `db:"redirect_url" json:"redirect_url"`
	CreatedAt          time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt          sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

type FormDatum struct {
	ID        uuid.UUID       `db:"id" json:"id"`
	FormID    string          `db:"form_id" json:"form_id"`
	Status    string          `db:"status" json:"status"`
	Data      json.RawMessage `db:"data" json:"data"`
	CreatedAt time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt time.Time       `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime    `db:"deleted_at" json:"deleted_at"`
}

type Project struct {
	ID          uuid.UUID    `db:"id" json:"id"`
	TeamID      uuid.UUID    `db:"team_id" json:"team_id"`
	Name        string       `db:"name" json:"name"`
	Description string       `db:"description" json:"description"`
	Enabled     bool         `db:"enabled" json:"enabled"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

type Team struct {
	ID          uuid.UUID    `db:"id" json:"id"`
	Name        string       `db:"name" json:"name"`
	Description string       `db:"description" json:"description"`
	Enabled     bool         `db:"enabled" json:"enabled"`
	OwnerID     uuid.UUID    `db:"owner_id" json:"owner_id"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

type User struct {
	ID        uuid.UUID    `db:"id" json:"id"`
	FirstName string       `db:"first_name" json:"first_name"`
	LastName  string       `db:"last_name" json:"last_name"`
	Email     string       `db:"email" json:"email"`
	Enabled   bool         `db:"enabled" json:"enabled"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`
}