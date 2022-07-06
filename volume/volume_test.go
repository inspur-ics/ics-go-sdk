package volume

import (
	"context"
	"encoding/json"
	"fmt"
	icsgo "github.com/inspur-ics/ics-go-sdk"
	"github.com/inspur-ics/ics-go-sdk/client/types"
	"testing"
	"time"
)

var volumeClient *VolumeService
var icsConnection icsgo.ICSConnection
var ctx context.Context
var volumeid string

func TestCreateVolume(t *testing.T) {
	fmt.Println("********************CreateVolume**************")
	icsConnection = icsgo.ICSConnection{
		//Username: "admin",
		//Password: "Cloud@s1",
		AccessKeyID:     "09cb1VN5xu5T7469d1Yo",
		AccessKeySecret: "6e4711WY1pY310qPG8ntv7RdIuM8vDO07XKM2GeX",
		Hostname:        "10.49.34.107",
		Port:            "443",
		Insecure:        true,
	}
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}
	createVolumeReq := types.VolumeReq{
		Name:          "test2",
		Size:          "20",
		DataStoreId:   "8ab1a2eb81989a090181998f6ffb0005",
		DataStoreType: "LOCAL",
		VolumePolicy:  "THIN",
		Format:        "RAW",
		Description:   "k8s",
		Usage:         "POD",
		Bootable:      false,
		Shared:        false,
	}
	volumeClient = NewVolumeService(icsConnection.Client)
	task, err := volumeClient.CreateVolume(ctx, createVolumeReq)
	fmt.Println(task.TaskId)
	time.Sleep(5 * time.Second)
}

func TestGetVolumesInDatastore(t *testing.T) {
	icsConnection = icsgo.ICSConnection{
		//Username: "admin",
		//Password: "Cloud@s1",
		AccessKeyID:     "09cb1VN5xu5T7469d1Yo",
		AccessKeySecret: "6e4711WY1pY310qPG8ntv7RdIuM8vDO07XKM2GeX",
		Hostname:        "10.49.34.107",
		Port:            "443",
		Insecure:        true,
	}
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	if icsConnection.AccessKeyID == "" {
		fmt.Println("access key id emtpy")
	} else {
		fmt.Printf("access key id %s\n", icsConnection.AccessKeyID)
	}

	volumeClient = NewVolumeService(icsConnection.Client)
	fmt.Println("********************GetVolumesInDatastore**************")
	volumes, err := volumeClient.GetVolumesInDatastore(ctx, "8ab1a2eb81989a090181998f6ffb0005")
	if volumes != nil {
		for i := 0; i < len(volumes); {
			volumeinfo := volumes[i]
			if volumeinfo.Name == "test2" {
				volumeid = volumeinfo.ID
				fmt.Printf("volume %s info: %+v \n", volumeinfo.Name, volumeinfo)
			} else {
				fmt.Printf("volume id:%s  name:%s \n", volumeinfo.ID, volumeinfo.Name)
			}
			i++
		}
	} else {
		fmt.Println(err.Error())
	}
}

func TestGetVolumeInfoById(t *testing.T) {
	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.49.34.22",
		Port:     "443",
		Insecure: true,
	}
	ctx = context.Background()
	icsConnection.Connect(ctx)
	volumeClient = NewVolumeService(icsConnection.Client)
	fmt.Println("********************TestGetVolumeInfoById**************")
	volumeinfo, err := volumeClient.GetVolumeInfoById(ctx, "8ab0b28d7867fb1d0178687e6c13006a")
	if err == nil {
		fmt.Println(volumeinfo.Name)
	} else {
		fmt.Println(err.Error())
	}
}

func TestDeleteVolume(t *testing.T) {
	fmt.Println("********************TestDeleteVolume**************")
	if volumeid != "" {
		task, err := volumeClient.DeleteVolume(ctx, volumeid, true)
		if err == nil {
			fmt.Println(task.TaskId)
		} else {
			fmt.Println(err.Error())
		}
	}
}

func TestSetVolume(t *testing.T) {
	fmt.Println("********************TestSetVolume**************")

	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.48.50.13",
		Port:     "443",
		Insecure: true,
	}
	ctx = context.Background()
	icsConnection.Connect(ctx)
	volumeId := "8ab0b28d7867fb1d0178687e6c13006a"
	volumeClient = NewVolumeService(icsConnection.Client)
	volumeInfo, err := volumeClient.GetVolumeInfoById(ctx, volumeId)
	if err != nil {
		t.Fatalf("Failed to get volume info. Error: %v", err)
	}

	volumeInfo.Size += 5
	task, err := volumeClient.SetVolume(ctx, volumeId, volumeInfo)
	if err != nil {
		t.Fatalf("Failed to set volume. Error: %v", err)
	}

	t.Logf("Waiting task %v finish.....\n", task.TaskId)
	taskInfo, err := volumeClient.TraceTaskProcess(&task)
	if err != nil {
		t.Fatalf("Failed to trace task. Error: %v", err)
	} else {
		taskJson, _ := json.MarshalIndent(taskInfo, "", "\t")
		t.Logf("Task Status: %v\n", string(taskJson))
	}
}
