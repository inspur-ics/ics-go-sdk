package types

type Storage struct {
	DataStoreType       string  `json:"dataStoreType"`
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	MountPath           string  `json:"mountPath"`
	Capacity            float64 `json:"capacity"`
	UsedCapacity        float64 `json:"usedCapacity"`
	AvailCapacity       float64 `json:"availCapacity"`
	DataCenterID        string  `json:"dataCenterId"`
	HostID              string  `json:"hostId"`
	MountStatus         string  `json:"mountStatus"`
	HostIP              string  `json:"hostIp"`
	UUID                string  `json:"uuid"`
	AbsolutePath        string  `json:"absolutePath"`
	DataCenterName      string  `json:"dataCenterName"`
	DataCenterOrHostDto struct {
		DataCenterOrHost string `json:"dataCenterOrHost"`
		DataCenterName   string `json:"dataCenterName"`
		HostName         string `json:"hostName"`
		Status           string `json:"status"`
	} `json:"dataCenterOrHostDto"`
	BlockDeviceDto    string `json:"blockDeviceDto"`
	DataCenterDto     string `json:"dataCenterDto"`
	HostNumbers       int    `json:"hostNumbers"`
	VMNumbers         int    `json:"vmNumbers"`
	VolumesNumbers    int    `json:"volumesNumbers"`
	VMTemplateNumbers int    `json:"vmTemplateNumbers"`
	Tags              string `json:"tags"`
	MaxSlots          int    `json:"maxSlots"`
	Creating          bool   `json:"creating"`
	StorageBackUp     bool   `json:"storageBackUp"`
	ExtensionType     string `json:"extensionType"`
	CanBeImageStorage bool   `json:"canBeImageStorage"`
	BlockDeviceUUID   string `json:"blockDeviceUuid"`
	OpHostIP          string `json:"opHostIp"`
	IsMount           string `json:"isMount"`
	HostDto           string `json:"hostDto"`
	ScvmOn            bool   `json:"scvmOn"`
}
type StoragePageResponse struct {
	PageResponse
	Items []Storage `json:"items"`
}

type StoragePageReq struct {
	PageReq
}
