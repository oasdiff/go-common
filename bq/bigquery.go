package bq

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/oasdiff/go-common/env"
	"github.com/oasdiff/go-common/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	DatasetDev     = "dev"
	TableTelemetry = "telemetry"
)

type Client interface {
	CreateDataset(dataset string) error
	DeleteDataset(dataset string) error
	CreateTable(dataset string, table string, metadata *bigquery.TableMetadata) error
	DeleteTable(dataset string, table string) error
	DatasetIterator(callback func(dataset *bigquery.Dataset)) error
	Insert(dataset string, table string, src interface{}) error
	Query(q string) error
	QueryIterator(q string, params []bigquery.QueryParameter) (Iterator, error)
	GetQueryStats(q string, params []bigquery.QueryParameter) (*bigquery.JobStatistics, error)
	Close() error
}

type ClientImpl struct {
	bqClient *bigquery.Client
}

func NewClient(project string) Client {

	client, err := bigquery.NewClient(context.Background(), project)
	if err != nil {
		log.Fatalf("failed to create bigquery client without token with '%v'", err)
	}

	return &ClientImpl{bqClient: client}
}

func NewClientAuth(project string) Client {

	conf, err := google.JWTConfigFromJSON([]byte(env.GetBigQueryToken()), bigquery.Scope)
	if err != nil {
		log.Fatalf("failed to config big-query JWT with '%v'", err)
	}

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, project, option.WithTokenSource(conf.TokenSource(ctx)))
	if err != nil {
		log.Fatalf("failed to create bigquery client with '%v'", err)
	}

	return &ClientImpl{bqClient: client}
}

func (client *ClientImpl) Insert(dataset string, table string, src interface{}) error {

	arr := util.ToInterfaceSlice(src)
	count := len(arr)
	start, end := 0, 99
	for {
		if end > count {
			end = count
		}
		err := client.bqClient.Dataset(dataset).Table(table).Inserter().Put(context.Background(), arr[start:end])
		if err != nil {
			log.Errorf("failed to persist '%s.%s' with '%v'", dataset, table, err)
			return err
		}
		log.Debugf("inserted '%d:%d' into '%s.%s'", start, end, dataset, table)
		if end == count {
			break
		}
		start += 100
		end += 100
	}

	return nil
}

func (client *ClientImpl) QueryIterator(q string, params []bigquery.QueryParameter) (Iterator, error) {

	query := client.bqClient.Query(q)
	query.Parameters = params
	rowIterator, err := query.Read(context.Background())
	if err != nil {
		return nil, err
	}

	return newRawIterator(rowIterator), err
}

func (client *ClientImpl) Query(q string) error {

	_, err := client.bqClient.Query(q).Read(context.Background())

	return err
}

func (client *ClientImpl) GetQueryStats(q string, params []bigquery.QueryParameter) (*bigquery.JobStatistics, error) {

	query := client.bqClient.Query(q)
	query.Parameters = params
	query.QueryConfig.DryRun = true

	job, err := query.Run(context.Background())
	if err != nil {
		log.Errorf("get query stats failed with '%v'", err)
		return nil, err
	}

	return job.LastStatus().Statistics, nil
}

func (client *ClientImpl) DatasetIterator(callback func(dataset *bigquery.Dataset)) error {

	datasetIterator := client.bqClient.Datasets(context.Background())
	for {
		currDataset, err := datasetIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Errorf("failed to iterate on datasets with '%v'", err)
			return err
		}
		if currDataset == nil {
			break
		}
		callback(currDataset)
	}

	return nil
}

func (client *ClientImpl) CreateDataset(dataset string) error {

	log.Infof("creating big-query dataset '%s'...", dataset)
	err := client.bqClient.Dataset(dataset).Create(context.Background(),
		&bigquery.DatasetMetadata{Name: dataset})
	if err != nil {
		log.Errorf("failed to create dataset '%s' with '%v'", dataset, err)
	}

	return err
}

func (client *ClientImpl) DeleteDataset(dataset string) error {

	log.Infof("deleting big-query dataset '%s'...", dataset)
	err := client.bqClient.Dataset(dataset).Delete(context.Background())
	if err != nil {
		log.Errorf("failed to delete dataset '%s' with '%v'", dataset, err)
	}

	return err
}

func (client *ClientImpl) CreateTable(dataset string, table string, metadata *bigquery.TableMetadata) error {

	log.Infof("creating big-query table '%s.%s'...", dataset, table)
	err := client.bqClient.Dataset(dataset).Table(table).Create(
		context.Background(), metadata)
	if err != nil {
		log.Errorf("failed to create table '%s.%s' with '%v'", dataset, table, err)
	}

	return err
}

func (client *ClientImpl) DeleteTable(dataset string, table string) error {

	log.Infof("deleting big-query table '%s.%s'...", dataset, table)
	err := client.bqClient.Dataset(dataset).Table(table).Delete(context.Background())
	if err != nil {
		log.Errorf("failed to delete table '%s.%s' with '%v'", dataset, table, err)
	}

	return err
}

func (client *ClientImpl) Close() error {

	return client.bqClient.Close()
}
