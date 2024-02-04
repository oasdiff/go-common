package env

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetLogLevel() string { return failIfEmpty("LOG_LEVEL") }

func GetGCPProject() string { return failIfEmpty("GCP_PROJECT") }

func GetGCPLocation() string { return failIfEmpty("GCP_LOCATION") }

func GetGCPQueue() string { return failIfEmpty("GCP_QUEUE") }

func GetGCPVertexAIToken() string { return failIfEmpty("GCP_VERTEX_AI_TOKEN") }

func GetGCPDatastoreToken() string { return failIfEmpty("GCP_DATASTORE_TOKEN") }

func GetGCPDatastoreNamespace() string { return failIfEmpty("GCP_DATASTORE_NAMESPACE") }

func GetBigQueryDataset() string { return failIfEmpty("GCP_BIG_QUERY_DATASET") }

func GetBigQueryToken() string { return failIfEmpty("GCP_BIG_QUERY_TOKEN") }

func GetGCPStorageBucket() string { return failIfEmpty("GCP_STORAGE_BUCKET") }

func GetGCPStorageKey() string { return os.Getenv("GCP_STORAGE_KEY") }

func GetTaskSubscriberUrl() string { return os.Getenv("OASDIFF_TASK_SUBSCRIBER_URL") }

func GetSlackInfoUrl() string { return failIfEmpty("SLACK_INFO_URL") }

func GetGoogleAnalyticsMeasurementId() string { return failIfEmpty("GOOGLE_ANALYTICS_MEASUREMENT_ID") }

func GetGoogleAnalyticsApiSecret() string { return failIfEmpty("GOOGLE_ANALYTICS_API_SECRET") }

func failIfEmpty(key string) string {

	res := os.Getenv(key)
	if res == "" {
		logrus.Fatalf("Please, add environment variable '%s'", key)
	}

	return res
}
