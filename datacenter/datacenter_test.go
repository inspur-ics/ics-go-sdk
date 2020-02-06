package datacenter

import (
    "context"
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