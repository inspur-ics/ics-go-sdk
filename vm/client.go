package vm

import (
    "github.com/inspur-ics/ics-go-sdk/client"
    "github.com/inspur-ics/ics-go-sdk/client/restful"
)

type VirtualMachineService struct {
    restAPITripper restful.RestAPITripper
}

// NewVirtualMachineService returns the session's virtual machine service.
func NewVirtualMachineService(c *client.Client) *VirtualMachineService {
    vm := VirtualMachineService{
        restAPITripper: c,
    }

    return &vm
}