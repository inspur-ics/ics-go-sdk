package dc

import (
    "context"
    "fmt"
    icsgo "github.com/inspur-ics/ics-go-sdk"
    "github.com/inspur-ics/ics-go-sdk/client/methods"
    "github.com/inspur-ics/ics-go-sdk/client/types"
//    "k8s.io/klog"
)

func  (d *HostService) GetHostSv(hostUUID string) (*types.Host, error) {
    ctx := context.Background()
    host, err := methods.GetHostById(ctx, d.RestAPITripper, hostUUID)
    return host, err
}

func  GetHost (hctx context.Context, connection *icsgo.ICSConnection, hostUUID string) (*types.Host, error){
    ctx := context.Background()
    err := connection.Connect(ctx)
    if err != nil {
        fmt.Println("Create ics connection error!")
    }

    Client := NewHostService(connection.Client)
    host, err := Client.GetHostSv(hostUUID)

    return host, err
}
