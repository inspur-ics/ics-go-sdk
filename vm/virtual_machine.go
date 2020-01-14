package vm

//import (
//    "context"
//    "fmt"
//    icsgo "github.com/inspur-ics/ics-go-sdk"
//    "github.com/inspur-ics/ics-go-sdk/client/methods"
//    "github.com/inspur-ics/ics-go-sdk/client/types"
//    ht "github.com/inspur-ics/ics-go-sdk/host"
//)
//
//func (v *VirtualMachineService) VM(id string) (*types.VirtualMachine, error) {
//    ctx := context.Background()
//    vm, err := methods.GetVMById(ctx, v.RestAPITripper, id)
//    return vm, err
//}
//
//func (v *VirtualMachineService) VMPageList(req *types.VMPageReq) (*types.VMPageResponse, error) {
//    ctx := context.Background()
//    vmPages, err := methods.GetVMPageList(ctx, v.RestAPITripper, req)
//    return vmPages, err
//}
//
//func (v *VirtualMachineService) PowerOnVM(id string) (*types.Task, error) {
//    ctx := context.Background()
//    task, err := methods.PowerOnVMById(ctx, v.RestAPITripper, id)
//    return task, err
//}
//
//func (v *VirtualMachineService) GetVMListSv() ([]types.VirtualMachine, error) {
//    ctx := context.Background()
//    vmlist, err := methods.GetVMList(ctx, v.RestAPITripper)
//    return vmlist.Items, err
//}
//
//func GetVMList(vmctx context.Context, connection *icsgo.ICSConnection) ([]types.VirtualMachine, error) {
//    ctx := context.Background()
//    err := connection.Connect(ctx)
//    if err != nil {
//        fmt.Println("Create ics connection error!")
//        return nil, err
//    }
//
//    Client := NewVirtualMachineService(connection.Client)
//    vmlist, err := Client.GetVMListSv()
//    if err != nil {
//        fmt.Println("Create ics connection error!")
//        return nil, err
//    }
//
//    return vmlist, err
//}
//
//func GetVMByUUId(vmctx context.Context, connection *icsgo.ICSConnection, vmUUID string) (*types.VirtualMachine, error) {
//    var tvm types.VirtualMachine
//    vmlist, err := GetVMList(vmctx, connection)
//    if err != nil {
//        fmt.Println("Create ics connection error!")
//        return nil, err
//    }
//
//    for _, vm := range vmlist {
//        if vmUUID == vm.UUID{
//            return &vm, err
//        }
//    }
//    return &tvm, err
//}
//
//func GetVMByName(vmctx context.Context, connection *icsgo.ICSConnection, vmName string) (*types.VirtualMachine, error) {
//    var tvm types.VirtualMachine
//    vmlist, err := GetVMList(vmctx, connection)
//    if err != nil {
//        fmt.Println("Create ics connection error!")
//        return nil, err
//    }
//
//    for _, vm := range vmlist {
//        if vmName == vm.Name{
//            return &vm, err
//        }
//    }
//    return &tvm, err
//}
//
//func GetHostByVMUUId(vmctx context.Context, connection *icsgo.ICSConnection, vmUUID string) (*types.Host, error) {
//    vm, err := GetVMByUUId(vmctx, connection, vmUUID)
//    if err != nil {
//        return nil, err
//    }
//
//    host, err := ht.GetHost(vmctx, connection, string(vm.HostID))
//    if err != nil {
//        return nil, err
//    }
//    return host, nil
//}
//
//func GetHostByVMName(vmctx context.Context, connection *icsgo.ICSConnection, vmName string) (*types.Host, error) {
//    vm, err := GetVMByName(vmctx, connection, vmName)
//    if err != nil {
//        return nil, err
//    }
//
//    host, err := ht.GetHost(vmctx, connection, string(vm.HostID))
//    if err != nil {
//        return nil, err
//    }
//    return host, nil
//}