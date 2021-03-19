package vapp

import (
	"github.com/inspur-ics/ics-go-sdk/client"
	"github.com/inspur-ics/ics-go-sdk/common"
)

type VappService struct {
	common.RestAPI
}

func NewVappService(c *client.Client) *VappService {
	vs := VappService{
		common.RestAPI{
			RestAPITripper: c,
		},
	}
	return &vs
}
