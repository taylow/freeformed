package data

const (
	DefaultLocalDataFilePath = "data.json"
	DefaultFormatJSON        = false
)

type LocalDataConfigOptions func(*LocalDataConfig)

type LocalDataConfig struct {
	dataFilePath string
	formatJSON   bool
}

func NewLocalDataConfig(opts ...LocalDataConfigOptions) *LocalDataConfig {
	c := &LocalDataConfig{
		dataFilePath: DefaultLocalDataFilePath,
		formatJSON:   DefaultFormatJSON,
	}
	c.Apply(opts...)
	return c
}

func (c *LocalDataConfig) Apply(opts ...LocalDataConfigOptions) {
	for _, opt := range opts {
		opt(c)
	}
}

func WithDataFilePath(dataFilePath string) LocalDataConfigOptions {
	return func(c *LocalDataConfig) {
		c.dataFilePath = dataFilePath
	}
}

func WithFormatJSON(formatJSON bool) LocalDataConfigOptions {
	return func(c *LocalDataConfig) {
		c.formatJSON = formatJSON
	}
}
