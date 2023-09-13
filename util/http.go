package util

import (
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/http/httpproxy"
)

func CreateHttpClient(reqURL *url.URL) *http.Client {

	if proxy := getProxyUrl(reqURL); proxy != nil {
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	}

	return &http.Client{}
}

func getProxyUrl(reqURL *url.URL) *url.URL {

	res, err := httpproxy.FromEnvironment().ProxyFunc()(reqURL)
	if err != nil {
		logrus.Infof("invalid proxy url '%s' with '%v'", res, err)
		return nil
	}

	return res
}
