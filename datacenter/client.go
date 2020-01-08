package dc

import (
    "github.com/inspur-ics/ics-go-sdk/client"
    "github.com/inspur-ics/ics-go-sdk/common"
)

type DatacenterService struct {
    common.RestAPI
}

// NewDatacenterService returns the session's virtual machine service.
func NewDatacenterService(c *client.Client) *DatacenterService {
    dc := DatacenterService{
        common.RestAPI{
            RestAPITripper: c,
        },
    }

    return &dc
}