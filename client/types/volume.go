package types

type Volume struct {
	ID                 string      `json:"id"`
	UUID               string      `json:"uuid"`
	Size               float64     `json:"size"`
	RealSize           float64     `json:"realSize"`
	Name               string      `json:"name"`
	FileName           string      `json:"fileName"`
	Offset             int         `json:"offset"`
	Shared             bool        `json:"shared"`
	DeleteModel        string      `json:"deleteModel"`
	VolumePolicy       string      `json:"volumePolicy"`
	Format             string      `json:"format"`
	BlockDeviceID      string      `json:"blockDeviceId,omitempty"`
	DiskType           string      `json:"diskType,omitempty"`
	DataStoreID        string      `json:"dataStoreId"`
	DataStoreName      string      `json:"dataStoreName"`
	DataStoreSize      float64     `json:"dataStoreSize"`
	FreeStorage        float64     `json:"freeStorage"`
	DataStoreType      string      `json:"dataStoreType"`
	DataStoreReplicate int         `json:"dataStoreReplicate"`
	VMName             string      `json:"vmName,omitempty"`
	VMStatus           string      `json:"vmStatus,omitempty"`
	Type               string      `json:"type,omitempty"`
	Description        string      `json:"description,omitempty"`
	Bootable           bool        `json:"bootable"`
	VolumeStatus       string      `json:"volumeStatus"`
	MountedHostIds     []string    `json:"mountedHostIds"`
	Md5                string      `json:"md5,omitempty"`
	DataSize           int         `json:"dataSize"`
	OpenStackID        string      `json:"openStackId,omitempty"`
	VvSourceDto        interface{} `json:"vvSourceDto"`
	FormatDisk         bool        `json:"formatDisk"`
	ToBeConverted      bool        `json:"toBeConverted"`
	RelatedVms         interface{} `json:"relatedVms"`
	ClusterSize        int         `json:"clusterSize"`
	ScsiID             string      `json:"scsiId"`
	Usage              string      `json:"usage"`
}

type VolumeReq struct {
	Name          string `json:"name"`
	Size          string `json:"size"`
	DataStoreType string `json:"dataStoreType"`
	DataStoreId   string `json:"dataStoreId"`
	VolumePolicy  string `json:"volumePolicy"`
	Description   string `json:"description"`
	Bootable      bool   `json:"bootable"`
	Shared        bool   `json:"shared"`
	Format        string `json:"format"`
	Usage         string `json:"usage"`
}

type VolumeListRsp struct {
	Items []Volume `json:"items"`
}
