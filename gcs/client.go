package gcs

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"github.com/oasdiff/go-common/env"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type Client interface {
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
			logrus.Fatalf("failed to config storage JWT from JSON key with '%v'", err)
		}
		ctx := context.Background()
		opt := []option.ClientOption{option.WithTokenSource(conf.TokenSource(ctx))}

		client, err := storage.NewClient(ctx, opt...)
		if err != nil {
			logrus.Fatalf("failed to create datastore client with '%v'", err)
		}

		return &Store{client: client, bucket: env.GetGCPStorageBucket()}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		logrus.Fatalf("failed to create storage client with '%v'", err)
	}

	return &Store{client: client, bucket: env.GetGCPStorageBucket()}
}

func (store *Store) Upload(path string, file []byte) error {

	w := store.client.Bucket(store.bucket).Object(path).NewWriter(context.Background())
	defer func() {
		if err := w.Close(); err != nil {
			logrus.Errorf("failed to close gcs bucket '%s' writer file '%s' with '%v'",
				store.bucket, path, err)
		}
	}()

	if _, err := w.Write(file); err != nil {
		logrus.Errorf("failed to create file in GCS bucket '%s' file '%s' with '%v'",
			store.bucket, path, err)
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
		logrus.Error(msg)
		return nil, errors.New(msg)
	}
	defer func() {
		if err := rc.Close(); err != nil {
			logrus.Errorf("failed to close reader with '%v'", err)
		}
	}()

	data, err := io.ReadAll(rc)
	if err != nil {
		msg := fmt.Sprintf("failed to read file '%s' with '%v'", path, err)
		logrus.Errorf(msg)
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
