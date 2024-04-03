package storage

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	db "github.com/taylow/freeformed/db/form"
	"github.com/taylow/freeformed/form"
)

var _ (form.DataRepository) = (*postgresDataRepository)(nil)

// postgresDataRepository is a data repository backed by a postgres database
type postgresDataRepository struct {
	db db.Querier
}

// NewPostgresDataRepository returns a new postgres data repository
func NewPostgresDataRepository(db db.Querier) *postgresDataRepository {
	return &postgresDataRepository{
		db: db,
	}
}

// Close implements form.DataRepository.
func (r *postgresDataRepository) Close() error {
	panic("unimplemented")
}

// SaveData implements form.DataRepository.
func (r *postgresDataRepository) SaveData(ctx context.Context, formID string, entryID string, data map[string][]string) error {
	id, err := uuid.Parse(entryID)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = r.db.CreateFormData(ctx, db.CreateFormDataParams{
		ID:     id,
		FormID: formID,
		Data:   jsonData,
	})

	return err
}
