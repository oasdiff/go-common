package bq

import (
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
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

func NewTelemetry(app string, appVersion string, t time.Time, machineId string,
	runtime string, platform string, command string, args []string, flags map[string]string) *Telemetry {

	// *** temporary fix https://github.com/googleapis/google-cloud-go/issues/6409
	ct := civil.DateTimeOf(t)
	ct.Time.Nanosecond = 0
	// ***

	return &Telemetry{
		Application:        app,
		ApplicationVersion: appVersion,
		Time:               ct,
		MachineId:          machineId,
		Runtime:            runtime,
		Platform:           platform,
		Command:            command,
		Args:               getArgs(args),
		Flags:              toKeyValue(flags),
	}
}

func getArgs(args []string) []string {

	if len(args) == 0 {
		return []string{}
	}

	return args
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
