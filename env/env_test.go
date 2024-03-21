package env_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/oasdiff/go-common/env"
	"github.com/stretchr/testify/require"
)

func TestMust(t *testing.T) {

	key := uuid.NewString()
	const value = "test-me"
	os.Setenv(key, value)
	require.Equal(t, value, env.Must(key))
}

func TestGetWithDefault(t *testing.T) {

	const value = "http://test.me.com"
	require.Equal(t, value, env.GetWithDefault(uuid.NewString(), value))
}

func TestGetIntWithDefault(t *testing.T) {

	const value = 14
	require.Equal(t, value, env.GetIntWithDefault(uuid.NewString(), value))
}

func TestGetIntWithDefault_Convert(t *testing.T) {

	key := uuid.NewString()
	const value = 14
	os.Setenv(key, strconv.Itoa(value))
	require.Equal(t, value, env.GetIntWithDefault(key, value))
}

func TestGetIntWithDefault_InvalidType(t *testing.T) {

	key := uuid.NewString()
	const value = 14
	os.Setenv(key, "test-me")
	require.Equal(t, value, env.GetIntWithDefault(key, value))
}
