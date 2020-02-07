package vm

import (
    "context"
    "github.com/inspur-ics/ics-go-sdk/client/methods"
    "github.com/inspur-ics/ics-go-sdk/client/types"
)

func (v *VirtualMachineService) GetVM(ctx context.Context, id string) (*types.VirtualMachine, error) {
   vm, err := methods.GetVMById(ctx, v.RestAPITripper, id)
   return vm, err
}

//FIXME TODO.WANGYONGCHAO
func (v *VirtualMachineService) GetVMByUUID(ctx context.Context, vmUUID string) (*types.VirtualMachine, error) {
    vm, err := methods.GetVMById(ctx, v.RestAPITripper, vmUUID)
    return vm, err
}

func (v *VirtualMachineService) GetVMByIP(ctx context.Context, ipAddy string) (*types.VirtualMachine, error) {
    vm, err := methods.GetVMByIP(ctx, v.RestAPITripper, ipAddy)
    return vm, err
}

func (v *VirtualMachineService) GetVMByName(ctx context.Context, name string) (*types.VirtualMachine, error) {
    vm, err := methods.GetVMByName(ctx, v.RestAPITripper, name)
    return vm, err
}

func (v *VirtualMachineService) GetVMByPath(ctx context.Context, path string) (*types.VirtualMachine, error) {
    vm, err := methods.GetVMById(ctx, v.RestAPITripper, path)
    return vm, err
}

func (v *VirtualMachineService) GetVMList(ctx context.Context) ([]types.VirtualMachine, error) {
    vmlist, err := methods.GetVMList(ctx, v.RestAPITripper)
    return vmlist.Items, err
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