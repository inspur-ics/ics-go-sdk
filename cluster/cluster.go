package cluster

import (
	"context"
	"github.com/inspur-ics/ics-go-sdk/client/methods"
	"github.com/inspur-ics/ics-go-sdk/client/types"
	//    "k8s.io/klog"
)

func (h *ClusterService) GetClusterList(ctx context.Context) ([]types.Cluster, error) {
	clusterlist, err := methods.GetClusterList(ctx, h.RestAPITripper)
	return clusterlist, err
}
