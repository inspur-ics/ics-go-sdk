package methods

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/inspur-ics/ics-go-sdk/client/restful"
	"github.com/inspur-ics/ics-go-sdk/client/types"
)

func GetHostById(ctx context.Context, r restful.RestAPITripper, hostUUID string) (*types.Host, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Host{}

	if len(hostUUID) <= 0 {
		hostUUID = "anonymous"
	}
	api.Api = fmt.Sprintf("/hosts/%s", hostUUID)
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

func GetHostHealthInfoById(ctx context.Context, r restful.RestAPITripper, hostUUID string) (*types.HostHealthInfo, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.HostHealthInfo{}

	if len(hostUUID) <= 0 {
		hostUUID = "anonymous"
	}
	api.Api = fmt.Sprintf("/hosts/%s/healthperform", hostUUID)
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

func GetHostAvailStorages(ctx context.Context, r restful.RestAPITripper, hostUUID string) ([]types.Storage, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response []types.Storage

	if len(hostUUID) <= 0 {
		hostUUID = "anonymous"
	}
	api.Api = fmt.Sprintf("/hosts/%s/availstorages", hostUUID)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func GetHostList(ctx context.Context, r restful.RestAPITripper) (types.HostPageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response types.HostPageResponse

	api.Api = "/hosts"
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func GetHostListByStorageID(ctx context.Context, r restful.RestAPITripper, storageID string) (types.HostPageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response types.HostPageResponse

	api.Api = fmt.Sprintf("/storages/%s/hosts", storageID)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}
	return response, err
}

func GetAvailHostListByStorageID(ctx context.Context, r restful.RestAPITripper, storageID string) ([]types.Host, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response []types.Host

	api.Api = fmt.Sprintf("/storages/%s/availhosts", storageID)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}
	return response, err
}

func GetHostListBySwitchID(ctx context.Context, r restful.RestAPITripper, switchID string) (types.HostPageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response types.HostPageResponse

	api.Api = fmt.Sprintf("/vswitchs/%s/hosts", switchID)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}
	return response, err
}

func GetHostListByDC(ctx context.Context, r restful.RestAPITripper, datacenterPath string) (*types.HostPageResponse, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response types.HostPageResponse

	api.Api = fmt.Sprintf("/datacenters/%s/hosts", datacenterPath)
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

func GetHostAccessibleDatastoreList(ctx context.Context, r restful.RestAPITripper, hostId string) ([]types.Storage, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response []types.Storage

	api.Api = fmt.Sprintf("hosts/%s/availstorages", hostId)
	api.Token = true

	resp, err := r.GetTrip(ctx, api, reqBody)
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return response, err
}

func PowerOnHostById(ctx context.Context, r restful.RestAPITripper, hostId string) (*types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	if len(hostId) <= 0 {
		hostId = "anonymous"
	}
	api.Api = fmt.Sprintf("/hosts/%s?action=poweron", hostId)
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

func PowerOffHostById(ctx context.Context, r restful.RestAPITripper, hostId string) (*types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	if len(hostId) <= 0 {
		hostId = "anonymous"
	}
	api.Api = fmt.Sprintf("/hosts/%s?action=poweroff", hostId)
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


func PowerOffSafelyHostById(ctx context.Context, r restful.RestAPITripper, hostId string) (*types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	if len(hostId) <= 0 {
		hostId = "anonymous"
	}
	api.Api = fmt.Sprintf("/hosts/%s?action=poweroffsafely", hostId)
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

func ShutdownHostById(ctx context.Context, r restful.RestAPITripper, hostId string) (*types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	if len(hostId) <= 0 {
		hostId = "anonymous"
	}
	api.Api = fmt.Sprintf("/hosts/%s?action=shutdown", hostId)
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

func RebootHostById(ctx context.Context, r restful.RestAPITripper, hostId string, forceFlag bool) (*types.Task, error) {
	var reqBody *types.Common
	var api types.ICSApi
	var response = types.Task{}

	if len(hostId) <= 0 {
		hostId = "anonymous"
	}
	api.Api = fmt.Sprintf("/hosts/%s?action=reboot", hostId)
	api.Token = true

	var resp *restful.Response
	var err error

	if forceFlag {
		var req = struct {
			ForceFlag bool `json:"forceFlag"`
		}{ ForceFlag: forceFlag }
		api.Api = api.Api + "&forceFlag=true"

		resp, err = r.PutTrip(ctx, api, &req)
	} else {
		resp, err = r.PutTrip(ctx, api, reqBody)
	}
	respBody, err1 := HandleResponse(resp, err)
	if err1 != nil {
		err = err1
	} else if respBody != nil {
		jsonErr := json.Unmarshal([]byte(respBody), &response)
		err = JsonError(jsonErr)
	}

	return &response, err
}

// func RestartHostById(ctx context.Context, r restful.RestAPITripper, hostId string) (*types.Task, error) {
// 	var reqBody *types.Common
// 	var api types.ICSApi
// 	var response = types.Task{}

// 	if len(hostId) <= 0 {
// 		hostId = "anonymous"
// 	}
// 	api.Api = fmt.Sprintf("/hosts/%s?action=restart", hostId)
// 	api.Token = true

// 	resp, err := r.PutTrip(ctx, api, reqBody)
// 	respBody, err1 := HandleResponse(resp, err)
// 	if err1 != nil {
// 		err = err1
// 	} else if respBody != nil {
// 		jsonErr := json.Unmarshal([]byte(respBody), &response)
// 		err = JsonError(jsonErr)
// 	}

// 	return &response, err
// }