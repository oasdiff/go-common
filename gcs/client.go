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
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Client interface {
	UploadSpec(tenantId string, name string, file []byte) error
	ReadLatestSpec(tenantId string) ([]byte, error)
	Close() error
}

type Store struct {
	client *storage.Client
	bucket string
}

func NewStore() Client {

	if key := env.GetStorageKey(); key != "" {
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

		return &Store{client: client, bucket: env.GetBucket()}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		logrus.Fatalf("failed to create storage client with '%v'", err)
	}

	return &Store{client: client, bucket: env.GetBucket()}
}

// Buckets/syncc/{tenant-id}/spec/[]spec
func (store *Store) UploadSpec(tenantId string, name string, file []byte) error {

	w := store.client.Bucket(store.bucket).
		Object(fmt.Sprintf("%s/spec/%s", tenantId, name)).
		NewWriter(context.Background())
	defer func() {
		if err := w.Close(); err != nil {
			logrus.Errorf("failed to close gcs bucket '%s' writer file '%s' with '%v'",
				store.bucket, name, err)
		}
	}()

	if _, err := w.Write(file); err != nil {
		logrus.Errorf("failed to create file in GCS bucket '%s' file '%s' with '%v'",
			store.bucket, name, err)
		return err
	}

	return nil
}

func (store *Store) ReadLatestSpec(tenantId string) ([]byte, error) {

	path, err := store.getLatestSpec(tenantId)
	if err != nil {
		return nil, err
	}
	if path == "" {
		logrus.Infof("no specs for tenant '%s'", tenantId)
		return []byte{}, nil
	}

	rc, err := store.client.Bucket(store.bucket).
		Object(path).
		NewReader(context.Background())
	if err != nil {
		msg := fmt.Sprintf("failed to create reader for file '%s' with '%v' tenant '%s'",
			path, err, tenantId)
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
		msg := fmt.Sprintf("failed to read file '%s' with '%v' tenant '%s'", path, err, tenantId)
		logrus.Errorf(msg)
		return nil, errors.New(msg)
	}

	return data, nil
}

func (store *Store) Close() error {

	return store.client.Close()
}

func (store *Store) getLatestSpec(tenantId string) (string, error) {

	res := ""
	latestCreated := time.Date(1980, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	folder := fmt.Sprintf("%s/spec", tenantId)
	it := store.client.Bucket(store.bucket).
		Objects(context.Background(), &storage.Query{Prefix: folder})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logrus.Errorf("failed to iterate over '%s' with '%v'", folder, err)
			return "", err
		}

		if latestCreated.Before(attrs.Created) {
			latestCreated = attrs.Created
			// attrs.Name looks like this: "{tenant-id}/spec/1685962955"
			res = attrs.Name
		}
	}

	return res, nil
}
