package gcs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"time"

	"cloud.google.com/go/storage"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oasdiff/go-common/env"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"gopkg.in/yaml.v3"
)

type Client interface {
	UploadSpec(tenant string, webhook string, name string, spec *openapi3.T) error
	Upload(path string, file []byte) error
	Read(path string) ([]byte, error)
	Close() error
}

type Store struct {
	client *storage.Client
	bucket string
}

func NewStore() Client {

	if key := env.GetGCPStorageKey(); key != "" {
		conf, err := google.JWTConfigFromJSON([]byte(key), storage.ScopeFullControl)
		if err != nil {
			slog.Error("failed to config storage JWT from JSON key", "error", err)
			return nil
		}
		ctx := context.Background()
		opt := []option.ClientOption{option.WithTokenSource(conf.TokenSource(ctx))}

		client, err := storage.NewClient(ctx, opt...)
		if err != nil {
			slog.Error("failed to create datastore client", "error", err)
			return nil
		}

		return &Store{client: client, bucket: env.GetGCPStorageBucket()}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		slog.Error("failed to create storage client", "error", err)
		return nil
	}

	return &Store{client: client, bucket: env.GetGCPStorageBucket()}
}

func (store *Store) UploadSpec(tenant string, webhook string, name string, spec *openapi3.T) error {

	payload, err := yaml.Marshal(spec)
	if err != nil {
		slog.Error("failed to marshal OpenAPI spec", "error", err, "tenant", tenant, "webhook", webhook)
		return err
	}

	err = store.Upload(GetSpecPath(tenant, webhook, name), payload)
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) Upload(path string, file []byte) error {

	w := store.client.Bucket(store.bucket).Object(path).NewWriter(context.Background())
	defer func() {
		if err := w.Close(); err != nil {
			slog.Error("failed to close gcs bucket writer",
				"error", err, "bucket", store.bucket, "file", path)
		}
	}()

	if _, err := w.Write(file); err != nil {
		slog.Error("failed to create file in GCS bucket",
			"error", err, "bucket", store.bucket, "file", path)
		return err
	}

	return nil
}

func (store *Store) Read(path string) ([]byte, error) {

	rc, err := store.client.Bucket(store.bucket).
		Object(path).
		NewReader(context.Background())
	if err != nil {
		msg := fmt.Sprintf("failed to create reader for file '%s' with '%v'",
			path, err)
		slog.Error(msg)
		return nil, errors.New(msg)
	}
	defer func() {
		if err := rc.Close(); err != nil {
			slog.Error("failed to close reader", "error", err)
		}
	}()

	data, err := io.ReadAll(rc)
	if err != nil {
		msg := fmt.Sprintf("failed to read file '%s' with '%v'", path, err)
		slog.Error(msg)
		return nil, errors.New(msg)
	}

	return data, nil
}

func (store *Store) Close() error {

	return store.client.Close()
}

// Buckets/syncc/[]{tenant-id}/spec/[]{webhook-id}/[]spec
func GetSpecPath(tenant, webhook, name string) string {

	return fmt.Sprintf("%s/spec/%s/%s", tenant, webhook, name)
}
