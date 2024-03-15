package types

type Cluster struct {
	Name              string `json:"name"`    //"cluster"
	Id                string `json:"id"`      //"8a878bda6f7012c7016f87e979120798"
	HostNum           int    `json:"hostNum"` //1
	FreeCpu           string `json:"freeCpu"`
	UsedCpu           string `json:"usedCpu"`
	TotalCpu          string `json:"totalCpu"`
	CpuCoreNum        int    `json:"cpuCoreNum"`
	CpuSocketNum      int    `json:"cpuSocketNum"`
	CpuNum            int    `json:"cpuNum"`
	FreeMemory        string `json:"freeMemory"`
	UsedMemory        string `json:"usedMemory"`
	TotalMemory       string `json:"totalMemory"`
	FreeMemoryInByte  int    `json:"freeMemoryInByte,omitempty"`
	UsedMemoryInByte  int    `json:"usedMemoryInByte,omitempty"`
	TotalMemoryInByte int    `json:"totalMemoryInByte,omitempty"`
	FreeStorage       string `json:"freeStorage,omitempty"`
	UsedStorage       string `json:"usedStorage,omitempty"`
	TotalStorage      string `json:"totalStorage,omitempty"`
	VcpuUsed          int    `json:"vcpuUsed"`
	Drs               struct {
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
		FailoverHostIds       []string      `json:"failoverHostIds"`
		HAPriority            string        `json:"haPriority,omitempty"`
		NetHaEnabled          bool          `json:"netHaEnabled"`
		NetTolerance          string        `json:"netTolerance"`
		NetProcessStrategy    string        `json:"netProcessStrategy"`
		ContainerHaEnabled    bool          `json:"containerHaEnabled"`
		CpuReserve            int           `json:"cpuReserve"`
		MemoryReserve         int           `json:"memoryReserve"`
		UsedCpu               float64       `json:"usedCpu"`
		TotalCpu              float64       `json:"totalCpu"`
		UsedMemory            float64       `json:"usedMemory"`
		TotalMemory           int           `json:"totalMemory"`
		UsedMemoryInByte      int           `json:"usedMemoryInByte,omitempty"`
		TotalMemoryInByte     int           `json:"totalMemoryInByte,omitempty"`
		CdpProcessStrategy    string        `json:"cdpProcessStrategy,omitempty"`
	} `json:"ha"`
	DataCenterDto              Datacenter  `json:"dataCenterDto"`
	HostIds                    []string    `json:"hostIds"`
	Tags                       []Tag       `json:"tags"`
	CpuFreePercent             float64     `json:"cpuFreePercent"`
	MemFreePercent             float64     `json:"memFreePercent"`
	VmTotalCpu                 int         `json:"vmTotalCpu,omitempty"`
	VmTotalMemInMB             int         `json:"vmTotalMemInMB,omitempty"`
	HostTotalMemInMB           float64     `json:"hostTotalMemInMB,omitempty"`
	VmTotalMemInByte           int         `json:"vmTotalMemInByte,omitempty"`
	HostTotalMemInByte         int         `json:"hostTotalMemInByte,omitempty"`
	CpuTotalRatio              float64     `json:"cpuTotalRatio,omitempty"`
	MemTotalRatio              float64     `json:"memTotalRatio,omitempty"`
	TotalStorageInByte         int         `json:"totalStorageInByte,omitempty"`
	UsedStorageInByte          int         `json:"usedStorageInByte,omitempty"`
	FreeStorageInByte          int         `json:"freeStorageInByte,omitempty"`
	StorageUsage               float64     `json:"storageUsage,omitempty"`
	VirtualVolumeStorageInByte int         `json:"virtualVolumeStorageInByte,omitempty"`
	HostTotalStorageInByte     int         `json:"hostTotalStorageInByte,omitempty"`
	StorageRatio               float64     `json:"storageRatio,omitempty"`
	CpuUsage                   float64     `json:"cpuUsage,omitempty"`
	MemoryUsage                float64     `json:"memoryUsage,omitempty"`
	CpuArchType                string      `json:"cpuArchType,omitempty"`
	VmNum                      int         `json:"vmNum"`
	PodNum                     int         `json:"podNum"`
}

type ClusterListRsp struct {
	Items []Cluster `json:"items"`
}
