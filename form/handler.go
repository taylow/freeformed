package form

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	// "github.com/pkg/errors"
	db "github.com/taylow/freeformed/db/form"
	"github.com/taylow/freeformed/errors"
)

var _ (Handler) = (*handler)(nil)

const (
	PlaceholderProjectID = "projectID"
	PlaceholderEntryID   = "entryID"
)

// handler handles form data requests and forwards it to the relevant places
type handler struct {
	config *HandlerConfig

	formRepository db.Querier
	dataRepository DataRepository
	fileRepository FileRepository
}

// NewHandler creates a new instance of the processor
func NewHandler(config *HandlerConfig, formRepository db.Querier, dataRepository DataRepository, fileRepository FileRepository) Handler {
	return &handler{
		config:         config,
		formRepository: formRepository,
		dataRepository: dataRepository,
		fileRepository: fileRepository,
	}
}

// HandleForm handles the raw request and delegates to the appropriate form processor
func (p *handler) HandleForm(w http.ResponseWriter, r *http.Request) error {
	formID := r.PathValue("id")
	if formID == "" {
		slog.Error("Form ID not provided")
		return errors.ErrFormIDNotProvided
	}

	form, err := p.formRepository.Form(r.Context(), formID)
	if err != nil {
		slog.Error("Failed to read form", "error", err)
		return errors.FromErr(err).WithCode(http.StatusInternalServerError).WithMessage("failed to read form")
	}

	if !form.Enabled {
		slog.Error("Form is disabled")
		return errors.ErrFormDisabled
	}

	if form.DeletedAt.Valid {
		slog.Error("Form is deleted")
		return errors.ErrFormDeleted
	}

	entryID := uuid.New().String()

	if r.Method != http.MethodPost {
		slog.Error("Invalid request method", "method", r.Method)
		return errors.ErrInvalidRequestMethod
	}

	if r.ContentLength > p.config.MaxRequestSize && p.config.MaxRequestSize > 0 {
		slog.Error("Request size exceeds max request size", "requestSize", r.ContentLength, "maxRequestSize", p.config.MaxRequestSize)
		return errors.ErrMaxRequestSizeExceeded
	}

	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/x-www-form-urlencoded") {
		return p.HandleURLEncodedForm(form, entryID, w, r)
	} else if strings.Contains(contentType, "multipart/form-data") {
		return p.HandleMultiPartForm(form, entryID, w, r)
	} else if strings.Contains(contentType, "application/json") {
		return p.HandleJSONForm(form, entryID, w, r)
	} else {
		return errors.ErrInvalidContentType.WithMessage(fmt.Sprintf("invalid content type + %q", contentType))
	}
}

// HandleMultiPartForm handles a multipart form request, saves the form content and stores any files to the file repository
func (h *handler) HandleMultiPartForm(form db.Form, entryID string, w http.ResponseWriter, r *http.Request) error {
	// TODO look into remainder of this function
	err := r.ParseMultipartForm(h.config.MaxMemory)
	if err != nil {
		slog.Error("Failed to parse multipart form", "error", err)
		return errors.ErrFailedToParseMultipartForm.WithCause(err)
	}

	formData := r.MultipartForm.Value
	// TODO decide if we should truncate or error out
	// if truncating, move this code to a counter in the loop of formData
	if len(formData) > h.config.MaxFields && h.config.MaxFields > 0 {
		slog.Error("Max fields exceeded", "maxFields", h.config.MaxFields)
		return errors.ErrLimitsExceeded.WithMessage(fmt.Sprintf("max fields exceeded (%d)", h.config.MaxFields))
	}

	fileCount := 0
	for field := range r.MultipartForm.File {
		formData[field] = []string{}
		fileHeaders := r.MultipartForm.File[field]

		// check for duplicate filenames
		for i := 0; i < len(fileHeaders); i++ {
			for j := i + 1; j < len(fileHeaders); j++ {
				if fileHeaders[i].Filename == fileHeaders[j].Filename {
					slog.Error("Duplicate filename", "filename", fileHeaders[i].Filename)
					http.Error(w, fmt.Sprintf("Duplicate filename (%s)", fileHeaders[i].Filename), http.StatusConflict)
					return errors.ErrDuplicateFilename.WithMessage(fmt.Sprintf("duplicate filename in same field (%s)", fileHeaders[i].Filename))
				}
			}
		}

		// store files and add their info to the form data
		for _, fileHeader := range fileHeaders {
			if fileHeader.Size > h.config.MaxFileSize && h.config.MaxFileSize > 0 {
				slog.Error("File exceeds max file size", "filename", fileHeader.Filename, "size", fileHeader.Size, "maxFileSize", h.config.MaxFileSize)
				return errors.ErrLimitsExceeded.WithMessage(fmt.Sprintf("file exceeds max file size (%d)", h.config.MaxFileSize))
			}
			if fileCount >= h.config.MaxFiles && h.config.MaxFiles > 0 {
				slog.Error("Max files exceeded", "maxFiles", h.config.MaxFiles)
				return errors.ErrLimitsExceeded.WithMessage(fmt.Sprintf("max files exceeded (%d)", h.config.MaxFiles))
			}

			file, err := fileHeader.Open()
			if err != nil {
				slog.Error("Failed to open file", "error", err)
				return errors.FromErr(err).WithCode(http.StatusInternalServerError).WithMessage("failed to open file")
			}
			if err := h.fileRepository.SaveFile(r.Context(), form.ID, entryID, field, fileHeader.Filename, file); err != nil {
				slog.Error("Failed to save file", "error", err)
				return errors.FromErr(err).WithCode(http.StatusInternalServerError).WithMessage("failed to save file")
			}
			err = file.Close()
			if err != nil {
				slog.Error("Failed to close file", "error", err)
				return errors.FromErr(err).WithCode(http.StatusInternalServerError).WithMessage("failed to close file")
			}

			formData[field] = append(formData[field], fmt.Sprintf("%s/%s/%s/%s", form.ID, entryID, field, fileHeader.Filename))
			// TODO decide on what data to send here and how to format it
			// ideally we would send filename, path, size, and any other relevant data
			fileCount++
		}
	}

	if err := validateLengths(formData, h.config); err != nil {
		slog.Error("Failed to validate form data", "error", err)
		return err
	}

	formData["_id"] = []string{entryID}
	formData["_date"] = []string{time.Now().Format(time.RFC3339)}
	formData["_status"] = []string{"pending"}

	if err := h.dataRepository.SaveData(r.Context(), form.ID, entryID, formData); err != nil {
		slog.Error("Failed to store form data", "error", err)
		return errors.FromErr(err).WithCode(http.StatusInternalServerError).WithMessage("failed to store form data")
	}

	// TODO fire event

	fmt.Fprintln(w, "Form data stored successfully")

	return nil
}

