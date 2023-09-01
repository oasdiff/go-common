package task

import (
	"fmt"

	taskspb "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"github.com/oasdiff/go-common/env"
)

type TaskBuilder struct {
	requestParent string
	subscriberUrl string
}

func NewTaskBuilder() *TaskBuilder {

	return &TaskBuilder{
		requestParent: fmt.Sprintf("projects/%s/locations/%s/queues/%s", env.GetGCPProject(), env.GetGCPLocation(), env.GetGCPQueue()),
		subscriberUrl: env.GetTaskSubscriberUrl(),
	}
}

func (tb *TaskBuilder) CreateRequest(body []byte) *taskspb.CreateTaskRequest {

	return &taskspb.CreateTaskRequest{
		Parent: tb.requestParent,
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        tb.subscriberUrl,
					Body:       body,
				},
			},
		},
	}
}
