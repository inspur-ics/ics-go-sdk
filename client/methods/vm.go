package methods

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/inspur-ics/ics-go-sdk/client/restful"
    "github.com/inspur-ics/ics-go-sdk/client/types"
)

func GetVMById(ctx context.Context, r restful.RestAPITripper, vmId string) (*types.VirtualMachine, error) {
    var reqBody      *types.Common
    var api          types.ICSApi
    var response = types.VirtualMachine{}

    if len(vmId) <= 0 {
        vmId = "anonymous"
    }
    api.Api = fmt.Sprintf("/vms/%s",vmId)
    api.Token = true

    resp, err := r.GetTrip(ctx, api, reqBody)
    respBody, err1 := HandleResponse(resp, err)
    if err1 != nil {
        err = err1
    } else if respBody != nil {
        jsonErr := json.Unmarshal([]byte(respBody), &response)
        err = JsonError(jsonErr)
    }

    return &response, err
}

func GetVMPageList(ctx context.Context, r restful.RestAPITripper, req *types.VMPageReq) (*types.VMPageResponse, error) {
    var api          types.ICSApi
    var response = types.VMPageResponse{}

    api.Api = "/vms/"
    api.Token = true

    resp, err := r.GetTrip(ctx, api, req)
    respBody, err1 := HandleResponse(resp, err)
    if err1 != nil {
        err = err1
    } else if respBody != nil {
        jsonErr := json.Unmarshal([]byte(respBody), &response)
        err = JsonError(jsonErr)
    }

    return &response, err
}

func PowerOnVMById(ctx context.Context, r restful.RestAPITripper, vmId string) (*types.Task, error) {
    var reqBody      *types.Common
    var api          types.ICSApi
    var response = types.Task{}

    if len(vmId) <= 0 {
        vmId = "anonymous"
    }
    api.Api = fmt.Sprintf("/vms/%s?action=poweron",vmId)
    api.Token = true

    resp, err := r.PutTrip(ctx, api, reqBody)
    respBody, err1 := HandleResponse(resp, err)
    if err1 != nil {
        err = err1
    } else if respBody != nil {
        jsonErr := json.Unmarshal([]byte(respBody), &response)
        err = JsonError(jsonErr)
    }

    return &response, err
}