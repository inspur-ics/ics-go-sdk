package types

type Cluster struct {
	Name         string `json:"name"`    //"cluster"
	Id           string `json:"id"`      //"8a878bda6f7012c7016f87e979120798"
	HostNum      int    `json:"hostNum"` //1
	FreeCpu      string `json:"freeCpu"`
	UsedCpu      string `json:"usedCpu"`
	TotalCpu     string `json:"totalCpu"`
	FreeMemory   string `json:"freeMemory"`
	UsedMemory   string `json:"usedMemory"`
	TotalMemory  string `json:"totalMemory"`
	FreeStorage  string `json:"freeStorage,omitempty"`
	UsedStorage  string `json:"usedStorage,omitempty"`
	TotalStorage string `json:"totalStorage,omitempty"`
	CpuNum       int    `json:"cpuNum"`
	VcpuUsed     int    `json:"vcpuUsed"`
	Drs          struct {
		ID                string `json:"id,omitempty"`
		ClusterID         string `json:"clusterID,omitempty"`
		DrsEnabled        bool   `json:"drsEnabled"`
		CpuThreshold      int    `json:"cpuThreshold"`
		MemoryThreshold   int    `json:"memoryThreshold"`
		VmMigrationCount  int    `json:"vmMigrationCount"`
		RelMigrateEnabled bool   `json:"relMigrateEnabled"`
		DpmEnabled        bool   `json:"dpmEnabled"`
		CpuLowThreshold   int    `json:"cpuLowThreshold"`
		MemLowThreshold   int    `json:"memLowThreshold"`
		MinReserveHost    int    `json:"minReserveHost"`
	} `json:"drs"`
	HA struct {
		ID                    string        `json:"id,omitempty"`
		ClusterID             string        `json:"clusterID,omitempty"`
		HAEnabled             bool          `json:"haEnabled"`
		HAMaxLimit            int           `json:"haMaxLimit"`
		AccessControlEnabled  bool          `json:"accessControlEnabled"`
		AccessControlStrategy string        `json:"accessControlStrategy,omitempty"`
		FailoverHosts         []interface{} `json:"failoverHosts"`
		FailoverHostIds       interface{}   `json:"failoverHostIds"`
		HAPriority            string        `json:"haPriority,omitempty"`
	} `json:"ha"`
	DataCenterDto Datacenter  `json:"dataCenterDto"`
	HostIds       interface{} `json:"hostIds"`
	Tags          []Tag       `json:"tags"`
}

type ClusterListRsp struct {
	Items []Cluster `json:"items"`
}
