package types

type Storage struct {
	DataStoreType       string  `json:"dataStoreType"`
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	MountPath           string  `json:"mountPath"`
	Capacity            float64 `json:"capacity"`
	CapacityInByte      int     `json:"capacityInByte,omitempty"`
	UsedCapacity        float64 `json:"usedCapacity"`
	UsedCapacityInByte  int     `json:"usedCapacityInByte,omitempty"`
	AvailCapacity       float64 `json:"availCapacity"`
	AvailCapacityInByte int     `json:"availCapacityInByte,omitempty"`
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
	XactiveStoreName  string      `json:"xactiveStoreName,omitempty"`
	XactiveStoreId    string      `json:"xactiveStoreId,omitempty"`
	DataCenterDto     interface{} `json:"dataCenterDto,omitempty"`
	HostNumbers       int         `json:"hostNumbers"`
	PodHostNumbers    int         `json:"podHostNumbers"`
	VMNumbers         int         `json:"vmNumbers"`
	VolumesNumbers    int         `json:"volumesNumbers"`
	VMTemplateNumbers int         `json:"vmTemplateNumbers"`
	PodNumbers        int         `json:"podNumbers"`
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
	DatastoreRole     interface{} `json:"datastoreRole"`
	CanCreateXactive  bool        `json:"canCreateXactive"`
	Accelerator       string      `json:"accelerator"`
	IscsiServerId     string      `json:"iscsiServerId"`
	CanUmount         bool        `json:"canUmount"`
	Iops              float64     `json:"iops"`
	Kbps              float64     `json:"kbps"`
	MaxReadRate       int         `json:"maxReadRate"`
	MaxWriteRate      int         `json:"maxWriteRate"`
	DepthReadRate     int         `json:"depthReadRate"`
	DepthWriteRate    int         `json:"depthWriteRate"`
	ReadBandwidth     interface{} `json:"readBandwidth"`
	WriteBandwidth    interface{} `json:"writeBandwidth"`
	MaxReadDelay      float64     `json:"maxReadDelay"`
	MaxWriteDelay     float64     `json:"maxWriteDelay"`
	DepthReadDelay    float64     `json:"depthReadDelay"`
	DepthWriteDelay   float64     `json:"depthWriteDelay"`
	BlockDeviceUuid   string      `json:"blockDeviceUuid"`
	OpHostIp          string      `json:"opHostIp"`
	IsMount           bool        `json:"isMount"`
	DetectIORate      interface{} `json:"detectIORate"`
	HostDto           interface{} `json:"hostDto"`
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
