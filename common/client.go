package common

import (
    "context"
    "github.com/inspur-ics/ics-go-sdk/client/methods"
    "github.com/inspur-ics/ics-go-sdk/client/restful"
    "github.com/inspur-ics/ics-go-sdk/client/types"
)

type RestAPI struct {
    RestAPITripper restful.RestAPITripper
}

func(r *RestAPI) TraceTaskProcess(task *types.Task) (*types.TaskInfo, error) {
    ctx := context.Background()
    taskInfo, err := methods.GetTaskInfo(ctx, r.RestAPITripper, task)
    return taskInfo, err
}