package types
type Volume struct {
	ID                 string        `json:"id"`
	UUID               string        `json:"uuid"`
	Size               float64       `json:"size"`
	RealSize           float64       `json:"realSize"`
	Name               string        `json:"name"`
	FileName           string        `json:"fileName"`
	Offset             int           `json:"offset"`
	Shared             bool          `json:"shared"`
	DeleteModel        string        `json:"deleteModel"`
	VolumePolicy       string        `json:"volumePolicy"`
	Format             string        `json:"format"`
	BlockDeviceID      interface{}   `json:"blockDeviceId"`
	DiskType           interface{}   `json:"diskType"`
	DataStoreID        string        `json:"dataStoreId"`
	DataStoreName      string        `json:"dataStoreName"`
	DataStoreSize      float64       `json:"dataStoreSize"`
	FreeStorage        float64       `json:"freeStorage"`
	DataStoreType      string        `json:"dataStoreType"`
	VMName             interface{}   `json:"vmName"`
	VMStatus           interface{}   `json:"vmStatus"`
	Type               interface{}   `json:"type"`
	Description        interface{}   `json:"description"`
	Bootable           bool          `json:"bootable"`
	VolumeStatus       string        `json:"volumeStatus"`
	MountedHostIds     interface{}   `json:"mountedHostIds"`
	Md5                interface{}   `json:"md5"`
	DataSize           int           `json:"dataSize"`
	OpenStackID        interface{}   `json:"openStackId"`
	VvSourceDto        interface{}   `json:"vvSourceDto"`
	FormatDisk         bool          `json:"formatDisk"`
	ToBeConverted      bool          `json:"toBeConverted"`
	RelatedVms         interface{}   `json:"relatedVms"`
	ClusterSize        int           `json:"clusterSize"`
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
}
type VolumeListRsp struct {
	Items []Volume `json:"items"`
}