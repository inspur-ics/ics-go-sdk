package datacenter

import (
	"context"
	"encoding/json"
	"fmt"
	icsgo "github.com/inspur-ics/ics-go-sdk"
	"testing"
)

func TestGetAllDatacenters(t *testing.T) {
	icsConnection := &icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.49.34.161",
		Port:     "443",
		Insecure: true,
	}
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	dataCenterSvc := NewDatacenterService(icsConnection.Client)

	datacenterList, err := dataCenterSvc.GetAllDatacenters(ctx)
	if err != nil {
		t.Errorf("Failed to get all datacenters. Error: %v", err)
		fmt.Printf("%+v\n", datacenterList[0])
		fmt.Println(datacenterList)
	} else {
		for _, datacenter := range datacenterList {
			dcJson, _ := json.MarshalIndent(datacenter, "", "\t")
			t.Logf("DatacenterInfo: %v", string(dcJson))
		}
	}
}

func TestGetDatacenterByName(t *testing.T) {
	icsConnection := &icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.49.34.22",
		Port:     "443",
		Insecure: true,
	}
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	dcName := "Datacenter"
	dcClient := NewDatacenterService(icsConnection.Client)
	dcInfo, err := dcClient.GetDatacenterByName(ctx, dcName)
	if err != nil {
		t.Errorf("Failed to get datacenter info by name. Error: %v", err)
	} else {
		dcJson, _ := json.MarshalIndent(dcInfo, "", "\t")
		t.Logf("DatacenterInfo: %v", string(dcJson))
	}
}
