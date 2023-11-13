package slack

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/oasdiff/go-common/env"
	"github.com/sirupsen/logrus"
)

type Client interface {
	Info(message string) error
}

type ClientWrapper struct {
	info string
}

func NewClientWrapper() Client {

	return &ClientWrapper{
		info: env.GetSlackInfoUrl(),
	}
}

func (c *ClientWrapper) Info(message string) error {

	return Send(c.info, message)
}

func Send(channelHook string, message string) error {

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}{Type: "mrkdwn", Text: message})
	if err != nil {
		logrus.Errorf("failed to encode message '%s' with '%v'", message, err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, channelHook, &buf)
	if err != nil {
		logrus.Errorf("failed to create slack message request '%s' with '%v'", message, err)
		return err
	}
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("failed to send message to slack '%s' with '%v'", message, err)
		return err
	}
	if http.StatusOK != resp.StatusCode {
		logrus.Errorf("failed to send slack message '%s' with '%s'", message, resp.Status)
		return err
	}

	return nil
}
