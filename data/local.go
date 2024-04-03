package data

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"sync"

	"github.com/taylow/freeformed/form"
)

var _ (form.DataRepository) = (*localDataRepository)(nil)

// localDataRepository is a data repository backed by the local filesystem
type localDataRepository struct {
	config *LocalDataConfig

	mx sync.Mutex
}

// LocalData represents the local data structure
type LocalData struct {
	Form map[string]struct {
		Entry map[string]struct {
			Data map[string][]string `json:"data"`
		} `json:"entry"`
	} `json:"form,omitempty"`
}

// NewLocalDataRepository returns a new local data repository
func NewLocalDataRepository(config *LocalDataConfig) (*localDataRepository, error) {
	if config == nil {
		config = NewLocalDataConfig()
	}

	slog.Debug("creating local data repository", "dataFilePath", config.dataFilePath)
	if _, err := os.Stat(config.dataFilePath); os.IsNotExist(err) {
		if err := initFile(config.dataFilePath); err != nil {
			return nil, err
		}
	}

	slog.Debug("local data repository created")
	return &localDataRepository{
		config: config,
	}, nil
}

// Close closes the data repository
func (*localDataRepository) Close() error {
	return nil
}

// StoreData stores the form data to the local filesystem
func (r *localDataRepository) SaveData(ctx context.Context, formID, entryID string, data map[string][]string) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	slog.Debug("storing form data", "formID", formID, "entryID", entryID)

	// read existing data
	file, err := os.OpenFile(r.config.dataFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	existingData := make(map[string]map[string]map[string][]string)
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &existingData); err != nil {
		return err
	}

	// add new data
	if existingData[formID] == nil {
		existingData[formID] = make(map[string]map[string][]string)
	}
	existingData[formID][entryID] = data

	// write updated data
	file.Truncate(0)
	file.Seek(0, 0)

	enc := json.NewEncoder(file)
	if r.config.formatJSON {
		enc.SetIndent("", "  ")
	}
	if err := enc.Encode(existingData); err != nil {
		return err
	}

	return nil
}

// initFile creates a new file at the provided path and writes an empty JSON object to it
func initFile(filePath string) error {
	slog.Debug("initialising data file", "filePath", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	localData := LocalData{}
	enc := json.NewEncoder(file)
	if err := enc.Encode(localData); err != nil {
		return err
	}

	return nil
}
