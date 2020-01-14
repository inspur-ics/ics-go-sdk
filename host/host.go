package dc

import (
    "context"
    "github.com/inspur-ics/ics-go-sdk/client/methods"
    "github.com/inspur-ics/ics-go-sdk/client/types"
    //    "k8s.io/klog"
)

func (h *HostService) GetHost(ctx context.Context, hostUUID string) (*types.Host, error) {
    host, err := methods.GetHostById(ctx, h.RestAPITripper, hostUUID)
    return host, err
}

func (h *HostService) GetHostListByDC(ctx context.Context, datacenterPath string) ([]*types.Host, error) {
    var hostList []*types.Host
    //hosts, err := methods.GetAllDatacenterList(ctx, h.RestAPITripper)
    //for _, host := range hosts.Items {
    //    hostList = append(hostList, &host)
    //}
    return hostList, nil
}

//func  GetHost (hctx context.Context, connection *icsgo.ICSConnection, hostUUID string) (*types.Host, error){
//    ctx := context.Background()
//    err := connection.Connect(ctx)
//    if err != nil {
//        fmt.Println("Create ics connection error!")
//    }
//
//    Client := NewHostService(connection.Client)
//    host, err := Client.GetHostSv(hostUUID)
//
//    return host, err
//}
