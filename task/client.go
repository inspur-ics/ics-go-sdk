package task

import (
	"github.com/inspur-ics/ics-go-sdk/client"
	"github.com/inspur-ics/ics-go-sdk/common"
)

type TaskService struct {
	common.RestAPI
}

func NewTaskService(c *client.Client) *TaskService {
	task := TaskService{
		common.RestAPI{
			RestAPITripper: c,
		},
	}
	return &task
}
