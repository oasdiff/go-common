package env_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/oasdiff/go-common/env"
	"github.com/stretchr/testify/require"
)

func TestGetEnvWithDefault(t *testing.T) {

	const value = "http://test.me.com"
	require.Equal(t, value, env.GetEnvWithDefault("MY_TEST_CONNECTION", value))
}

func TestGetEnvIntWithDefault(t *testing.T) {

	const value = 14
	require.Equal(t, value, env.GetEnvIntWithDefault("MY_TEST_CONNECTION", value))
}

func TestGetEnvIntWithDefault_Convert(t *testing.T) {

	key := uuid.NewString()
	const value = 14

	os.Setenv(key, strconv.Itoa(value))
	require.Equal(t, value, env.GetEnvIntWithDefault(key, value))
}
