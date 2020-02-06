package volume

import (
    "context"
    "github.com/inspur-ics/ics-go-sdk/client/methods"
    "github.com/inspur-ics/ics-go-sdk/client/types"
)

func (volume *VolumeService)CreateVolume(ctx context.Context,volumereq types.VolumeReq) (types.Task, error) {
    task, err := methods.CreateVolume(ctx, volume.RestAPITripper,volumereq)
    return task , err
}

func (vs *VolumeService)DeleteVolume(ctx context.Context, volumeid string,deleteVolume bool) (types.Task, error) {
    task, err := methods.DeleteVolume(ctx, vs.RestAPITripper, volumeid,deleteVolume)
    return task, err
}
func (vs *VolumeService) GetVolumeInfoById(ctx context.Context, volumeid string) (types.Volume, error) {
    volumeinfo, err := methods.GetVolumeInfoById(ctx, vs.RestAPITripper, volumeid)
    return volumeinfo, err
}
func (vs *VolumeService) GetVolumesInDatastore(ctx context.Context, storeid string) ([]types.Volume, error) {
    volumes, err := methods.GetVolumesInDatastore(ctx, vs.RestAPITripper, storeid)
    return volumes.Items, err
}