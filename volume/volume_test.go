package volume

import (
    "context"
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
		Username: "admin",
		Password: "admin@inspur",
		Hostname: "10.7.11.90",
		Port:     "443",
		Insecure: true,
	}
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}
	createVolumeReq := types.VolumeReq{
		Name:          "test",
		Size:          "10",
		DataStoreId:   "8a878bda6f6f3ca4016f6f458b50003c",
		DataStoreType: "LOCAL",
		VolumePolicy:  "THIN",
		Description:   "k8s",
		Bootable:      false,
		Shared:        false,
	}
	volumeClient = NewVolumeService(icsConnection.Client)
	task, err := volumeClient.CreateVolume(ctx, createVolumeReq)
	fmt.Println(task.TaskId)
    time.Sleep(5*time.Second)
}

func TestGetVolumesInDatastore(t *testing.T) {
	fmt.Println("********************GetVolumesInDatastore**************")
	volumes, err := volumeClient.GetVolumesInDatastore(ctx, "8a878bda6f6f3ca4016f6f458b50003c")
	if volumes != nil {
		for i := 0; i <len(volumes); {
			volumeinfo := volumes[i]
			if volumeinfo.Name == "test" {
				volumeid = volumeinfo.ID
				fmt.Println("test volume id :"+volumeid)
			}
			i++
		}
	} else {
		fmt.Println(err.Error())
	}
}
func TestGetVolumeInfoById(t *testing.T) {
	fmt.Println("********************TestGetVolumeInfoById**************")
	volumeinfo, err := volumeClient.GetVolumeInfoById(ctx, volumeid)
	if err == nil {

		fmt.Println(volumeinfo.Name)
	} else {
		fmt.Println(err.Error())
	}
}
func TestDeleteVolume(t *testing.T) {
	fmt.Println("********************TestDeleteVolume**************")
    if volumeid!= "" {
        task, err := volumeClient.DeleteVolume(ctx, volumeid,true)
        if err == nil {
            fmt.Println(task.TaskId)
        } else {
            fmt.Println(err.Error())
        }
    }
}
