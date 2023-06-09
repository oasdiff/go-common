package env

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetBucket() string { return failIfEmpty("GCS_BUCKET") }

func GetGCloudProject() string { return failIfEmpty("GCLOUD_PROJECT") }

func GetDatastoreKey() string { return os.Getenv("DATASTORE_KEY") }

func GetStorageKey() string { return os.Getenv("STORAGE_KEY") }

func GetSlackInfoUrl() string { return failIfEmpty("SLACK_INFO_URL") }

func failIfEmpty(key string) string {

	res := os.Getenv(key)
	if res == "" {
		logrus.Fatalf("Please, add environment variable '%s'", key)
	}

	return res
}
