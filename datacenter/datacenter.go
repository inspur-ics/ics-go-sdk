package datacenter

import (
	"context"
	"github.com/inspur-ics/ics-go-sdk/client/methods"
	"github.com/inspur-ics/ics-go-sdk/client/types"
)

func (d *DatacenterService) GetDatacenter(ctx context.Context, datacenterid string) (*types.Datacenter, error) {
	dc, err := methods.GetDatacenterById(ctx, d.RestAPITripper, datacenterid)
	return dc, err
}

func (d *DatacenterService) GetAllDatacenters(ctx context.Context) ([]*types.Datacenter, error) {
	var dclist []*types.Datacenter
	datacenters, err := methods.GetAllDatacenterList(ctx, d.RestAPITripper)
	for idx := range datacenters.Items {
		dclist = append(dclist, &datacenters.Items[idx])
	}
	return dclist, err
}

//func  (d *DatacenterService) GetDatacenterSv(datacenterid string) (*types.Datacenter, error) {
//    ctx := context.Background()
//    dc, err := methods.GetDatacenterById(ctx, d.RestAPITripper, datacenterid)
//    return dc, err
//}
//
//func (d *DatacenterService) GetAllDatacenterSv() ([]*types.Datacenter, error) {
//
////    var datacenterlist types.DatacenterPageResponse
//    var dclist []*types.Datacenter
//
//    ctx := context.Background()
//    datacenters, err := methods.GetAllDatacenterList(ctx, d.RestAPITripper)
//
//    for _, datacenter := range datacenters.Items {
//        dclist = append(dclist, &datacenter)
//    }
//
//    return dclist, err
//}
//

func (d *DatacenterService) GetDatacenterVMList(ctx context.Context, datacenterid string) ([]*types.VirtualMachine, error) {

	var vmlist []*types.VirtualMachine
	dcvms, err := methods.GetDatacenterVMById(ctx, d.RestAPITripper, datacenterid)

	for idx := range dcvms.Items {
		vmlist = append(vmlist, &dcvms.Items[idx])
	}

	return vmlist, err
}

//func  GetAllDatacenter (dctx context.Context, connection *icsgo.ICSConnection) ([]*types.Datacenter, error){
//    ctx := context.Background()
//    err := connection.Connect(ctx)
//    if err != nil {
//        fmt.Println("Create ics connection error!")
//    }
//
//    Client := NewDatacenterService(connection.Client)
//    dclist, err := Client.GetAllDatacenterSv()
//
//    return dclist, err
//}
//
//// GetNumberOfDatacenters returns the number of DataCenters in this iCenter
//func  GetDatacenter (dctx context.Context, connection *icsgo.ICSConnection, dcid string) (*types.Datacenter, error){
//    ctx := context.Background()
//    err := connection.Connect(ctx)
//    if err != nil {
//        fmt.Println("Create ics connection error!")
//    }
//
//    Client := NewDatacenterService(connection.Client)
//    dc, err := Client.GetDatacenterSv(dcid)
//
//    return dc, err
//}
//
//// GetNumberOfDatacenters returns the number of DataCenters in this iCenter
//func GetNumberOfDatacenters(dctx context.Context, connection *icsgo.ICSConnection) (int, error) {
//    dc, err := GetAllDatacenter (dctx, connection)
//    if err != nil {
//        return 0, err
//    }
//    return len(dc), nil
//}
//
//// GetVMByIP gets the VM object from the given IP address
//func GetVMByIP(dctx context.Context, connection *icsgo.ICSConnection, dcUUID string, ipAddy string) (*types.VirtualMachine, error) {
//    var tvm types.VirtualMachine
//    ctx := context.Background()
//    err := connection.Connect(ctx)
//    if err != nil {
//        fmt.Println("Create ics connection error!")
//    }
//
//    Client := NewDatacenterService(connection.Client)
//    vmlist, err := Client.GetDatacenterVMSv(dcUUID)
//
//    //    fmt.Println(len(vmlist))
//    for _, vm := range vmlist {
//        for _, nic := range vm.Nics{
//            if ipAddy == nic.IP{
//                return &vm, err
//            }
//        }
//    }
//    return &tvm, err
//}
//
//// GetVMByDNSName is not implemented.
//func GetVMByDNSName(dctx context.Context, connection *icsgo.ICSConnection, dcUUID string, dnsName string) (*types.VirtualMachine, error) {
//    return nil, nil
//}
//
//// GetVMByUUID gets the VM object from the given vmUUID
//func GetVMByUUID(dctx context.Context, connection *icsgo.ICSConnection, dcUUID string, vmUUID string) (*types.VirtualMachine, error) {
//    var tvm types.VirtualMachine
//    ctx := context.Background()
//    err := connection.Connect(ctx)
//    if err != nil {
//        fmt.Println("Create ics connection error!")
//    }
//
//    Client := NewDatacenterService(connection.Client)
//    vmlist, err := Client.GetDatacenterVMSv(dcUUID)
//
////    fmt.Println(len(vmlist))
//    for _, vm := range vmlist {
//        if vmUUID == vm.UUID{
//            return &vm, err
//        }
//    }
//    return &tvm, err
//}
