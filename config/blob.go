package config

import (
	"context"

	awsv2cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

const (
	// DefaultBucket is the default bucket for storing files
	DefaultBucket            = "uploaded"
	DefaultRegion            = "eu-west2"
	DefaultEndpoint          = "http://localhost:9000"
	DefaultDisableSSL        = true
	DefaultForceCreateBucket = false
	DefaultForceS3PathStyle  = true
)

// BlobFileConfigOption is a configuration option for a file repository
type BlobFileConfigOption func(*BlobFileConfig)

// BlobFileConfig is a configuration for a blob-backed file repository
type BlobFileConfig struct {
	DataFileName string

	Bucket            string
	Region            string
	Endpoint          string
	DisableSSL        bool
	ForceCreateBucket bool
	ForceS3PathStyle  bool

	useStaticCredentials    bool
	staticCredentialsKey    string
	staticCredentialsSecret string

	useEnvCredentials bool
}

// NewBlobFileConfig returns a new FileConfig with default values, unless overridden by the provided options
func NewBlobFileConfig(opts ...BlobFileConfigOption) *BlobFileConfig {
	c := &BlobFileConfig{
		DataFileName:      DefaultDataFileName,
		Bucket:            DefaultBucket,
		Region:            DefaultRegion,
		Endpoint:          DefaultEndpoint,
		DisableSSL:        DefaultDisableSSL,
		ForceCreateBucket: DefaultForceCreateBucket,
		ForceS3PathStyle:  DefaultForceS3PathStyle,
	}
	c.Apply(opts...)
	return c
}

// Apply applies the provided options to the config
func (c *BlobFileConfig) Apply(opts ...BlobFileConfigOption) {
	for _, opt := range opts {
		opt(c)
	}
}

// ToAWSConfig convers the config to an aws.Config
func (c *BlobFileConfig) ToAWSConfig() *aws.Config {
	// TODO support more credential types
	// TODO support v2
	creds := credentials.NewEnvCredentials()
	if c.useStaticCredentials {
		creds = credentials.NewStaticCredentials(c.staticCredentialsKey, c.staticCredentialsSecret, "")
	} else if c.useEnvCredentials {
		creds = credentials.NewEnvCredentials()
	}

	return &aws.Config{
		Credentials:      creds,
		Endpoint:         aws.String(c.Endpoint),
		DisableSSL:       aws.Bool(c.DisableSSL),
		S3ForcePathStyle: aws.Bool(c.ForceS3PathStyle),
		Region:           aws.String(c.Region),
	}
}

// ToAWSConfigV2 convers the config to an aws.Config
func (c *BlobFileConfig) ToAWSConfigV2(ctx context.Context) (*awsv2cfg.Config, error) {
	// TODO support v2
	// cfg, err := awsv2cfg.LoadDefaultConfig(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// return &cfg, nil
	panic("not implemented")
}

// WithBucket sets the root path for storing files
func WithBucket(bucket string) BlobFileConfigOption {
	return func(c *BlobFileConfig) {
		c.Bucket = bucket
	}
}

// WithRegion sets the region for the file repository
func WithRegion(region string) BlobFileConfigOption {
	return func(c *BlobFileConfig) {
		c.Region = region
	}
}

// WithEndpoint sets the endpoint for the file repository
func WithEndpoint(endpoint string) BlobFileConfigOption {
	return func(c *BlobFileConfig) {
		c.Endpoint = endpoint
	}
}

// WithDisableSSL sets whether to disable SSL for the file repository
func WithDisableSSL(disableSSL bool) BlobFileConfigOption {
	return func(c *BlobFileConfig) {
		c.DisableSSL = disableSSL
	}
}

// WithForceCreateBucket sets whether to force create the bucket for the file repository
func WithForceCreateBucket(forceCreateBucket bool) BlobFileConfigOption {
	panic("force create bucket not implemented")
	return func(c *BlobFileConfig) {
		c.ForceCreateBucket = forceCreateBucket
	}
}

// WithForceS3PathStyle sets whether to force S3 path style for the file repository
func WithForceS3PathStyle(forceS3PathStyle bool) BlobFileConfigOption {
	return func(c *BlobFileConfig) {
		c.ForceS3PathStyle = forceS3PathStyle
	}
}

// WithStaticCredentials sets the static credentials for the file repository
func WithStaticCredentials(key, secret string) BlobFileConfigOption {
	return func(c *BlobFileConfig) {
		c.useStaticCredentials = true
		c.staticCredentialsKey = key
		c.staticCredentialsSecret = secret
	}
}

// WithDataFileName sets the filename for storing data
func WithBlobDataFileName(dataFileName string) BlobFileConfigOption {
	return func(c *BlobFileConfig) {
		c.DataFileName = dataFileName
	}
}
