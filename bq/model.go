package bq

import (
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"github.com/oasdiff/telemetry/model"
)

type KeyValue struct {
	Key   string `bigquery:"key"`
	Value string `bigquery:"value"`
}

type Telemetry struct {
	Application        string         `bigquery:"application"`
	ApplicationVersion string         `bigquery:"application_version"`
	Time               civil.DateTime `bigquery:"time"`
	MachineId          string         `bigquery:"machine_id"`
	Runtime            string         `bigquery:"runtime"`
	Platform           string         `bigquery:"platform"`
	Command            string         `bigquery:"command"`
	Args               []string       `bigquery:"args"`
	Flags              []KeyValue     `bigquery:"flags"`
}

func NewTelemetry(t *model.Telemetry) *Telemetry {

	return &Telemetry{
		Application:        t.Application,
		ApplicationVersion: t.ApplicationVersion,
		Time:               civil.DateTimeOf(time.Now()),
		MachineId:          t.MachineId,
		Runtime:            t.Runtime,
		Platform:           t.Platform,
		Command:            t.Command,
		Args:               t.Args,
		Flags:              toKeyValue(t.Flags),
	}
}

// Save implements the ValueSaver interface
func (t *Telemetry) Save() (map[string]bigquery.Value, string, error) {

	return map[string]bigquery.Value{
		"application":         t.Application,
		"application_version": t.ApplicationVersion,
		"time":                t.Time,
		"machine_id":          t.MachineId,
		"runtime":             t.Runtime,
		"platform":            t.Platform,
		"command":             t.Command,
		"args":                t.Args,
		"flags":               t.Flags,
	}, "", nil
}

func GetTelemetryTableMetadata() *bigquery.TableMetadata {

	const fieldTime = "time"

	return &bigquery.TableMetadata{Schema: bigquery.Schema{
		{Name: fieldTime, Type: bigquery.DateTimeFieldType},
		{Name: "application", Type: bigquery.StringFieldType},
		{Name: "application_version", Type: bigquery.StringFieldType},
		{Name: "machine_id", Type: bigquery.StringFieldType},
		{Name: "runtime", Type: bigquery.StringFieldType},
		{Name: "platform", Type: bigquery.StringFieldType},
		{Name: "command", Type: bigquery.StringFieldType},
		{Name: "args", Type: bigquery.StringFieldType, Repeated: true},
		{Name: "flags", Type: bigquery.RecordFieldType, Repeated: true, Schema: bigquery.Schema{
			{Name: "key", Type: bigquery.StringFieldType},
			{Name: "value", Type: bigquery.StringFieldType},
		}},
	}, TimePartitioning: &bigquery.TimePartitioning{
		Type:  bigquery.DayPartitioningType,
		Field: fieldTime,
	}}
}

func toKeyValue(items map[string]string) []KeyValue {

	res := make([]KeyValue, len(items))
	for key, value := range items {
		res = append(res, KeyValue{Key: key, Value: value})
	}

	return res
}
