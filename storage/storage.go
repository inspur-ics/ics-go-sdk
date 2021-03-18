package storage

import (
	"context"
	"github.com/inspur-ics/ics-go-sdk/client/methods"
	"github.com/inspur-ics/ics-go-sdk/client/types"
)

//FIXME TODO.WANGYONGCHAO
func (sts *StorageService) GetAllDatastores(ctx context.Context, datacenterId string) ([]*types.Storage, error) {
	var storages []*types.Storage
	var err error

	return storages, err
}

func (sts *StorageService) GetStoragesList(ctx context.Context) ([]types.Storage, error) {
	storagelist, err := methods.GetStorageList(ctx, sts.RestAPITripper)
	return storagelist.Items, err
}

func (sts *StorageService) GetStorageInfoByName(ctx context.Context, name string) (*types.Storage, error) {
	storageList, err := methods.GetStorageList(ctx, sts.RestAPITripper)
	if storageList == nil || len(storageList.Items) == 0 {
		return nil, err
	}
	for _, storage := range storageList.Items {
		if storage.Name == name {
			return &storage, nil
		}
	}
	return nil, err
}

func (sts *StorageService) GetStorageInfoById(ctx context.Context, id string) (*types.Storage, error) {
	storage, err := methods.GetStorageInfo(ctx, sts.RestAPITripper, id)
	return storage, err
}

func (sts *StorageService) GetStoragePageList(req *types.StoragePageReq) (*types.StoragePageResponse, error) {
	ctx := context.Background()
	storagePages, err := methods.GetStoragePageList(ctx, sts.RestAPITripper, req)
	return storagePages, err
}
