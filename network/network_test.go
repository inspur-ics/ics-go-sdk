package network

import (
	"context"
	"encoding/json"
	icsgo "github.com/inspur-ics/ics-go-sdk"
	//"github.com/inspur-ics/ics-go-sdk/client/types"
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

func TestGetNetworkByName(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	networkName := "manageNetwork0"
	networkClient := NewNetworkService(icsConnection.Client)
	network, err := networkClient.GetNetworkByName(ctx, networkName)
	if err != nil {
		t.Errorf("Failed to get network by name %s. Error: %v\n", networkName, err)
	} else {
		networkJson, _ := json.MarshalIndent(network, "", "\t")
		t.Logf("Network Info: %s\n", string(networkJson))
	}
}

func TestGetNetworkList(t *testing.T) {
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	networkClient := NewNetworkService(icsConnection.Client)
	networkList, err := networkClient.GetNetworkList(ctx)
	if err != nil {
		t.Errorf("Failed to get network list. Error: %v\n", err)
	} else {
		for _, network := range networkList {
			networkJson, _ := json.MarshalIndent(network, "", "\t")
			t.Logf("Network Info: %s\n", string(networkJson))
		}
	}
}
