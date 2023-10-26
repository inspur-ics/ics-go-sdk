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

func (v *VirtualMachineService) SetVM(ctx context.Context, vmInfo types.VirtualMachine) (*types.Task, error) {
	task, err := methods.SetVM(ctx, v.RestAPITripper, vmInfo)
	return task, err
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

func (v *VirtualMachineService) GetVMByUUID(ctx context.Context, vmUUID string) (*types.VirtualMachine, error) {
	vms, err := v.GetVMList(ctx)
	for _, vm := range vms {
		if vmUUID == vm.UUID {
			return &vm, err
		}
	}
	return nil, err
}

func (v *VirtualMachineService) GetVMByIP(ctx context.Context, ip string) (*types.VirtualMachine, error) {
	vm, err := methods.GetVMByIP(ctx, v.RestAPITripper, ip)
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

func (v *VirtualMachineService) PowerOnVM(ctx context.Context, id string) (*types.Task, error) {
	task, err := methods.PowerOnVMById(ctx, v.RestAPITripper, id)
	return task, err
}

func (v *VirtualMachineService) PowerOffVM(ctx context.Context, id string) (*types.Task, error) {
	task, err := methods.PowerOffVMById(ctx, v.RestAPITripper, id)
	return task, err
}

func (v *VirtualMachineService) ShutdownVM(ctx context.Context, id string) (*types.Task, error) {
	task, err := methods.ShutdownVMById(ctx, v.RestAPITripper, id)
	return task, err
}

func (v *VirtualMachineService) RestartVM(ctx context.Context, id string) (*types.Task, error) {
	task, err := methods.RestartVMById(ctx, v.RestAPITripper, id)
	return task, err
}

func (v *VirtualMachineService) DeleteVM(ctx context.Context, id string, deleteFile bool, removeData bool) (*types.Task, error) {
	task, err := methods.DeleteVMById(ctx, v.RestAPITripper, id, deleteFile, removeData)
	return task, err
}

func (v *VirtualMachineService) CreateVMByTemplate(ctx context.Context, vmSpec types.VirtualMachine, quickClone bool) (*types.Task, error) {
	task, err := methods.CreateVMByTemplate(ctx, v.RestAPITripper, vmSpec, quickClone)
	return task, err
}

func (v *VirtualMachineService) GetVMTemplateList(ctx context.Context) ([]types.VirtualMachine, error) {
	vmTemplateList, err := methods.GetVMTemplateList(ctx, v.RestAPITripper)
	return vmTemplateList.Items, err
}

func (v *VirtualMachineService) GetVMTemplate(ctx context.Context, id string) (*types.VirtualMachine, error) {
	vmt, err := methods.GetVMTemplateById(ctx, v.RestAPITripper, id)
	return vmt, err
}

func (v *VirtualMachineService) GetVMTemplateByUUID(ctx context.Context, uuid string) (*types.VirtualMachine, error) {
	vmtList, err := v.GetVMTemplateList(ctx)
	for _, vmt := range vmtList {
		if uuid == vmt.UUID {
			return &vmt, err
		}
	}
	return nil, err
}

func (v *VirtualMachineService) GetVMTemplateByName(ctx context.Context, name string) (*types.VirtualMachine, error) {
	vmtList, err := v.GetVMTemplateList(ctx)
	for _, vmt := range vmtList {
		if name == vmt.Name {
			return &vmt, err
		}
	}
	return nil, err
}

func (v *VirtualMachineService) GetVMPowerStateByID(ctx context.Context, id string) (*types.VMPowerState, error) {
	vm, err := methods.GetVMById(ctx, v.RestAPITripper, id)

	if vm.Status == "STARTED" {
		vm.PowerState = "poweredOn"
		return &vm.PowerState, err
	} else if vm.Status == "STOPPED" {
		vm.PowerState = "poweredOff"
		return &vm.PowerState, err
	} else if vm.Status == "PAUSED" {
		vm.PowerState = "paused"
		return &vm.PowerState, err
	} else {
		vm.PowerState = "other"
		return &vm.PowerState, err
	}
}

func (v *VirtualMachineService) GetVMNetState(ctx context.Context, id string) ([]types.Nic, error) {
	vm, err := methods.GetVMById(ctx, v.RestAPITripper, id)
	nic := vm.Nics
	return nic, err
}
