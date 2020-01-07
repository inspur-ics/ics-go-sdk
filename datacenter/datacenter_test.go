package dc

import (
    "context"
    "fmt"
    icsgo "github.com/inspur-ics/ics-go-sdk"
    "testing"
)

func TestGetAllDatacenter(t *testing.T) {
    icsConnection := &icsgo.ICSConnection{
        Username: "admin",
        Password: "admin@inspur",
        Hostname: "10.7.11.90",
        Port:     "443",
        Insecure: true,
    }
    ctx := context.Background()

    dc, err := GetAllDatacenter (ctx, icsConnection)

    if err != nil {
        fmt.Println("func failed!!!")
    } else {
        for _, d := range dc {
            fmt.Println(d.ID)
        }
    }
}

func TestGetDatacenter(t *testing.T) {
    icsConnection := &icsgo.ICSConnection{
        Username: "admin",
        Password: "admin@inspur",
        Hostname: "10.7.11.90",
        Port:     "443",
        Insecure: true,
    }
    ctx := context.Background()

    dc, err := GetDatacenter (ctx, icsConnection, "3f0094542ebb11eaa2691a130ac12531")

    if err != nil {
        fmt.Println("func failed!!!")
    } else {
        fmt.Println(dc.Name)
        fmt.Println("!!!!90039")
    }
}

func TestGetNumberOfDatacenters(t *testing.T) {
    icsConnection := &icsgo.ICSConnection{
        Username: "admin",
        Password: "admin@inspur",
        Hostname: "10.7.11.90",
        Port:     "443",
        Insecure: true,
    }
    ctx := context.Background()

    num, err := GetNumberOfDatacenters (ctx, icsConnection)

    if err != nil {
        fmt.Println("func failed!!!")
    } else {
        fmt.Println(num)
        fmt.Println("ok!!!")
    }
}