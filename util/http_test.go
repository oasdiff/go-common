package util_test

import (
	"net/url"
	"os"
	"testing"

	"github.com/oasdiff/go-common/util"
	"github.com/stretchr/testify/require"
)

func TestCreateHttpClient_Proxy(t *testing.T) {

	envProxy := os.Getenv("HTTPS_PROXY")
	require.NoError(t, os.Setenv("HTTPS_PROXY", "https://proxy.me"))
	defer func() { require.NoError(t, os.Setenv("HTTPS_PROXY", envProxy)) }()
	request, err := url.Parse("https://test.me")
	require.NoError(t, err)
	require.NotNil(t, util.CreateHttpClient(request).Transport)
}

func TestCreateHttpClient_NoProxy(t *testing.T) {

	envProxy := os.Getenv("HTTPS_PROXY")
	require.NoError(t, os.Setenv("HTTPS_PROXY", ""))
	defer func() { require.NoError(t, os.Setenv("HTTPS_PROXY", envProxy)) }()
	request, err := url.Parse("https://test.me")
	require.NoError(t, err)
	require.Nil(t, util.CreateHttpClient(request).Transport)
}
