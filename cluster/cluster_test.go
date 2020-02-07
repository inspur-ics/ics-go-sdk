package cluster
import (
	"context"
	"fmt"
	icsgo "github.com/inspur-ics/ics-go-sdk"
	"testing"
)

var clusterClient *ClusterService
var icsConnection icsgo.ICSConnection
var ctx context.Context
func TestClusterList(t *testing.T) {
	fmt.Println("********************TestClusterList**************")
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
	clusterClient=NewClusterService(icsConnection.Client)
	clusterlist, err := clusterClient.GetClusterList(ctx)
	if err == nil {
		for i := 0; i < len(clusterlist); {
			fmt.Println(clusterlist[i].Id)
			fmt.Println(clusterlist[i].Name)
			i++
		}
	} else {
		fmt.Println(err.Error())
	}
}