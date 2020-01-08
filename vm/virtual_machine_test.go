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

    vm, err := vmClient.VM("8a878bda6f7012c7016f7de1935d047a")
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

func TestGetHostByVMId(t *testing.T) {
    icsConnection := &icsgo.ICSConnection{
        Username: "admin",
        Password: "admin@inspur",
        Hostname: "10.7.11.90",
        Port:     "443",
        Insecure: true,
    }
    ctx := context.Background()

    host, err := GetHostByVMUUId(ctx, icsConnection, "cf35e100-9fff-49f1-b5c4-b06940c89d13")

    if err != nil {
        fmt.Println("func failed!!!")
    } else {
        fmt.Println(host.Name)
        fmt.Println(host.HostName)
        fmt.Println(host.CPUSocket)
        fmt.Println("ok!!!")
    }
}

func TestGetHostByVMName(t *testing.T) {
    icsConnection := &icsgo.ICSConnection{
        Username: "admin",
        Password: "admin@inspur",
        Hostname: "10.7.11.90",
        Port:     "443",
        Insecure: true,
    }
    ctx := context.Background()

    host, err := GetHostByVMName(ctx, icsConnection, "11.92_guozhen")

    if err != nil {
        fmt.Println("func failed!!!")
    } else {
        fmt.Println(host.Name)
        fmt.Println(host.HostName)
        fmt.Println(host.CPUSocket)
        fmt.Println("ok!!!")
    }
}