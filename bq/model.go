package bq

import (
	"cloud.google.com/go/bigquery"
)

type Telemetry struct {
	Id                 string   `bigquery:"id"`
	Time               int64    `bigquery:"time"`
	MachineId          string   `bigquery:"machine_id"`
	Runtime            string   `bigquery:"runtime"`
	Platform           string   `bigquery:"platform"`
	Command            string   `bigquery:"command"`
	Args               []string `bigquery:"args"`
	Flags              string   `bigquery:"flags"`
	Application        string   `bigquery:"application"`
	ApplicationVersion string   `bigquery:"application_version"`
}

// Save implements the ValueSaver interface
func (t *Telemetry) Save() (map[string]bigquery.Value, string, error) {

	return map[string]bigquery.Value{
		"id":                  t.Id,
		"time":                t.Time,
		"machine_id":          t.MachineId,
		"runtime":             t.Runtime,
		"platform":            t.Platform,
		"command":             t.Command,
		"args":                t.Args,
		"flags":               t.Flags,
		"application":         t.Application,
		"application_version": t.ApplicationVersion,
	}, "", nil
}
