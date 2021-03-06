package types

type Disk struct {
	ID              string  `json:"id"`
	Label           string  `json:"label"`
	ScsiID          string  `json:"scsiId"`
	Enabled         bool    `json:"enabled"`
	WriteBps        int     `json:"writeBps"`
	ReadBps         int     `json:"readBps"`
	TotalBps        int     `json:"totalBps"`
	TotalIops       int     `json:"totalIops"`
	WriteIops       int     `json:"writeIops"`
	ReadIops        int     `json:"readIops"`
	Volume          Volume  `json:"volume"`
	BusModel        string  `json:"busModel"`
	Usage           float64 `json:"usage"`
	MonReadIops     float64 `json:"monReadIops"`
	MonWriteIops    float64 `json:"monWriteIops"`
	ReadThroughput  float64 `json:"readThroughput"`
	WriteThroughput float64 `json:"writeThroughput"`
	ReadWriteModel  string  `json:"readWriteModel"`
	EnableNativeIO  bool    `json:"enableNativeIO"`
	EnableKernelIO  bool    `json:"enableKernelIO"`
	L2CacheSize     int     `json:"l2CacheSize"`
	QueueNum        int     `json:"queueNum"`
}
type Nic struct {
	ID              string      `json:"id"`
	AutoGenerated   bool        `json:"autoGenerated"`
	Name            string      `json:"name"`
	NolocalName     string      `json:"nolocalName"`
	InnerName       string      `json:"innerName,omitempty"`
	DevName         string      `json:"devName"`
	IP              string      `json:"ip,omitempty"`
	Netmask         string      `json:"netmask,omitempty"`
	Gateway         string      `json:"gateway,omitempty"`
	Mac             string      `json:"mac"`
	Model           string      `json:"model"`
	DeviceID        string      `json:"deviceId"`
	DeviceName      string      `json:"deviceName"`
	DeviceType      string      `json:"deviceType"`
	SwitchType      string      `json:"switchType"`
	VswitchID       string      `json:"vswitchId"`
	UplinkRate      int         `json:"uplinkRate"`
	UplinkBurst     int         `json:"uplinkBurst"`
	DownlinkRate    int         `json:"downlinkRate"`
	DownlinkBurst   int         `json:"downlinkBurst"`
	DownlinkQueue   string      `json:"downlinkQueue,omitempty"`
	Enable          bool        `json:"enable"`
	Status          string      `json:"status"`
	InboundRate     float64     `json:"inboundRate"`
	OutboundRate    float64     `json:"outboundRate"`
	ConnectStatus   bool        `json:"connectStatus"`
	VMName          string      `json:"vmName,omitempty"`
	VMID            string      `json:"vmId,omitempty"`
	VMStatus        string      `json:"vmStatus,omitempty"`
	VMTemplate      bool        `json:"vmTemplate"`
	NetworkName     string      `json:"networkName"`
	NetworkVlan     string      `json:"networkVlan,omitempty"`
	VlanRange       interface{} `json:"vlanRange"`
	NetworkID       string      `json:"networkId"`
	NetworkType     interface{} `json:"networkType"`
	HostIP          string      `json:"hostIp,omitempty"`
	HostStatus      string      `json:"hostStatus,omitempty"`
	HostID          string      `json:"hostId,omitempty"`
	DirectObjName   string      `json:"directObjName,omitempty"`
	TotalOctets     float64     `json:"totalOctets"`
	TotalDropped    float64     `json:"totalDropped"`
	TotalPackets    float64     `json:"totalPackets"`
	TotalBytes      float64     `json:"totalBytes"`
	TotalErrors     float64     `json:"totalErrors"`
	WriteOctets     float64     `json:"writeOctets"`
	WriteDropped    float64     `json:"writeDropped"`
	WritePackets    float64     `json:"writePackets"`
	WriteBytes      float64     `json:"writeBytes"`
	WriteErrors     float64     `json:"writeErrors"`
	ReadOctets      float64     `json:"readOctets"`
	ReadDropped     float64     `json:"readDropped"`
	ReadPackets     float64     `json:"readPackets"`
	ReadBytes       float64     `json:"readBytes"`
	ReadErrors      float64     `json:"readErrors"`
	SecurityGroups  interface{} `json:"securityGroups"`
	AdvancedNetIP   interface{} `json:"advancedNetIp"`
	PortID          interface{} `json:"portId"`
	SdnVFID         interface{} `json:"sdnVFId"`
	OpenstackID     interface{} `json:"openstackId"`
	BindIPEnable    bool        `json:"bindIpEnable"`
	BindIP          interface{} `json:"bindIp"`
	PriorityEnabled bool        `json:"priorityEnabled"`
	NetPriority     interface{} `json:"netPriority"`
	VMType          string      `json:"vmType"`
	SystemVMType    interface{} `json:"systemVmType"`
	Dhcp            bool        `json:"dhcp"`
	DhcpIP          interface{} `json:"dhcpIp"`
	UsedDpdk        bool        `json:"usedDpdk"`
	Queues          int         `json:"queues"`
}

