package storage

import (
    "context"
    "github.com/inspur-ics/ics-go-sdk/client/methods"
    "github.com/inspur-ics/ics-go-sdk/client/types"
)

//FIXME TODO.WANGYONGCHAO
func (sts *StorageService)GetAllDatastores(ctx context.Context, datacenterId string) ([]*types.Storage, error) {
    var storages []*types.Storage
    var err error

    return storages, err
}

func (sts *StorageService)GetStoragesList(ctx context.Context) ([]types.Storage, error) {
    storagelist, err := methods.GetStorageList(ctx, sts.RestAPITripper)
    return storagelist.Items, err
}

//FIXME TODO.WANGYONGCHAO
func (sts *StorageService)GetDatastoreByName(ctx context.Context, name string) (*types.Storage, error) {
    var storage *types.Storage
    var err error

    return storage, err
}

func (sts *StorageService)GetStorageInfoById(ctx context.Context, id string) (*types.Storage, error) {
    storage, err := methods.GetStorageInfo(ctx, sts.RestAPITripper, id)
    return storage, err
}
func (v *StorageService) GetStoragePageList(req *types.StoragePageReq) (*types.StoragePageResponse, error) {
    ctx := context.Background()
    storagePages, err := methods.GetStoragePageList(ctx, v.RestAPITripper, req)
    return storagePages, err
}