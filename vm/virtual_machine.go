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

//FIXME TODO.WANGYONGCHAO
func (v *VirtualMachineService) GetVMByIP(ctx context.Context, ipAddy string) (*types.VirtualMachine, error) {
    vm, err := methods.GetVMById(ctx, v.RestAPITripper, ipAddy)
    return vm, err
}

//FIXME TODO.WANGYONGCHAO
func (v *VirtualMachineService) GetVMByName(ctx context.Context, name string) (*types.VirtualMachine, error) {
    vm, err := methods.GetVMById(ctx, v.RestAPITripper, name)
    return vm, err
}
//FIXME TODO.WANGYONGCHAO
func (v *VirtualMachineService) GetVMByPath(ctx context.Context, name string) (*types.VirtualMachine, error) {
    return nil, nil
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

//func GetVMList(vmctx context.Context, connection *icsgo.ICSConnection) ([]types.VirtualMachine, error) {
//   ctx := context.Background()
//   err := connection.Connect(ctx)
//   if err != nil {
//       fmt.Println("Create ics connection error!")
//       return nil, err
//   }
//
//   Client := NewVirtualMachineService(connection.Client)
//   vmlist, err := Client.GetVMListSv()
//   if err != nil {
//       fmt.Println("Create ics connection error!")
//       return nil, err
//   }
//
//   return vmlist, err
//}
//
//func GetVMByUUId(vmctx context.Context, connection *icsgo.ICSConnection, vmUUID string) (*types.VirtualMachine, error) {
//   var tvm types.VirtualMachine
//   vmlist, err := GetVMList(vmctx, connection)
//   if err != nil {
//       fmt.Println("Create ics connection error!")
//       return nil, err
//   }
//
//   for _, vm := range vmlist {
//       if vmUUID == vm.UUID{
//           return &vm, err
//       }
//   }
//   return &tvm, err
//}
//
//func GetVMByName(vmctx context.Context, connection *icsgo.ICSConnection, vmName string) (*types.VirtualMachine, error) {
//   var tvm types.VirtualMachine
//   vmlist, err := GetVMList(vmctx, connection)
//   if err != nil {
//       fmt.Println("Create ics connection error!")
//       return nil, err
//   }
//
//   for _, vm := range vmlist {
//       if vmName == vm.Name{
//           return &vm, err
//       }
//   }
//   return &tvm, err
//}
//
//func GetHostByVMUUId(vmctx context.Context, connection *icsgo.ICSConnection, vmUUID string) (*types.Host, error) {
//   vm, err := GetVMByUUId(vmctx, connection, vmUUID)
//   if err != nil {
//       return nil, err
//   }
//
//   host, err := ht.GetHost(vmctx, connection, string(vm.HostID))
//   if err != nil {
//       return nil, err
//   }
//   return host, nil
//}
//
//func GetHostByVMName(vmctx context.Context, connection *icsgo.ICSConnection, vmName string) (*types.Host, error) {
//   vm, err := GetVMByName(vmctx, connection, vmName)
//   if err != nil {
//       return nil, err
//   }
//
//   host, err := ht.GetHost(vmctx, connection, string(vm.HostID))
//   if err != nil {
//       return nil, err
//   }
//   return host, nil
//}