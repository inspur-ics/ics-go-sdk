package vm

import (
    "context"
    "github.com/inspur-ics/ics-go-sdk/client/methods"
    "github.com/inspur-ics/ics-go-sdk/client/types"
)

func (v *VirtualMachineService) VM(id string) (*types.VM, error) {
    ctx := context.Background()
    vm, err := methods.GetVMById(ctx, v.restAPITripper, id)
    return vm, err
}

func (v *VirtualMachineService) VMPageList(req *types.VMPageReq) (*types.VMPageResponse, error) {
    ctx := context.Background()
    vmPages, err := methods.GetVMPageList(ctx, v.restAPITripper, req)
    return vmPages, err
}