type Floppy struct {
	Path      string      `json:"path"`
	DataStore interface{} `json:"dataStore"`
	VfdType   string      `json:"vfdType"`
}

type Cdrom struct {
	Path           string      `json:"path"`
	Type           string      `json:"type"`
	Connected      bool        `json:"connected"`
	StartConnected bool        `json:"startConnected"`
	CifsDto        interface{} `json:"cifsDto"`
	DataStore      interface{} `json:"dataStore"`
}

type GuestOSAuthInfo struct {
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}

type GuestOsInfo struct {
	Model               string `json:"model"`
	SocketLimit         int    `json:"socketLimit"`
	SupportCPUHotPlug   bool   `json:"supportCpuHotPlug"`
	SupportMemHotPlug   bool   `json:"supportMemHotPlug"`
	SupportDiskHotPlug  bool   `json:"supportDiskHotPlug"`
	SupportUefiBootMode bool   `json:"supportUefiBootMode"`
}

type VirtualMachine struct {
	ID                       string          `json:"id"`
	CustomVmId               string          `json:"customVmId"`
	Name                     string          `json:"name"`
	PowerState               VMPowerState    `json:"state"`
	Status                   string          `json:"status"`
	HostID                   string          `json:"hostId"`
	HostName                 string          `json:"hostName"`
	HostIP                   string          `json:"hostIp"`
	HostStatus               string          `json:"hostStatus"`
	HostMemory               float64         `json:"hostMemory"`
	DataCenterID             string          `json:"dataCenterId"`
	HaEnabled                bool            `json:"haEnabled"`
	RouterFlag               bool            `json:"routerFlag"`
	Migratable               bool            `json:"migratable"`
	HostBinded               bool            `json:"hostBinded"`
	ToolsInstalled           bool            `json:"toolsInstalled"`
	ToolsVersion             string          `json:"toolsVersion"`
	ToolsType                string          `json:"toolsType"`
	ToolsVersionStatus       string          `json:"toolsVersionStatus"`
	ToolsRunningStatus       string          `json:"toolsRunningStatus"`
	ToolsNeedUpdate          bool            `json:"toolsNeedUpdate"`
	Description              string          `json:"description"`
	HaMaxLimit               int             `json:"haMaxLimit"`
	Template                 bool            `json:"template"`
	Initialized              bool            `json:"initialized"`
	GuestosLabel             string          `json:"guestosLabel"`
	GuestosType              string          `json:"guestosType"`
	GuestOsInfo              GuestOsInfo     `json:"guestOsInfo"`
	InnerName                string          `json:"innerName"`
	UUID                     string          `json:"uuid"`
	MaxMemory                int             `json:"maxMemory"`
	Memory                   int             `json:"memory"`
	MemoryUsage              float64         `json:"memoryUsage"`
	MemHotplugEnabled        bool            `json:"memHotplugEnabled"`
	EnableHugeMemPage        bool            `json:"enableHugeMemPage"`
	CPUNum                   int             `json:"cpuNum"`
	CPUSocket                int             `json:"cpuSocket"`
	CPUCore                  int             `json:"cpuCore"`
	CPUUsage                 float64         `json:"cpuUsage"`
	MaxCPUNum                int             `json:"maxCpuNum"`
	CPUHotplugEnabled        bool            `json:"cpuHotplugEnabled"`
	CPUModelType             string          `json:"cpuModelType"`
	CPUModelEnabled          bool            `json:"cpuModelEnabled"`
	RunningTime              float64         `json:"runningTime"`
	Boot                     string          `json:"boot"`
	BootMode                 string          `json:"bootMode"`
	SplashTime               int             `json:"splashTime"`
	StoragePriority          int             `json:"storagePriority"`
	Usb                      interface{}     `json:"usb"`
	Usbs                     []interface{}   `json:"usbs"`
	Cdrom                    Cdrom           `json:"cdrom"`
	Floppy                   Floppy          `json:"floppy"`
	Disks                    []Disk          `json:"disks"`
	Nics                     []Nic           `json:"nics"`
	Gpus                     []interface{}   `json:"gpus"`
	VMPcis                   []interface{}   `json:"vmPcis"`
	ConfigLocation           string          `json:"configLocation"`
	HotplugEnabled           bool            `json:"hotplugEnabled"`
	VncPort                  int             `json:"vncPort"`
	VncPasswd                string          `json:"vncPasswd"`
	VncSharePolicy           string          `json:"vncSharePolicy"`
	CpuBindType              string          `json:"cpuBindType"`
	VcpuPin                  string          `json:"vcpuPin"`
	VcpuPins                 []string        `json:"vcpuPins"`
	CPUShares                int             `json:"cpuShares"`
	PanickPolicy             string          `json:"panickPolicy"`
	DataStoreID              string          `json:"dataStoreId"`
	SdsdomainID              string          `json:"sdsdomainId"`
	ClockModel               string          `json:"clockModel"`
	CPULimit                 int             `json:"cpuLimit"`
	MemShares                int             `json:"memShares"`
	CPUReservation           int             `json:"cpuReservation"`
	MemReservation           int             `json:"memReservation"`
	LastBackup               interface{}     `json:"lastBackup"`
	VMType                   string          `json:"vmType"`
	SystemVMType             interface{}     `json:"systemVmType"`
	MemBalloonEnabled        bool            `json:"memBalloonEnabled"`
	Completed                bool            `json:"completed"`
	GraphicsCardModel        string          `json:"graphicsCardModel"`
	GraphicsCardMemory       int             `json:"graphicsCardMemory"`
	GraphicsCards            interface{}     `json:"graphicsCards"`
	VMHostName               string          `json:"vmHostName"`
	DiskTotalSize            float64         `json:"diskTotalSize"`
	DiskUsedSize             float64         `json:"diskUsedSize"`
	DiskUsage                float64         `json:"diskUsage"`
	Tags                     interface{}     `json:"tags"`
	StartPriority            string          `json:"startPriority"`
	OwnerName                string          `json:"ownerName"`
	Version                  string          `json:"version"`
	EnableReplicate          bool            `json:"enableReplicate"`
	ReplicationDatastoreId   string          `json:"replicationDatastoreId"`
	ReplicationDatastoreName string          `json:"replicationDatastoreName"`
	RecoveryFlag             bool            `json:"recoveryFlag"`
	SpiceUsbNum              int             `json:"spiceUsbNum"`
	CdpInfo                  interface{}     `json:"cdpInfo"`
	GuestOSAuthInfo          GuestOSAuthInfo `json:"guestOSAuthInfo"`
	AwareNumaEnabled         bool            `json:"awareNumaEnabled"`
	ExtendData               string          `json:"extendData,omitempty"`
	CloudInited              bool            `json:"cloudInited,omitempty"`
}

type VMPageResponse struct {
	PageResponse
	Items []VirtualMachine `json:"items"`
}

type VMPageReq struct {
	PageReq
}

type VMPowerState string
