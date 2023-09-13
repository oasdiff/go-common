package util

import (
	"net/http"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
)

func CreateHttpClient() *http.Client {

	if proxy := getProxyUrl(); proxy != nil {
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	}

	return &http.Client{}
}

func getProxyUrl() *url.URL {

	proxy := os.Getenv("HTTP_PROXY")
	if proxy == "" {
		proxy = os.Getenv("HTTPS_PROXY")
	}
	if proxy != "" {
		res, err := url.Parse(proxy)
		if err != nil {
			logrus.Infof("invalid proxy url '%s' with '%v'", proxy, err)
			return nil
		}

		return res
	}

	return nil
}
