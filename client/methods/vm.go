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

func GetVMByIP(ctx context.Context, r restful.RestAPITripper, ipAddy string) (*types.VirtualMachine, error) {
    var reqBody      *types.Common
    var api          types.ICSApi
    var vmPages = types.VMPageResponse{}
    var response = types.VirtualMachine{}

    if len(ipAddy) <= 6 {
        ipAddy = "255.255.255.255"
    }
    api.Api = fmt.Sprintf("/vms?pageSize=1&currentPage=1&sortField=&sort=desc&vnicIp=%s",ipAddy)
    api.Token = true

    resp, err := r.GetTrip(ctx, api, reqBody)
    respBody, err1 := HandleResponse(resp, err)
    if err1 != nil {
        err = err1
    } else if respBody != nil {
        jsonErr := json.Unmarshal([]byte(respBody), &vmPages)
        err = JsonError(jsonErr)
    }
    if vmPages.TotalSize == 1 {
        response = vmPages.Items[0]
    }

    return &response, err
}

func GetVMByName(ctx context.Context, r restful.RestAPITripper, name string) (*types.VirtualMachine, error) {
    var reqBody      *types.Common
    var api          types.ICSApi
    var vmPages = types.VMPageResponse{}
    var response = types.VirtualMachine{}

    if len(name) <= 0 {
        name = "vm@anonymous"
    }
    api.Api = fmt.Sprintf("/vms?pageSize=1&currentPage=1&sortField=&sort=desc&name=%s",name)
    api.Token = true

    resp, err := r.GetTrip(ctx, api, reqBody)
    respBody, err1 := HandleResponse(resp, err)
    if err1 != nil {
        err = err1
    } else if respBody != nil {
        jsonErr := json.Unmarshal([]byte(respBody), &vmPages)
        err = JsonError(jsonErr)
    }
    if vmPages.TotalSize == 1 {
        response = vmPages.Items[0]
    }

    return &response, err
}

func GetVMPageList(ctx context.Context, r restful.RestAPITripper, req *types.VMPageReq) (*types.VMPageResponse, error) {
    var api          types.ICSApi
    var response = types.VMPageResponse{}

    api.Api = "/vms"
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

func GetVMList(ctx context.Context, r restful.RestAPITripper) (*types.VMPageResponse, error) {
    var reqBody      *types.Common
    var api          types.ICSApi
    var response = types.VMPageResponse{}

    api.Api = "/vms/"
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