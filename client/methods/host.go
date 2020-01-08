package methods

import (
	"context"
	"encoding/json"
	"fmt"

	//	"fmt"
	"github.com/inspur-ics/ics-go-sdk/client/restful"
	"github.com/inspur-ics/ics-go-sdk/client/types"
)


func GetHostById(ctx context.Context, r restful.RestAPITripper, hostUUID string) (*types.Host, error) {
	var reqBody      *types.Common
	var api          types.ICSApi
	var response = types.Host{}

	if len(hostUUID) <= 0 {
		hostUUID = "anonymous"
	}
	api.Api = fmt.Sprintf("/hosts/%s",hostUUID)
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