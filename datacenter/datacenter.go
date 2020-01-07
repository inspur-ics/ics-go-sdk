package dc

import (
    "context"
    "fmt"
    icsgo "github.com/inspur-ics/ics-go-sdk"
    "github.com/inspur-ics/ics-go-sdk/client/methods"
    "github.com/inspur-ics/ics-go-sdk/client/types"
)

func  (d *DatacenterService) GetDatacenterSv(datacenterid string) (*types.Datacenter, error) {
    ctx := context.Background()
    dc, err := methods.GetDatacenterById(ctx, d.RestAPITripper, datacenterid)
    return dc, err
}

func (d *DatacenterService) GetAllDatacenterSv() ([]*types.Datacenter, error) {

//    var datacenterlist types.DatacenterPageResponse
    var dclist []*types.Datacenter

    ctx := context.Background()
    datacenters, err := methods.GetAllDatacenterList(ctx, d.RestAPITripper)

    for _, datacenter := range datacenters.Items {
        dclist = append(dclist, &datacenter)
    }

    return dclist, err
}

func  GetAllDatacenter (dctx context.Context, connection *icsgo.ICSConnection) ([]*types.Datacenter, error){
    ctx := context.Background()
    err := connection.Connect(ctx)
    if err != nil {
        fmt.Println("Create ics connection error!")
    }

    Client := NewDatacenterService(connection.Client)
    dclist, err := Client.GetAllDatacenterSv()

    return dclist, err
}

// GetNumberOfDatacenters returns the number of DataCenters in this iCenter
func  GetDatacenter (dctx context.Context, connection *icsgo.ICSConnection, dcid string) (*types.Datacenter, error){
    ctx := context.Background()
    err := connection.Connect(ctx)
    if err != nil {
        fmt.Println("Create ics connection error!")
    }

    Client := NewDatacenterService(connection.Client)
    dc, err := Client.GetDatacenterSv(dcid)

    return dc, err
}

// GetNumberOfDatacenters returns the number of DataCenters in this iCenter
func GetNumberOfDatacenters(dctx context.Context, connection *icsgo.ICSConnection) (int, error) {
    dc, err := GetAllDatacenter (dctx, connection)
    if err != nil {
        return 0, err
    }
    return len(dc), nil
}
