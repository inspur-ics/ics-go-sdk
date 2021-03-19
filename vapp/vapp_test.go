package vapp

import (
	"context"
	"encoding/json"
	//"fmt"
	icsgo "github.com/inspur-ics/ics-go-sdk"
	"github.com/inspur-ics/ics-go-sdk/client/types"
	"testing"
)

var (
	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "admin@inspur",
		Hostname: "10.7.11.90",
		Port:     "443",
		Insecure: true,
	}
)

func TestGetVappList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}
	vappClient := NewVappService(icsConnection.Client)
	vappList, err := vappClient.GetVappList(ctx)
	if err == nil {
		for i := range vappList {
			vappJson, _ := json.MarshalIndent(vappList[i], "", "\t")
			t.Logf("VappInfo: %v\n", string(vappJson))
		}
	} else {
		t.Errorf("Failed to get vapp list. Error: %v\n", err.Error())
	}
}

func TestGetVappByName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappName := "wyc"
	vappClient := NewVappService(icsConnection.Client)
	vappInfo, err := vappClient.GetVappByName(ctx, vappName)
	if vappInfo != nil {
		vappJson, _ := json.MarshalIndent(vappInfo, "", "\t")
		t.Logf("VappInfo: %v\n", string(vappJson))
	} else {
		t.Errorf("Failed to get vapp info by name. Error: %v\n", err.Error())
	}
}

func TestCreateVapp(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappReq := types.VappCreateReq{
		Name:         "test_vapp",
		Description:  "VAPP001",
		DataCenterID: "3f0094542ebb11eaa2691a130ac12531",
	}
	vappClient := NewVappService(icsConnection.Client)
	task, err := vappClient.CreateVapp(ctx, vappReq)
	if err != nil {
		t.Errorf("Failed to create vapp. Error: %v\n", err)
	} else {

		taskJson, _ := json.MarshalIndent(task, "", "\t")
		t.Logf("taskInfo: %v\n", string(taskJson))
	}
}

func TestDeleteVapp(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vappID := "8a878bda781f145e017849e3a1b701e6"
	vappClient := NewVappService(icsConnection.Client)
	task, err := vappClient.DeleteVapp(ctx, vappID)
	if err != nil {
		t.Errorf("Failed to delete vapp. Error: %v\n", err)
	} else {

		taskJson, _ := json.MarshalIndent(task, "", "\t")
		t.Logf("taskInfo: %v\n", string(taskJson))
	}
}
