package common

import (
    "context"
    "github.com/inspur-ics/ics-go-sdk/client/methods"
    "github.com/inspur-ics/ics-go-sdk/client/restful"
    "github.com/inspur-ics/ics-go-sdk/client/types"
    "sync"
    "time"
)



type RestAPI struct {
    RestAPITripper restful.RestAPITripper
}

func (r *RestAPI) TraceTaskProcess(task *types.Task) (*types.TaskInfo, error) {
    var wg sync.WaitGroup
    wg.Add(1)
    go r.tracing(task, wg)
    wg.Wait()

    ctx := context.Background()
    taskInfo, err := methods.GetTaskInfo(ctx, r.RestAPITripper, task)
    return taskInfo, err
}

func (r *RestAPI) tracing(task *types.Task, wg sync.WaitGroup) {
    defer wg.Done()
    ctx := context.Background()
    count := 1
    for {
        taskInfo, err := methods.GetTaskInfo(ctx, r.RestAPITripper, task)
        if err != nil {
            break
        }
        if taskInfo == nil || taskInfo.State == "FINISHED" || taskInfo.State == "ERROR" {
            break
        }
        if count > 72000 {
            break
        }
        time.Sleep(100 * time.Millisecond)
    }
}