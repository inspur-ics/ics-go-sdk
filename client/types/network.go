package types

type Datacenter struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	NfsPath            string        `json:"nfsPath"`
	Type               string        `json:"type"`
	Description        string        `json:"description"`
	HostNum            int           `json:"hostNum"`
	VMNum              int           `json:"vmNum"`
	ClusterNum         int           `json:"clusterNum"`
	StorageNum         int           `json:"storageNum"`
	NetNum             int           `json:"netNum"`
	NeutronNetNum      int           `json:"neutronNetNum"`
	CPUCapacity        string        `json:"cpuCapacity"`
	CPUAvailable       string        `json:"cpuAvailable"`
	CPUUsed            string        `json:"cpuUsed"`
	CPUUtilization     string        `json:"cpuUtilization"`
	MemoryCapacity     string        `json:"memoryCapacity"`
	MemoryAvailable    string        `json:"memoryAvailable"`
	MemoryUsed         string        `json:"memoryUsed"`
	MemoryUtilization  string        `json:"memoryUtilization"`
	StorageCapacity    string        `json:"storageCapacity"`
	StorageAvailable   string        `json:"storageAvailable"`
	StorageUsed        string        `json:"storageUsed"`
	StorageUtilization string        `json:"storageUtilization"`
	DatastoreNum       int           `json:"datastoreNum"`
	NetworkType        interface{}   `json:"networkType"`
	VswitchDtos        interface{}   `json:"vswitchDtos"`
	SdnNetworkDtos     []interface{} `json:"sdnNetworkDtos"`
	SdnInit            bool          `json:"sdnInit"`
}

type DatacenterPageResponse struct {
	PageResponse
	Items               []Datacenter         `json:"items"`
}