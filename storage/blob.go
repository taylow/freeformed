package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/taylow/freeformed/config"
	"github.com/taylow/freeformed/form"

	// "gocloud.dev/blob/s3blob"
	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
	_ "gocloud.dev/blob/s3blob"
)

var _ (form.FileRepository) = (*blobFormRepository)(nil)
var _ (form.DataRepository) = (*blobFormRepository)(nil)

// blobFormRepository is a file repository backed by the blob filesystem
type blobFormRepository struct {
	config *config.BlobFileConfig

	bucket *blob.Bucket
}

// NewBlobFileRepository returns a new blob file repository
func NewBlobFileRepository(ctx context.Context, cfg *config.BlobFileConfig) (*blobFormRepository, error) {
	if cfg == nil {
		return nil, fmt.Errorf("blob repository config is required")
	}

	sess, err := session.NewSession(cfg.ToAWSConfig())
	if err != nil {
		return nil, err
	}

	bucket, err := s3blob.OpenBucket(ctx, sess, cfg.Bucket, nil)
	if err != nil {
		return nil, err
	}

	// TODO replace with v2
	// see https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/ for more info.
	// cfg, err := awsv2cfg.LoadDefaultConfig(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// clientV2 := s3v2.NewFromConfig(config.ToAWSConfig())
	// bucket, err := s3blob.OpenBucketV2(ctx, clientV2, "my-bucket", nil)
	// if err != nil {
	// 	return nil, err
	// }

	return &blobFormRepository{
		config: cfg,
		bucket: bucket,
	}, nil
}

// Close closes the file repository
func (r *blobFormRepository) Close() error {
	return r.bucket.Close()
}

// SaveData saves the provided data file to the repository with the given filename to a bucket
func (r *blobFormRepository) SaveData(
	ctx context.Context,
	formID, entryID string,
	data map[string][]string,
) error {
	key := fmt.Sprintf("%s/%s/%s", formID, entryID, r.config.DataFileName)

	opts := &blob.WriterOptions{
		ContentType: "application/json",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := r.bucket.WriteAll(ctx, key, jsonData, opts); err != nil {
		return err
	}

	slog.Debug("file saved to blob repository", "key", key)

	return nil
}

// SaveFile saves the provided file to the repository with the given filename to a bucket
func (r *blobFormRepository) SaveFile(
	ctx context.Context,
	formID, entryID, fieldName, filename string,
	file io.Reader,
) error {
	key := fmt.Sprintf("%s/%s/%s/%s", formID, entryID, fieldName, filename)

	opts := &blob.WriterOptions{
		ContentType: "application/octet-stream",
	}

	if err := r.bucket.Upload(ctx, key, file, opts); err != nil {
		return err
	}

	slog.Debug("file saved to blob repository", "key", key)

	return nil
}
