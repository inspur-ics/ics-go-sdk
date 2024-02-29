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
	HostID              string  `json:"hostId,omitempty"`
	MountStatus         string  `json:"mountStatus,omitempty"`
	HostIP              string  `json:"hostIp,omitempty"`
	UUID                string  `json:"uuid"`
	AbsolutePath        string  `json:"absolutePath,omitempty"`
	DataCenterName      string  `json:"dataCenterName,omitempty"`
	DataCenterOrHostDto struct {
		DataCenterOrHost string `json:"dataCenterOrHost"`
		DataCenterName   string `json:"dataCenterName,omitempty"`
		HostName         string `json:"hostName,omitempty"`
		Status           string `json:"status,omitempty"`
	} `json:"dataCenterOrHostDto"`
	BlockDeviceDto    interface{} `json:"blockDeviceDto,omitempty"`
	DataCenterDto     interface{} `json:"dataCenterDto,omitempty"`
	HostNumbers       int         `json:"hostNumbers"`
	VMNumbers         int         `json:"vmNumbers"`
	VolumesNumbers    int         `json:"volumesNumbers"`
	VMTemplateNumbers int         `json:"vmTemplateNumbers"`
	Tags              interface{} `json:"tags,omitempty"`
	MaxSlots          int         `json:"maxSlots"`
	Creating          bool        `json:"creating"`
	StorageBackUp     bool        `json:"storageBackUp"`
	ExtensionType     string      `json:"extensionType"`
	CanBeImageStorage bool        `json:"canBeImageStorage"`
	MultiplexRatio    float64     `json:"multiplexRatio"`
	Oplimit           bool        `json:"oplimit"`
	Maxop             int         `json:"maxop"`
	MountStateCount   interface{} `json:"mountStateCount"`
	ScvmOn            bool        `json:"scvmOn"`
	AllocPolicy       string      `json:"allocPolicy,omitempty"`
}

type StoragePageResponse struct {
	PageResponse
	Items []Storage `json:"items"`
}

type StoragePageReq struct {
	PageReq
}

type ImageFileInfo struct {
	Name           string `json:"name"`
	SourceType     string `json:"sourceType"`
	Format         string `json:"format"`
	FileType       string `json:"fileType"`
	Date           string `json:"date"`
	Path           string `json:"path"`
	FtpServer      string `json:"ftpServer"`
	DataStoreID    string `json:"dataStoreId"`
	DataStoreName  string `json:"dataStoreName"`
	ServerID       string `json:"serverId"`
	Md5            string `json:"md5"`
	FileSizeInByte int    `json:"fileSizeInByte"`
	RealSizeInByte int    `json:"realSizeInByte"`
}

type ImageFilePageResponse struct {
	PageResponse
	Items []ImageFileInfo `json:"items"`
}
