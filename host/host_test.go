package host

import (
	"context"
	"fmt"
	icsgo "github.com/inspur-ics/ics-go-sdk"
	"testing"
)

var hostClient *HostService
var icsConnection icsgo.ICSConnection
var ctx context.Context
var hostid string
func TestHostList(t *testing.T) {
	fmt.Println("********************TestHostList**************")
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

	hostClient = NewHostService(icsConnection.Client)
	hostlist, err := hostClient.GetHostList(ctx)
	if hostlist != nil {
		for i := 0; i < len(hostlist); {
			fmt.Println(hostlist[i].Name)
			fmt.Println(hostlist[i].IP)
			hostid = hostlist[i].ID
			i++
		}
	} else {
		fmt.Println(err.Error())
	}
}
func TestGetHostAccessibleDatastoreList(t *testing.T) {
	fmt.Println("********************GetHostAccessibleDatastoreList**************")
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	hostClient = NewHostService(icsConnection.Client)
	storagelist, err := hostClient.GetHostAccessibleDatastoreList(ctx,hostid)
	if storagelist != nil {
		for i := 0; i < len(storagelist); {
			fmt.Println(storagelist[i].Name)
			i++
		}
	} else {
		fmt.Println(err.Error())
	}
}

func TestGetHostListByDC(t *testing.T) {
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

	hostClient = NewHostService(icsConnection.Client)
	hostlist, err := hostClient.GetHostListByDC(ctx, "3f0094542ebb11eaa2691a130ac12531")
	if hostlist != nil && len(hostlist) >= 1{
		fmt.Printf("%+v", hostlist[0])
	} else if err != nil {
		fmt.Println(err.Error())
	}
}