// HandleURLEncodedForm handles a url-encoded form request and saves the form content
func (h *handler) HandleURLEncodedForm(form db.Form, entryID string, w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		slog.Error("Failed to parse form", "error", err)
		return errors.ErrFailedToParseForm.WithCause(err)
	}

	formData := r.Form
	if err := validateLengths(formData, h.config); err != nil {
		slog.Error("Failed to validate form data", "error", err)
		return err
	}

	formData["_id"] = []string{entryID}
	formData["_date"] = []string{time.Now().Format(time.RFC3339)}
	formData["_status"] = []string{"pending"}

	if err := h.dataRepository.SaveData(r.Context(), form.ID, entryID, formData); err != nil {
		slog.Error("Failed to store form data", "error", err)
		return errors.FromErr(err).WithCode(http.StatusInternalServerError).WithMessage("failed to store form data")
	}

	// TODO fire event

	fmt.Fprintln(w, "Form data stored successfully")

	return nil
}

// HandleJSONForm handles a json-encoded form request and saves the content
func (h *handler) HandleJSONForm(form db.Form, entryID string, w http.ResponseWriter, r *http.Request) error {
	var data map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		return errors.FromErr(err).WithCode(http.StatusBadRequest).WithMessage("Failed to decode json")
	}

	formData := r.Form
	if err := validateLengths(formData, h.config); err != nil {
		slog.Error("Failed to validate form data", "error", err)
		return err
	}

	formData["_id"] = []string{entryID}
	formData["_date"] = []string{time.Now().Format(time.RFC3339)}
	formData["_status"] = []string{"pending"}

	if err := h.dataRepository.SaveData(r.Context(), form.ID, entryID, formData); err != nil {
		slog.Error("Failed to store form data", "error", err)
		return errors.FromErr(err).WithCode(http.StatusInternalServerError).WithMessage("failed to store form data")
	}

	// TODO fire event

	fmt.Fprintln(w, "Form data stored successfully")

	return nil
}

// validateLengths validates the lengths of the form data
func validateLengths(formData map[string][]string, config *HandlerConfig) error {
	valueCount := 0
	for field, values := range formData {
		if len(values) > config.MaxValuesPerField && config.MaxValuesPerField > 0 {
			return errors.ErrLimitsExceeded.WithMessage(fmt.Sprintf("values per field exceeds max values per field (%d)", len(values)))
		}

		valueCount += len(values)
		if valueCount > config.MaxValues && config.MaxValues > 0 {
			return errors.ErrLimitsExceeded.WithMessage(fmt.Sprintf("value count exceeds max values: %d", valueCount))
		}

		if len(field) > config.MaxFieldNameLen && config.MaxFieldNameLen > 0 {
			return errors.ErrLimitsExceeded.WithMessage(fmt.Sprintf("field name length exceeds max field name length (%d)", len(field)))
		}

		for _, value := range values {
			if int64(len(value)) > config.MaxValueLen && config.MaxValueLen > 0 {
				return errors.ErrLimitsExceeded.WithMessage(fmt.Sprintf("value length exceeds max value length (%d)", len(value)))
			}
		}
	}

	return nil
}

// // saveFormData saves the form data to the data repository
// func (p *processor) saveFormData(ctx context.Context, projectID, entryID string, formData map[string][]string) error {
// 	var buf bytes.Buffer
// 	enc := json.NewEncoder(&buf)
// 	if p.config.FormatJSON {
// 		enc.SetIndent("", "  ")
// 	}
// 	if err := enc.Encode(formData); err != nil {
// 		return err
// 	}

// 	if err := p.dataRepository.SaveData(ctx, projectID, entryID, &buf); err != nil {
// 		return err
// 	}

// 	return nil
// }
