package methods

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/inspur-ics/ics-go-sdk/client/restful"
	"github.com/inspur-ics/ics-go-sdk/client/types"
	"reflect"
)

var TypeFunc types.Func

func init() {
	TypeFunc = types.TypeFunc()
}

func HandleResponse(resp *restful.Response, resBody  interface{}) ([]byte, error) {
	var err    error
	var code, msg    string
	var response    []byte

	reflect.TypeOf(resBody)

	if resp.StatusCode() == 200 {
		response = resp.Body()
	} else if resp.StatusCode() == 202 || resp.StatusCode() == 401 || resp.StatusCode() == 403 {
		e := types.SDKError{}
		err := json.Unmarshal([]byte(resp.Body()), &e)
		if err != nil {
			code = "501"
			msg = "Service response error"
			err = &types.SDKError{
				Code: code,
				Message: msg,
			}
		} else {
			err = &e
		}
	} else {
		code = "500"
		msg = "Service unknown error"
		err = &types.SDKError{
			Code: code,
			Message: msg,
		}
	}
	return response, err
}

func ValidUserSession(ctx context.Context, r restful.RestAPITripper, user *types.UserSession) error {
	var reqBody    *interface{}
	var resBody    *interface{}
	var api        types.ICSApi
	var err    error
	var code, msg    string

	userId := "anonymous"
	if len(user.UserId) > 0 {
		userId = user.UserId
	}
	api.Api = fmt.Sprintf("/users/%s/themes",userId)
	api.Token = true

	resp, errs := r.GetTrip(ctx, api, reqBody, resBody)
	if err != nil {
		code = "404"
		msg = errs.Error()
		err = &types.SDKError{
			Code: code,
			Message: msg,
		}
	} else {
		if resp.StatusCode() == 200 {
			err = nil
		} else if resp.StatusCode() == 401 {
			code = "401"
			msg = "the session is not authenticated"
			err = &types.SDKError{
				Code: code,
				Message: msg,
			}
		} else {
			code = "400"
			msg = "Client Request Error"
			err = &types.SDKError{
				Code: code,
				Message: msg,
			}
		}
	}
	return err
}

func Login(ctx context.Context, r restful.RestAPITripper, req *types.Login) (*types.LoginResponse, error) {
	var reqBody    *types.Login
	var resBody    *types.LoginResponse
	var api        types.ICSApi
	var code, msg    string

	var response = types.LoginResponse{}

	reqBody = req
	api.Api = "/authentication"
	api.Token = false

	resp, err := r.PostTrip(ctx, api, reqBody, resBody)
	if err != nil {
		code = "404"
		msg = err.Error()
		err = &types.SDKError{
			Code: code,
			Message: msg,
		}
	} else {
		respBody, err1 := HandleResponse(resp, resBody)
		if err1 != nil {
			err = err1
		} else if respBody != nil {
			err = json.Unmarshal([]byte(respBody), &response)
			if err != nil {
				code = "400"
				msg = "SDK unmarshal data error"
				err = &types.SDKError{
					Code: code,
					Message: msg,
				}
			}
		}
	}

	return &response, err
}