package env

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetGCPProject() string { return failIfEmpty("GCP_PROJECT") }

func GetGCPLocation() string { return failIfEmpty("GCP_LOCATION") }

func GetGCPQueue() string { return failIfEmpty("GCP_QUEUE") }

func GetGCPDatastoreKey() string { return os.Getenv("GCP_DATASTORE_KEY") }

func GetGCPStorageBucket() string { return failIfEmpty("GCP_STORAGE_BUCKET") }

func GetGCPStorageKey() string { return os.Getenv("GCP_STORAGE_KEY") }

func GetTaskSubscriberUrl() string { return os.Getenv("OASDIFF_TASK_SUBSCRIBER_URL") }

func GetSlackInfoUrl() string { return failIfEmpty("SLACK_INFO_URL") }

func failIfEmpty(key string) string {

	res := os.Getenv(key)
	if res == "" {
		logrus.Fatalf("Please, add environment variable '%s'", key)
	}

	return res
}
