package host

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
	var hostslist []*types.Host
	hosts, err := methods.GetHostListByDC(ctx, h.RestAPITripper, datacenterPath)
	if err == nil && len(hosts.Items) > 0 {
		for index := range hosts.Items {
			hostslist = append(hostslist, &hosts.Items[index])
		}
	}
	return hostslist, err
}

func (hostserver *HostService) GetHostListByClusterID(ctx context.Context, clusterID string) ([]*types.Host, error) {
	var hostslist []*types.Host
	hostResp, err := methods.GetHostList(ctx, hostserver.RestAPITripper)
	if err != nil {
		return hostslist, err
	}
	for i := range hostResp.Items {
		if hostResp.Items[i].ClusterID == clusterID {
			hostslist = append(hostslist, &hostResp.Items[i])
		}
	}
	return hostslist, err
}

func (hostserver *HostService) GetHostListByClusterName(ctx context.Context, clusterName string) ([]*types.Host, error) {
	var hostslist []*types.Host
	hostResp, err := methods.GetHostList(ctx, hostserver.RestAPITripper)
	if err != nil {
		return hostslist, err
	}
	for i := range hostResp.Items {
		if hostResp.Items[i].ClusterName == clusterName {
			hostslist = append(hostslist, &hostResp.Items[i])
		}
	}
	return hostslist, err
}

func (h *HostService) GetHostAvailStorages(ctx context.Context, hostUUID string) ([]*types.Storage, error) {
	var storageList []*types.Storage
	storages, err := methods.GetHostAvailStorages(ctx, h.RestAPITripper, hostUUID)

	for idx := range storages {
		storageList = append(storageList, &storages[idx])
	}

	return storageList, err
}

func (hostserver *HostService) GetHostList(ctx context.Context) ([]types.Host, error) {
	hostlist, err := methods.GetHostList(ctx, hostserver.RestAPITripper)
	return hostlist.Items, err
}
func (hostserver *HostService) GetHostAccessibleDatastoreList(ctx context.Context, hostid string) ([]types.Storage, error) {
	storagelist, err := methods.GetHostAccessibleDatastoreList(ctx, hostserver.RestAPITripper, hostid)
	return storagelist, err
}
