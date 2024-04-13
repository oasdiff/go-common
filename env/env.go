package env

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

func Must(key string) string {

	res, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Environment variable '%s' not found", key))
	}

	return res
}

func GetWithDefault(key string, defaultValue string) string {

	res, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return res
}

func GetIntWithDefault(key string, defaultValue int) int {

	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	res, err := strconv.Atoi(value)
	if err != nil {
		slog.Warn("failed to parse environment variable that should be int. using default",
			"error", err, "env key", key, "value", value, "default", defaultValue)
		return defaultValue
	}

	return res
}

func GetFloat32WithDefault(key string, defaultValue float32) float32 {

	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	res, err := strconv.ParseFloat(value, 32)
	if err != nil {
		slog.Warn("failed to parse environment variable that should be float32. using default",
			"error", err, "env key", key, "value", value, "default", defaultValue)
		return defaultValue
	}

	return float32(res)
}

func GetLogLevel() string { return GetWithDefault("LOG_LEVEL", "info") }

func GetGCPProject() string { return Must("GCP_PROJECT") }

func GetGCPLocation() string { return Must("GCP_LOCATION") }

func GetGCPQueue() string { return Must("GCP_QUEUE") }

func GetGCPVertexAIToken() string { return Must("GCP_VERTEX_AI_TOKEN") }

func GetGCPDatastoreToken() string { return Must("GCP_DATASTORE_TOKEN") }

func GetGCPDatastoreNamespace() string { return Must("GCP_DATASTORE_NAMESPACE") }

func GetBigQueryDataset() string { return Must("GCP_BIG_QUERY_DATASET") }

func GetBigQueryToken() string { return Must("GCP_BIG_QUERY_TOKEN") }

func GetGCPStorageBucket() string { return Must("GCP_STORAGE_BUCKET") }

func GetGCPStorageKey() string { return os.Getenv("GCP_STORAGE_KEY") }

func GetTaskSubscriberUrl() string { return os.Getenv("OASDIFF_TASK_SUBSCRIBER_URL") }

func GetSlackInfoUrl() string { return Must("SLACK_INFO_URL") }

func GetGoogleAnalyticsMeasurementId() string { return Must("GOOGLE_ANALYTICS_MEASUREMENT_ID") }

func GetGoogleAnalyticsApiSecret() string { return Must("GOOGLE_ANALYTICS_API_SECRET") }
