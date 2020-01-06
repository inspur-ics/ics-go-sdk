package icsgo

import (
    "context"
    "fmt"
    "testing"
)

func TestConnect(t *testing.T) {
    icsConnection := &ICSConnection{
        Username:    "admin",
        Password:    "admin@inspur",
        Hostname:    "192.168.220.164",
        Port:        "443",
        Insecure:    true,
    }

    ctx := context.Background()

    err := icsConnection.Connect(ctx)
    if err != nil {
        t.Fatal("Create ics connection error!")
    } else {
        fmt.Println("Create ics connection success!")
    }
    icsConnection.Connect(ctx)
}

func TestNewClient(t *testing.T) {
    icsConnection := &ICSConnection{
        Username:    "admin",
        Password:    "admin@inspur",
        Hostname:    "192.168.220.164",
        Port:        "443",
        Insecure:    true,
    }

    ctx := context.Background()

    client, err := icsConnection.NewClient(ctx)
    if err != nil {
        t.Fatal("Create ics connection error!")
    }
    fmt.Println(fmt.Sprintf("Client token %s", client.Authorization))
}
