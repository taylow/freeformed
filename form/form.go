package form

import (
	"context"
	"io"
	"net/http"

	db "github.com/taylow/freeformed/db/form"
)

// Closer is the interface for closing resources
type Closer interface {
	// Close closes the resource
	Close() error
}

// Handler handles form data requests and forwards it to the relevant places
type Handler interface {
	// HandleForm handles the raw request and delegates to the appropriate form handler
	HandleForm(w http.ResponseWriter, r *http.Request) error
	// HandleMultiPartForm handles the request as a multipart form
	HandleMultiPartForm(form db.Form, entryID string, w http.ResponseWriter, r *http.Request) error
	// HandleURLEncodedForm handles the request as a url encoded form
	HandleURLEncodedForm(form db.Form, entryID string, w http.ResponseWriter, r *http.Request) error
	// HandleJSONForm handles the request as a url encoded form
	HandleJSONForm(form db.Form, entryID string, w http.ResponseWriter, r *http.Request) error
}

// DataRepository is the interface for storing form data
type DataRepository interface {
	Closer

	// SaveData stores the provided form data
	SaveData(
		ctx context.Context,
		formID, entryID string,
		data map[string][]string,
	) error
}

// FileRepository is the interface for storing multipart-form files
type FileRepository interface {
	Closer

	// SaveFile saves the provided file to the repository with the given filename
	SaveFile(
		ctx context.Context,
		formID, entryID, fieldName, fileName string,
		file io.Reader,
	) error
}
