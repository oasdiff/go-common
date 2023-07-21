package ds

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/oasdiff/go-common/env"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type FilterField struct {
	Name     string
	Operator string
	Value    interface{}
}

type Kind string

const (
	KindTenant  Kind = "tenant"
	KindWebhook Kind = "webhook"
)

type Client interface {
	Get(kind Kind, id string, dst interface{}) error
	GetAll(kind Kind, dst interface{}) error
	GetFilter(kind Kind, filters []FilterField, dst interface{}) error
	Put(kind Kind, id string, src interface{}) error
	Close() error
}

type ClientWrapper struct {
	dsc       *datastore.Client
	namespace string
}

func NewClientWrapper(project string, namespace string) Client {

	if key := env.GetDatastoreKey(); key != "" {
		conf, err := google.JWTConfigFromJSON([]byte(key), datastore.ScopeDatastore)
		if err != nil {
			logrus.Fatalf("failed to config datastore JWT from JSON key with '%v'", err)
		}

		ctx := context.Background()
		opt := []option.ClientOption{option.WithTokenSource(conf.TokenSource(ctx))}

		dataStoreEndPoint := os.Getenv("DATASTORE_ENDPOINT")
		if dataStoreEndPoint != "" {
			opt = append(opt, option.WithEndpoint(dataStoreEndPoint))
		}

		client, err := datastore.NewClient(ctx, project, opt...)
		if err != nil {
			logrus.Fatalf("failed to create datastore client with '%v'", err)
		}

		return &ClientWrapper{dsc: client, namespace: namespace}
	}

	client, err := datastore.NewClient(context.Background(), project)
	if err != nil {
		logrus.Fatalf("failed to create datastore client without token with '%v'", err)
	}

	return &ClientWrapper{dsc: client, namespace: namespace}
}

func (c *ClientWrapper) Get(kind Kind, id string, dst interface{}) error {

	err := c.dsc.Get(context.Background(), getKey(c.namespace, kind, id), dst)
	if err != nil {
		logrus.Errorf("failed to get '%s' id '%s' from datastore namespace '%s' with '%v'", kind, id, c.namespace, err)
	}

	return err
}

func (c *ClientWrapper) GetAll(kind Kind, dst interface{}) error {

	q := datastore.NewQuery(string(kind)).Namespace(c.namespace)
	_, err := c.dsc.GetAll(context.Background(), q, dst)
	if err != nil {
		logrus.Errorf("failed to get all '%s' from datastore namespace '%s' with '%v'", kind, c.namespace, err)
	}

	return err
}

func (c *ClientWrapper) GetFilter(kind Kind, filters []FilterField, dst interface{}) error {

	q := datastore.NewQuery(string(kind)).Namespace(c.namespace)
	for _, currFilter := range filters {
		q = q.FilterField(currFilter.Name, currFilter.Operator, currFilter.Value)
	}
	_, err := c.dsc.GetAll(context.Background(), q, dst)
	if err != nil {
		msg := fmt.Sprintf("failed to get kind '%s' filter '%+v' from datastore namespace '%s' with '%v'",
			kind, filters, c.namespace, err)
		if IsNoSuchEntityError(err) {
			logrus.Debug(msg)
		} else {
			logrus.Error(msg)
		}
	}

	return err
}

func (c *ClientWrapper) Put(kind Kind, id string, src interface{}) error {

	_, err := c.dsc.Put(context.Background(), getKey(c.namespace, kind, id), src)
	if err != nil {
		logrus.Errorf("failed to update '%s/%s' item '%+v' type: '%T' with '%v'", c.namespace, kind, src, src, err)
	} else {
		logrus.Infof("created '%s/%s: %+v' type: '%T'", c.namespace, kind, src, src)
	}

	return err
}

func (c *ClientWrapper) Close() error { return c.dsc.Close() }

func getKey(namespace string, kind Kind, id string) *datastore.Key {

	res := datastore.NameKey(string(kind), id, nil)
	res.Namespace = namespace

	return res
}

func IsNoSuchEntityError(err error) bool {

	if err == nil {
		return false
	}

	return err.Error() == datastore.ErrNoSuchEntity.Error()
}
