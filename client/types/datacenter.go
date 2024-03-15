package types

type Datacenter struct {
	ID                         string        `json:"id"`
	Name                       string        `json:"name"`
	NfsPath                    string        `json:"nfsPath"`
	NfsVersion                 string        `json:"nfsVersion,omitempty"`
	Type                       string        `json:"type"`
	Description                string        `json:"description"`
	HostNum                    int           `json:"hostNum"`
	VMNum                      int           `json:"vmNum"`
	PodNum                     int           `json:"podNum"`
	ClusterNum                 int           `json:"clusterNum"`
	CfsDomainNum               int           `json:"cfsDomainNum"`
	SdsDomainNum               int           `json:"sdsDomainNum"`
	StorageNum                 int           `json:"storageNum"`
	ImageIsoNum                int           `json:"imageIsoNum"`
	NetNum                     int           `json:"netNum"`
	NeutronNetNum              int           `json:"neutronNetNum"`
	CPUCapacity                string        `json:"cpuCapacity"`
	CPUAvailable               string        `json:"cpuAvailable"`
	CPUUsed                    string        `json:"cpuUsed"`
	CPUUtilization             string        `json:"cpuUtilization"`
	CPUSocketNum               int           `json:"cpuSocketNum"`
	CPUCoreNum                 int           `json:"cpuCoreNum"`
	CPUNum                     int           `json:"cpuNum"`
	MemoryCapacity             string        `json:"memoryCapacity"`
	MemoryAvailable            string        `json:"memoryAvailable"`
	MemoryUsedInByte           int           `json:"memoryUsedInByte,omitempty"`
	MemoryCapacityInByte       int           `json:"memoryCapacityInByte,omitempty"`
	MemoryAvailableInByte      int           `json:"memoryAvailableInByte,omitempty"`
	MemoryUsed                 string        `json:"memoryUsed"`
	MemoryUtilization          string        `json:"memoryUtilization"`
	StorageCapacity            string        `json:"storageCapacity"`
	StorageAvailable           string        `json:"storageAvailable"`
	StorageUsed                string        `json:"storageUsed"`
	StorageUtilization         string        `json:"storageUtilization"`
	StorageCapacityInByte      int           `json:"storageCapacityInByte,omitempty"`
	StorageAvailableInByte     int           `json:"storageAvailableInByte,omitempty"`
	StorageUsedInByte          int           `json:"storageUsedInByte,omitempty"`
	DatastoreNum               int           `json:"datastoreNum"`
	LocalStoreNum              int           `json:"localstoreNum"`
	CfsStoreNum                int           `json:"cfsstoreNum"`
	RawStoreNum                int           `json:"rawstoreNum"`
	NfsStoreNum                int           `json:"nfsstoreNum"`
	XactiveStoreNum            int           `json:"xactivestoreNum"`
	NetworkType                string        `json:"networkType"`
	VswitchDtos                interface{}   `json:"vswitchDtos"`
	SdnNetworkDtos             []interface{} `json:"sdnNetworkDtos"`
	SdnInit                    bool          `json:"sdnInit"`
	SdnSpeedUp                 bool          `json:"sdnSpeedUp"`
	SdnConfigDto               interface{}   `json:"sdnConfigDto"`
	CpuArchType                string        `json:"cpuArchType"`
	VmTotalCpu                 int           `json:"vmTotalCpu,omitempty"`
	VmTotalMemInMB             int           `json:"vmTotalMemInMB,omitempty"`
	HostTotalMemInMB           float64       `json:"hostTotalMemInMB,omitempty"`
	VmTotalMemInByte           int           `json:"vmTotalMemInByte,omitempty"`
	HostTotalMemInByte         int           `json:"hostTotalMemInByte,omitempty"`
	CpuTotalRatio              float64       `json:"cpuTotalRatio,omitempty"`
	MemTotalRatio              float64       `json:"memTotalRatio,omitempty"`
	VirtualVolumeStorageInGB   float64       `json:"virtualVolumeStorageInGB,omitempty"`
	HostTotalStorageInGB       float64       `json:"hostTotalStorageInGB,omitempty"`
	VirtualVolumeStorageInByte int           `json:"virtualVolumeStorageInByte,omitempty"`
	HostTotalStorageInByte     int           `json:"hostTotalStorageInByte,omitempty"`
	StorageRatio               float64       `json:"storageRatio,omitempty"`
}

type DatacenterPageResponse struct {
	PageResponse
	Items []Datacenter `json:"items"`
}
