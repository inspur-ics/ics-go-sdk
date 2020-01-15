package host

import (
    "github.com/inspur-ics/ics-go-sdk/client"
    "github.com/inspur-ics/ics-go-sdk/common"
)

type HostService struct {
    common.RestAPI
}

// NewDatacenterService returns the session's host service.
func NewHostService(c *client.Client) *HostService {
    ht := HostService{
        common.RestAPI{
            RestAPITripper: c,
        },
    }

    return &ht
}