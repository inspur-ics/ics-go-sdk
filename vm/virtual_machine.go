package vm

import (
    "context"
    "github.com/inspur-ics/ics-go-sdk/client/methods"
    "github.com/inspur-ics/ics-go-sdk/client/types"
)

func (v *VirtualMachineService) VM(id string) (*types.VM, error) {
    ctx := context.Background()
    vm, err := methods.GetVMById(ctx, v.RestAPITripper, id)
    return vm, err
}

func (v *VirtualMachineService) VMPageList(req *types.VMPageReq) (*types.VMPageResponse, error) {
    ctx := context.Background()
    vmPages, err := methods.GetVMPageList(ctx, v.RestAPITripper, req)
    return vmPages, err
}
func (v *VirtualMachineService) PowerOnVM(id string) (*types.Task, error) {
    ctx := context.Background()
    task, err := methods.PowerOnVMById(ctx, v.RestAPITripper, id)
    return task, err
}