package types

type Host struct {
	ID                  string        `json:"id"`
	IP                  string   `json:"ip"`
	SwitchUplinkPortDto interface{}   `json:"switchUplinkPortDto"`
	UplinkTopoDto       interface{}   `json:"uplinkTopoDto"`
	Pnics               interface{}   `json:"pnics"`
	Disks               interface{}   `json:"disks"`
	Name                string        `json:"name"`
	HostName            string        `json:"hostName"`
	NodeVersion         string        `json:"nodeVersion"`
	Password            string        `json:"password"`
	DataCenterID        string        `json:"dataCenterId"`
	DataCenterName      string        `json:"dataCenterName"`
	ClusterName         string        `json:"clusterName"`
	ClusterID           string        `json:"clusterId"`
	Status              string        `json:"status"`
	CPUSocket           int           `json:"cpuSocket"`
	CPUCorePerSocket    int           `json:"cpuCorePerSocket"`
	CPUThreadPerCore    int           `json:"cpuThreadPerCore"`
	LogicCPUNum         int           `json:"logicCpuNum"`
	LogicalProcessor    int           `json:"logicalProcessor"`
	CPUFrequency        float64       `json:"cpuFrequency"`
	CPUUsage            float64       `json:"cpuUsage"`
	CPUTotalHz          float64       `json:"cpuTotalHz"`
	FreeCPU             float64       `json:"freeCpu"`
	UsedCPU             float64       `json:"usedCpu"`
	TotalMem            float64       `json:"totalMem"`
	LogicTotalMem       float64       `json:"logicTotalMem"`
	MemoryUsage         float64       `json:"memoryUsage"`
	FreeMemory          float64       `json:"freeMemory"`
	UsedMemory          float64       `json:"usedMemory"`
	LogicUsedMemory     float64       `json:"logicUsedMemory"`
	LogicFreeMemory     float64       `json:"logicFreeMemory"`
	PnicNum             int           `json:"pnicNum"`
	NormalRunTime       float64       `json:"normalRunTime"`
	Model               string        `json:"model"`
	CPUType             string        `json:"cpuType"`
	VtDegree            float64       `json:"vtDegree"`
	Powerstate          interface{}   `json:"powerstate"`
	HostBmcDto          interface{}   `json:"hostBmcDto"`
	Tags                []interface{} `json:"tags"`
	MountPath           interface{}   `json:"mountPath"`
	MonMountState       interface{}   `json:"monMountState"`
	CPUModel            interface{}   `json:"cpuModel"`
	NetworkDtos         interface{}   `json:"networkDtos"`
	PortIP              interface{}   `json:"portIp"`
	Monstatus           bool          `json:"monstatus"`
	HostIqn             interface{}   `json:"hostIqn"`
	VxlanPortDto        interface{}   `json:"vxlanPortDto"`
	SdnUpLinks          interface{}   `json:"sdnUpLinks"`
	AllPNicsCount       int           `json:"allPNicsCount"`
	AvailablePNicsCount int           `json:"availablePNicsCount"`
	CfsDomainStatus     interface{}   `json:"cfsDomainStatus"`
	SerialNumber        interface{}   `json:"serialNumber"`
	Manufacturer        interface{}   `json:"manufacturer"`
	IndicatorStatus     interface{}   `json:"indicatorStatus"`
	EntryTemperature    interface{}   `json:"entryTemperature"`
	MulticastEnabled    bool          `json:"multicastEnabled"`
	Pcies               interface{}   `json:"pcies"`
	VgpuEnable          bool          `json:"vgpuEnable"`
	SpecialFailover     bool          `json:"specialFailover"`
	VswitchDtos         interface{}   `json:"vswitchDtos"`
	HotfixVersion       string        `json:"hotfixVersion"`
}

type HostPageResponse struct {
	TotalPage   int     `json:"totalPage"`
	CurrentPage int     `json:"currentPage"`
	TotalSize   int     `json:"totalSize"`
	Items       []Host  `json:"items"`
}
