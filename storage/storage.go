package storage

import (
    "context"
    "github.com/inspur-ics/ics-go-sdk/client/types"
)

//FIXME TODO.WANGYONGCHAO
func (sts *StorageService)GetAllDatastores(ctx context.Context, datacenterId string) ([]*types.Datastore, error) {
    var storages []*types.Datastore
    var err error

    return storages, err
}

//FIXME TODO.WANGYONGCHAO
func (sts *StorageService)GetDatastoreByName(ctx context.Context, name string) (*types.Datastore, error) {
    var storage  *types.Datastore
    var err error

    return storage, err
}