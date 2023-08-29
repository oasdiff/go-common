package bq

import "cloud.google.com/go/bigquery"

type ClientInMemory struct{}

func NewClientInMemory() Client {

	return &ClientInMemory{}
}

func (c *ClientInMemory) GetQueryStats(q string, params []bigquery.QueryParameter) (*bigquery.JobStatistics, error) {
	return nil, nil
}

func (c *ClientInMemory) Insert(string, string, interface{}) error {

	return nil
}

func (*ClientInMemory) QueryIterator(string, []bigquery.QueryParameter) (Iterator, error) {

	panic("invalid operation")
}

func (*ClientInMemory) Query(string) error {

	return nil
}

func (*ClientInMemory) DatasetIterator(func(*bigquery.Dataset)) error { return nil }

func (*ClientInMemory) DeleteDataset(string) error { return nil }

func (*ClientInMemory) CreateDataset(string) error { return nil }

func (*ClientInMemory) CreateTable(string, string, *bigquery.TableMetadata) error {

	panic("invalid operation")
}

func (*ClientInMemory) DeleteTable(string, string) error {

	panic("invalid operation")
}

func (*ClientInMemory) Close() error {

	panic("invalid operation")
}
