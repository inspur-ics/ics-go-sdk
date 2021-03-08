package vm

import (
	"context"
	"encoding/json"
	"fmt"
	icsgo "github.com/inspur-ics/ics-go-sdk"
	"github.com/inspur-ics/ics-go-sdk/client/types"
	"testing"
)

var (
	icsConnection = &icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.48.50.13",
		Port:     "443",
		Insecure: true,
	}
)

func TestVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	vm, err := vmClient.GetVM(ctx, "8a878bda6f7012c7016f7de1935d047a")
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

func TestGetVMByUUID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)

	vm, err := vmClient.GetVMByUUID(ctx, "6c535536-6fe0-46ee-a20f-772dc0cd61ff")
	if vm != nil {
		fmt.Printf("%+v", vm)
	} else {
		fmt.Println("No VM be found by you point id.")
	}
}

func TestGetVMByIP(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)

	vm, err := vmClient.GetVMByIP(ctx, "10.7.11.81")
	if vm != nil {
		fmt.Printf("%+v", vm)
	} else {
		fmt.Println("No VM be found by you point id.")
	}
}

func TestGetVMByName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)

	vm, err := vmClient.GetVMByName(ctx, "master")
	if vm != nil {
		fmt.Printf("%+v", vm)
	} else {
		fmt.Println("No VM be found by you point id.")
	}
}

func TestVMPageList(t *testing.T) {
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

func TestGetVMTemplateList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	vmtList, err := vmClient.GetVMTemplateList(ctx)
	if err != nil {
		t.Errorf("Failed to get vm template list. Error: %v", err)
	} else {
		t.Logf("VMTemplateList: %+v", vmtList)
	}
}

func TestGetVMTemplate(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmtID := "8ab0b28d77be994a0177fc16f5b20ab0"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	vmt, err := vmClient.GetVMTemplate(ctx, vmtID)
	if err != nil {
		t.Errorf("Failed to get vm template by id. Error: %v", err)
	} else {
		vmtJson, _ := json.MarshalIndent(vmt, "", "\t")
		t.Logf("VMTemplate: %v", string(vmtJson))
	}
}

func TestGetVMTemplateByUUID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	uuid := "b9a6059a-5057-4efa-b501-78e2cb8d3e64"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	vmt, err := vmClient.GetVMTemplateByUUID(ctx, uuid)
	if err != nil {
		t.Errorf("Failed to get vm template by uuid. Error: %v", err)
	} else {
		vmtJson, _ := json.MarshalIndent(vmt, "", "\t")
		t.Logf("VMTemplate: %v", string(vmtJson))
	}
}

func TestGetVMTemplateByName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmtName := "template_test"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	vmt, err := vmClient.GetVMTemplateByName(ctx, vmtName)
	if err != nil {
		t.Errorf("Failed to get vm template by name. Error: %v", err)
	} else {
		vmtJson, _ := json.MarshalIndent(vmt, "", "\t")
		t.Logf("VMTemplate: %v", string(vmtJson))
	}
}

func TestCreateVMByTemplate(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmtName := "template_test"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	vmt, err := vmClient.GetVMTemplateByName(ctx, vmtName)
	if err != nil {
		t.Fatalf("Failed to get vm template by name. Error: %v", err)
	} else {
		vmtJson, _ := json.MarshalIndent(vmt, "", "\t")
		fmt.Printf("VMTemplate: %v\n", string(vmtJson))
	}

	quickClone := false
	vmt.Name = "vm_create_by_template_001"
	task, err := vmClient.CreateVMByTemplate(ctx, *vmt, quickClone)
	if err != nil {
		t.Fatalf("Failed to create vm by template. Error: %v", err)
	} else {
		fmt.Printf("Clone VM Task: %+v\n", task)
	}

	fmt.Printf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		fmt.Printf("Task Status: %v\n", string(taskJson))
	}

}
