package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/taylow/freeformed/config"
	"github.com/taylow/freeformed/form"
)

var _ (form.FileRepository) = (*localFormRepository)(nil)
var _ (form.DataRepository) = (*localFormRepository)(nil)

// localFormRepository is a file repository backed by the local filesystem
type localFormRepository struct {
	config *config.LocalFileConfig
}

// NewLocalFileRepository returns a new local file repository
func NewLocalFileRepository(cfg *config.LocalFileConfig) (*localFormRepository, error) {
	if cfg == nil {
		cfg = config.NewLocalFileConfig()
	}

	if _, err := os.Stat(cfg.RootPath); os.IsNotExist(err) {
		if err := os.MkdirAll(cfg.RootPath, 0755); err != nil {
			return nil, err
		}
	}

	return &localFormRepository{
		config: cfg,
	}, nil
}

// Close closes the file repository
func (r *localFormRepository) Close() error {
	// TODO close the open file (if we ever keep it open)
	return nil
}

// SaveData saves the provided file to the repository with the given filename to disk
func (r *localFormRepository) SaveData(
	ctx context.Context,
	formID, entryID string,
	data map[string][]string) error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s/%s", r.config.RootPath, formID, entryID)); os.IsNotExist(err) {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s/%s", r.config.RootPath, formID, entryID), 0755); err != nil {
			return err
		}
	}
	out, err := os.Create(fmt.Sprintf("%s/%s/%s/%s", r.config.RootPath, formID, entryID, r.config.DataFileName))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = fmt.Fprintf(out, "%v", data)
	return err
}

// SaveFile saves the provided file to the repository with the given filename to disk
func (r *localFormRepository) SaveFile(
	ctx context.Context,
	formID, entryID, fieldName, fileName string,
	file io.Reader,
) error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s/%s/%s", r.config.RootPath, formID, entryID, fieldName)); os.IsNotExist(err) {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s/%s/%s", r.config.RootPath, formID, entryID, fieldName), 0755); err != nil {
			return err
		}
	}
	out, err := os.Create(fmt.Sprintf("%s/%s/%s/%s/%s", r.config.RootPath, formID, entryID, fieldName, fileName))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}
