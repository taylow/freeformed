package config

const (
	// DefaultRootPath is the default root directory for storing files
	DefaultRootPath = "uploaded"
	// DefaultDataFileName is the default filename for storing data
	DefaultDataFileName = "data.json"
)

// LocalFileConfigOption is a configuration option for a file repository
type LocalFileConfigOption func(*LocalFileConfig)

// LocalFileConfig is a configuration for a file repository
type LocalFileConfig struct {
	RootPath     string
	DataFileName string
}

// NewLocalFileConfig returns a new FileConfig with default values, unless overridden by the provided options
func NewLocalFileConfig(opts ...LocalFileConfigOption) *LocalFileConfig {
	c := &LocalFileConfig{
		RootPath:     DefaultRootPath,
		DataFileName: DefaultDataFileName,
	}
	c.Apply(opts...)
	return c
}

// Apply applies the provided options to the config
func (c *LocalFileConfig) Apply(opts ...LocalFileConfigOption) {
	for _, opt := range opts {
		opt(c)
	}
}

// WithRootPath sets the root path for storing files
func WithRootPath(rootPath string) LocalFileConfigOption {
	return func(c *LocalFileConfig) {
		c.RootPath = rootPath
	}
}

// WithDataFileName sets the filename for storing data
func WithDataFileName(dataFileName string) LocalFileConfigOption {
	return func(c *LocalFileConfig) {
		c.DataFileName = dataFileName
	}
}
