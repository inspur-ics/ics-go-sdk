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

	dataCenterSvc := NewDatacenterService(icsConnection.Client)

	datacenterList, err := dataCenterSvc.GetAllDatacenters(ctx)
	if datacenterList != nil {
		fmt.Printf("%+v\n", datacenterList[0])
		fmt.Println(datacenterList)
	}
}

func TestGetDatacenterByName(t *testing.T) {
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

	dcName := "datacenter02"
	dcClient := NewDatacenterService(icsConnection.Client)
	dcInfo, err := dcClient.GetDatacenterByName(ctx, dcName)
	if err != nil {
		t.Errorf("Failed to get datacenter info by name. Error: %v", err)
	} else {
		dcJson, _ := json.MarshalIndent(dcInfo, "", "\t")
		t.Logf("DatacenterInfo: %v", string(dcJson))
	}
}
