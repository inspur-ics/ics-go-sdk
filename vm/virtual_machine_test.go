package vm

import (
   "context"
   "fmt"
   icsgo "github.com/inspur-ics/ics-go-sdk"
    "github.com/inspur-ics/ics-go-sdk/client/types"
    "testing"
)

func TestVM(t *testing.T) {
   icsConnection := &icsgo.ICSConnection{
       Username: "admin",
       Password: "admin@inspur",
       Hostname: "10.7.11.90",
       Port:     "443",
       Insecure: true,
   }
   ctx := context.Background()
   err := icsConnection.Connect(ctx)
   if err != nil {
       t.Fatal("Create ics connection error!")
   }

   vmClient := NewVirtualMachineService(icsConnection.Client)

   vm, err := vmClient.GetVM(ctx,"8a878bda6f7012c7016f7de1935d047a")
   if vm != nil {
       fmt.Println(vm.Name)
       fmt.Println(vm.ID)
       fmt.Println(vm.UUID)
       fmt.Println(vm.HostID)
       fmt.Println(vm.HostIP)
       fmt.Println(vm.Status)
       fmt.Println(vm.State)
   } else {
       fmt.Println("No VM be found by you point id.")
   }
}

func TestGetVMByIP(t *testing.T) {
    icsConnection := &icsgo.ICSConnection{
        Username: "admin",
        Password: "admin@inspur",
        Hostname: "10.7.11.90",
        Port:     "443",
        Insecure: true,
    }
    ctx := context.Background()
    err := icsConnection.Connect(ctx)
    if err != nil {
        t.Fatal("Create ics connection error!")
    }

    vmClient := NewVirtualMachineService(icsConnection.Client)

    vm, err := vmClient.GetVMByIP(ctx,"10.7.11.81")
    if vm != nil {
        fmt.Printf("%+v", vm)
    } else {
        fmt.Println("No VM be found by you point id.")
    }
}

func TestGetVMByName(t *testing.T) {
    icsConnection := &icsgo.ICSConnection{
        Username: "admin",
        Password: "admin@inspur",
        Hostname: "10.7.11.90",
        Port:     "443",
        Insecure: true,
    }
    ctx := context.Background()
    err := icsConnection.Connect(ctx)
    if err != nil {
        t.Fatal("Create ics connection error!")
    }

    vmClient := NewVirtualMachineService(icsConnection.Client)

    vm, err := vmClient.GetVMByName(ctx,"master")
    if vm != nil {
        fmt.Printf("%+v", vm)
    } else {
        fmt.Println("No VM be found by you point id.")
    }
}

func TestVMPageList(t *testing.T) {
    icsConnection := &icsgo.ICSConnection{
        Username: "admin",
        Password: "admin@inspur",
        Hostname: "10.7.11.90",
        Port:     "443",
        Insecure: true,
    }
    ctx := context.Background()
    err := icsConnection.Connect(ctx)
    if err != nil {
        t.Fatal("Create ics connection error!")
    }

    vmClient := NewVirtualMachineService(icsConnection.Client)

    pakg := types.PageReq{
        1000,
        1,
        "",
        "desc",
    }

    req := &types.VMPageReq{
        pakg,
    }
    vmpagelist, err := vmClient.VMPageList(req)
    if vmpagelist != nil {
        fmt.Printf("%+v", vmpagelist)
    }
}