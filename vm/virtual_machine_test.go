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
		Hostname: "10.49.34.161",
		Port:     "443",
		Insecure: true,
	}
)

func TestGetVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	vm, err := vmClient.GetVM(ctx, "8ab0b28d77be994a017819e3f65f1169")
	if err != nil {
		t.Errorf("Failed to get vm info by specified id. Error: %v", err)
	} else {
		vmJson, _ := json.MarshalIndent(vm, "", "\t")
		t.Logf("VM Info: %s\n", string(vmJson))
	}
}

func TestSetVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	vm, err := vmClient.GetVM(ctx, "8a878bda6f6f3ca4016f6f6eb8d300bb")
	if err != nil {
		t.Fatalf("Failed to get vm info by specified id. Error: %v", err)
	}

	vm.Name = "11.82_zhanghuijian1"
	//vm.VncPasswd = "12345678"
	task, err := vmClient.SetVM(ctx, *vm)
	if err != nil {
		t.Fatalf("Failed to set vm. Error: %v", err)
	} else {
		t.Logf("Set VM Task: %+v\n", task)
	}

	taskInfo, err := vmClient.GetTaskInfo(task)
	if err != nil {
		t.Fatalf("Failed to get task info. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Info: %v\n", string(taskJson))
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err = vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestGetVMByUUID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)

	vm, err := vmClient.GetVMByUUID(ctx, "71909586-f83b-45de-b620-4650aceb7865")
	if err != nil || vm == nil {
		t.Errorf("Failed to get vm info by specified uuid. Error: %v", err)
	} else {
		vmJson, _ := json.MarshalIndent(vm, "", "\t")
		t.Logf("VM Info: %s\n", string(vmJson))
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
		fmt.Println("No VM be found by specified id.")
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
		fmt.Println("No VM be found by specified id.")
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

func TestPowerOnVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8ab0b28d77be994a0178110f28f110ce"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	task, err := vmClient.PowerOnVM(ctx, vmId)
	if err != nil {
		t.Fatalf("Failed to poweron vm. Error: %v", err)
	} else {
		t.Logf("PowerOn VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestPowerOffVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8ab0b28d77be994a0178110f28f110ce"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	task, err := vmClient.PowerOffVM(ctx, vmId)
	if err != nil {
		t.Fatalf("Failed to poweroff vm. Error: %v", err)
	} else {
		t.Logf("PowerOff VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestShutdownVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8ab0b28d77be994a0178110f28f110ce"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	task, err := vmClient.ShutdownVM(ctx, vmId)
	if err != nil {
		t.Fatalf("Failed to shutdown vm. Error: %v", err)
	} else {
		t.Logf("Shutdown VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestRestartVM(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8a878bda6f6f3ca4016f6f6eb8d300bb"
	vmClient := NewVirtualMachineService(icsConnection.Client)
	task, err := vmClient.RestartVM(ctx, vmId)
	if err != nil {
		t.Fatalf("Failed to restart vm. Error: %v", err)
	} else {
		t.Logf("Restart VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}

func TestDeleteVM(t *testing.T) {
	var task *types.Task
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmId := "8ab1a2218dc27ce1018dc6e0a218004f"
	deleteFile := true
	removeData := true

	vmClient := NewVirtualMachineService(icsConnection.Client)
	needAuth, err := vmClient.IsDeleteNeedIdentityAuth(ctx)
	if needAuth {
		t.Logf("Delete VM %s with check params", vmId)
		task, err = vmClient.DeleteVMWithCheckParams(ctx, vmId, deleteFile, removeData, icsConnection.Password)
	} else {
		task, err = vmClient.DeleteVM(ctx, vmId, deleteFile, removeData)
	}
	if err != nil {
		t.Fatalf("Failed to delete vm. Error: %v", err)
	} else {
		t.Logf("Delete VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
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
		t.Logf("VMTemplate: %v\n", string(vmtJson))
	}

	quickClone := false
	vmt.Name = "vm_create_by_template_001"
	task, err := vmClient.CreateVMByTemplate(ctx, *vmt, quickClone)
	if err != nil {
		t.Fatalf("Failed to create vm by template. Error: %v", err)
	} else {
		t.Logf("Clone VM Task: %+v\n", task)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := vmClient.TraceTaskProcess(task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}

}

func TestGetVMPowerStateByID(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	status, err := vmClient.GetVMPowerStateByID(ctx, "8a878bda781f145e0178456029a8016c")
	if status != nil {
		fmt.Println(*status)
	} else {
		fmt.Println("No VM be found by you point id.")
	}
}

func TestGetVMNetState(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmClient := NewVirtualMachineService(icsConnection.Client)
	nic, err := vmClient.GetVMNetState(ctx, "8a878bda6f7012c7016f70b40ed000a1")
	//nic, err := vmClient.GetVMNetState(ctx, "8a878bda781f145e0178456029a8016c")

	if err != nil {
		t.Fatalf("No VM be found by you point id.")
	} else {
		for _, data := range nic {
			fmt.Println("-----------------------")
			fmt.Println(data)
			fmt.Println("-----------------------")
		}
	}
}
