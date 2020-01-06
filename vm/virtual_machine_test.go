package vm

import (
    "context"
    "fmt"
    icsgo "github.com/inspur-ics/ics-go-sdk"
    "testing"
)

func TestVM(t *testing.T) {
    icsConnection := &icsgo.ICSConnection{
        Username: "admin",
        Password: "admin@inspur",
        Hostname: "192.168.220.164",
        Port:     "443",
        Insecure: true,
    }
    ctx := context.Background()
    err := icsConnection.Connect(ctx)
    if err != nil {
        t.Fatal("Create ics connection error!")
    }

    vmClient := NewVirtualMachineService(icsConnection.Client)

    vm, err := vmClient.VM("test")
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