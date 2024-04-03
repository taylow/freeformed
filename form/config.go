package form

// TODO allow for -1 to mean no limit and 0 to mean disabled

const (
	// DefaultFormatJSON is the default format for form data (false = minified, true = pretty-printed JSON)
	DefaultFormatJSON = true

	// DefaultMaxMemory is the maximum memory used for parsing a form
	DefaultMaxMemory = 32 << 20 // 32 MB
	// DefaultMaxFileSize is the maximum file size for a form file
	DefaultMaxFileSize = 32 << 20 // 32 MB
	// DefaultMaxFiles is the maximum number of files for a form
	DefaultMaxFiles = 3
	// DefaultMaxFields is the maximum number of fields for a form
	DefaultMaxFields = 3
	// DefaultMaxFieldNameLen is the maximum length of a form field name
	DefaultMaxFieldNameLen = 128
	// DefaultMaxValues is the maximum number of values for the entire form
	DefaultMaxValues = 10
	// DefaultMaxValuesPerField is the maximum number of values for a form field
	DefaultMaxValuesPerField = 3
	// DefaultMaxValueLen is the maximum length of a form value
	DefaultMaxValueLen = 1024
	// DefaultMaxRequestSize is the maximum size of a form request
	DefaultMaxRequestSize = 32 << 20 // 32 MB
)

type ConfigOption func(*HandlerConfig)

type HandlerConfig struct {
	FormatJSON bool

	MaxMemory         int64
	MaxFileSize       int64
	MaxFiles          int
	MaxFields         int
	MaxFieldNameLen   int
	MaxValues         int
	MaxValuesPerField int
	MaxValueLen       int64
	MaxRequestSize    int64

	StaticFormID string
}

// NewProcessorConfig returns a new ProcessorConfig with default values, unless overridden by the provided options.
func NewProcessorConfig(opts ...ConfigOption) *HandlerConfig {
	c := &HandlerConfig{
		MaxMemory:         DefaultMaxMemory,
		MaxFileSize:       DefaultMaxFileSize,
		MaxFiles:          DefaultMaxFiles,
		MaxFields:         DefaultMaxFields,
		MaxFieldNameLen:   DefaultMaxFieldNameLen,
		MaxValues:         DefaultMaxValues,
		MaxValuesPerField: DefaultMaxValuesPerField,
		MaxValueLen:       DefaultMaxValueLen,
		MaxRequestSize:    DefaultMaxRequestSize,
	}
	c.Apply(opts...)
	return c
}

func (c *HandlerConfig) SetMaxMemory(maxMemory int64) *HandlerConfig {
	c.MaxMemory = maxMemory
	return c
}

func (c *HandlerConfig) Apply(opts ...ConfigOption) {
	for _, opt := range opts {
		opt(c)
	}
}

// WithMaxMemory sets the maximum memory used for parsing a form
func WithMaxMemory(maxMemory int64) ConfigOption {
	return func(c *HandlerConfig) {
		c.MaxMemory = maxMemory
	}
}

// WithMaxFileSize sets the maximum file size for a form file
func WithMaxFileSize(maxFileSize int64) ConfigOption {
	return func(c *HandlerConfig) {
		c.MaxFileSize = maxFileSize
	}
}

// WithMaxFiles sets the maximum number of files for a form
func WithMaxFiles(maxFiles int) ConfigOption {
	return func(c *HandlerConfig) {
		c.MaxFiles = maxFiles
	}
}

// WithMaxFields sets the maximum number of fields for a form
func WithMaxFields(maxFields int) ConfigOption {
	return func(c *HandlerConfig) {
		c.MaxFields = maxFields
	}
}

// WithMaxFieldNameLen sets the maximum length of a form field name
func WithMaxFieldNameLen(maxFieldNameLen int) ConfigOption {
	return func(c *HandlerConfig) {
		c.MaxFieldNameLen = maxFieldNameLen
	}
}

// WithMaxValues sets the maximum number of values for the entire form
func WithMaxValues(maxValues int) ConfigOption {
	return func(c *HandlerConfig) {
		c.MaxValues = maxValues
	}
}

// WithMaxValuesPerField sets the maximum length of a form value
func WithMaxValuesPerField(maxValuesPerField int) ConfigOption {
	return func(c *HandlerConfig) {
		c.MaxValuesPerField = maxValuesPerField
	}
}

// WithMaxValueLen sets the maximum length of a form value
func WithMaxValueLen(maxValueLen int64) ConfigOption {
	return func(c *HandlerConfig) {
		c.MaxValueLen = maxValueLen
	}
}

// WithMaxRequestSize sets the maximum size of a form request
func WithMaxRequestSize(maxRequestSize int64) ConfigOption {
	return func(c *HandlerConfig) {
		c.MaxRequestSize = maxRequestSize
	}
}

// WithStaticFormID sets the static form ID for the form handler
func WithStaticFormID(staticFormID string) ConfigOption {
	return func(c *HandlerConfig) {
		c.StaticFormID = staticFormID
	}
}
