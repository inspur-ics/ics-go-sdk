package types

type Disk struct {
	ID              string  `json:"id,omitempty"`
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
	ID                 string      `json:"id"`
	AutoGenerated      bool        `json:"autoGenerated"`
	Name               string      `json:"name"`
	NolocalName        string      `json:"nolocalName"`
	InnerName          string      `json:"innerName,omitempty"`
	DevName            string      `json:"devName"`
	IP                 string      `json:"ip,omitempty"`
	IPv6               string      `json:"ipv6,omitempty"`
	Netmask            string      `json:"netmask,omitempty"`
	Gateway            string      `json:"gateway,omitempty"`
	Mac                string      `json:"mac"`
	Model              string      `json:"model"`
	DeviceID           string      `json:"deviceId"`
	DeviceName         string      `json:"deviceName"`
	DeviceType         string      `json:"deviceType"`
	SwitchType         string      `json:"switchType"`
	VswitchID          string      `json:"vswitchId"`
	UplinkRate         int         `json:"uplinkRate"`
	UplinkBurst        int         `json:"uplinkBurst"`
	DownlinkRate       int         `json:"downlinkRate"`
	DownlinkBurst      int         `json:"downlinkBurst"`
	DownlinkQueue      string      `json:"downlinkQueue,omitempty"`
	Enable             bool        `json:"enable"`
	Status             string      `json:"status"`
	InboundRate        float64     `json:"inboundRate"`
	OutboundRate       float64     `json:"outboundRate"`
	ConnectStatus      bool        `json:"connectStatus"`
	VMName             string      `json:"vmName,omitempty"`
	VMID               string      `json:"vmId,omitempty"`
	VMStatus           string      `json:"vmStatus,omitempty"`
	VMTemplate         bool        `json:"vmTemplate"`
	NetworkName        string      `json:"networkName"`
	NetworkVlan        string      `json:"networkVlan,omitempty"`
	VlanRange          interface{} `json:"vlanRange"`
	NetworkID          string      `json:"networkId"`
	NetworkType        interface{} `json:"networkType"`
	AnotherNetworkType string      `json:"anotherNetworkType"`
	HostIP             string      `json:"hostIp,omitempty"`
	HostStatus         string      `json:"hostStatus,omitempty"`
	HostID             string      `json:"hostId,omitempty"`
	DirectObjName      string      `json:"directObjName,omitempty"`
	TotalOctets        float64     `json:"totalOctets"`
	TotalDropped       float64     `json:"totalDropped"`
	TotalPackets       float64     `json:"totalPackets"`
	TotalBytes         float64     `json:"totalBytes"`
	TotalErrors        float64     `json:"totalErrors"`
	WriteOctets        float64     `json:"writeOctets"`
	WriteDropped       float64     `json:"writeDropped"`
	WritePackets       float64     `json:"writePackets"`
	WriteBytes         float64     `json:"writeBytes"`
	WriteErrors        float64     `json:"writeErrors"`
	ReadOctets         float64     `json:"readOctets"`
	ReadDropped        float64     `json:"readDropped"`
	ReadPackets        float64     `json:"readPackets"`
	ReadBytes          float64     `json:"readBytes"`
	ReadErrors         float64     `json:"readErrors"`
	SecurityGroups     interface{} `json:"securityGroups"`
	AdvancedNetIP      interface{} `json:"advancedNetIp"`
	PortID             interface{} `json:"portId"`
	SdnVFID            interface{} `json:"sdnVFId"`
	OpenstackID        interface{} `json:"openstackId"`
	BindIPEnable       bool        `json:"bindIpEnable"`
	BindIP             interface{} `json:"bindIp"`
	PriorityEnabled    bool        `json:"priorityEnabled"`
	NetPriority        string      `json:"netPriority"`
	VMType             string      `json:"vmType"`
	SystemVMType       interface{} `json:"systemVmType"`
	Dhcp               bool        `json:"dhcp"`
	DhcpIP             interface{} `json:"dhcpIp"`
	UsedDpdk           bool        `json:"usedDpdk"`
	Queues             int         `json:"queues"`
	Speed              string      `json:"speed"`
	FloatIp            string      `json:"floatIp"`
	NatGatewayId       string      `json:"natGatewayId"`
	QueueLengthSet     bool        `json:"queueLengthSet"`
	SendQueueLength    int         `json:"sendQueueLength"`
	ReceiveQueueLength int         `json:"receiveQueueLength"`
	StaticIp           bool        `json:"staticIp"`
	UserIp             string      `json:"userIp"`
	Ipv4PrimaryDNS     string      `json:"ipv4PrimaryDNS"`
	Ipv4SecondDNS      string      `json:"ipv4SecondDNS"`
	Ipv4Netmask        string      `json:"ipv4Netmask"`
	Ipv4Gateway        string      `json:"ipv4Gateway"`
	TenantId           string      `json:"tenantId"`
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
	UserName        string      `json:"userName"`
	UserPwd         string      `json:"userPwd"`
	Domain          string      `json:"domain"`
	DomainOU        string      `json:"domainOU"`
	DomainAdmin     string      `json:"domainAdmin"`
	DomainAdminPass string      `json:"domainAdminPass"`
	InitializeSid   bool        `json:"initializeSid"`
	Scripts         interface{} `json:"scripts"`
}

type GuestOsInfo struct {
	Model               string `json:"model"`
	SocketLimit         int    `json:"socketLimit"`
	SupportCPUHotPlug   bool   `json:"supportCpuHotPlug"`
	SupportCpuHotReduce bool   `json:"supportCpuHotReduce"`
	SupportMemHotPlug   bool   `json:"supportMemHotPlug"`
	SupportMemHotReduce bool   `json:"supportMemHotReduce"`
	SupportDiskHotPlug  bool   `json:"supportDiskHotPlug"`
	SupportUefiBootMode bool   `json:"supportUefiBootMode"`
	SupportVnicHotPlug  bool   `json:"supportVnicHotPlug"`
	SupportStaticIp     bool   `json:"supportStaticIp"`
}

type CloudInit struct {
	MetaData       string `json:"metadata"`
	UserData       string `json:"userdata"`
	DataSourceType string `json:"dataSourceType"`
}

type BootDevice struct {
	ID             string `json:"id"`
	Order          int    `json:"order"`
	BootDeviceType string `json:"bootDeviceType"`
	BootDeviceId   string `json:"bootDeviceId"`
	ConfigId       string `json:"configId"`
}

type GraphicsCard struct {
	GraphicsCardModel  string `json:"graphicsCardModel"`
	GraphicsCardMemory int    `json:"graphicsCardMemory"`
	ScreenNumbers      int    `json:"screenNumbers"`
}

type CdpInfo struct {
	CdpBackupDatastoreId  string `json:"cdpBackupDatastoreId"`
	BackupDataStoreName   string `json:"backupDataStoreName"`
	StartTime             string `json:"startTime"`
	EndTime               string `json:"endTime"`
	EnableCDP             bool   `json:"enableCDP"`
	CdpAvgWriteMBps       int    `json:"cdpAvgWriteMBps"`
	CdpRemainTimes        int    `json:"cdpRemainTimes"`
	CdpLogSpaceSize       int    `json:"cdpLogSpaceSize"`
	CdpLogSpaceSizeInByte int    `json:"cdpLogSpaceSizeInByte"`
	IntervalTime          int    `json:"intervalTime"`
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
	MaxMemoryInByte          int             `json:"maxMemoryInByte"`
	Memory                   int             `json:"memory"`
	MemoryInByte             int             `json:"memoryInByte"`
	MemoryUsage              float64         `json:"memoryUsage"`
	MemHotplugEnabled        bool            `json:"memHotplugEnabled"`
	MemHotplugNumaEnabled    bool            `json:"memHotplugNumaEnabled"`
	EnableHugeMemPage        bool            `json:"enableHugeMemPage"`
	TransPriority            interface{}     `json:"transPriority"`
	CPUNum                   int             `json:"cpuNum"`
	CPUSocket                int             `json:"cpuSocket"`
	CPUCore                  int             `json:"cpuCore"`
	CPUUsage                 float64         `json:"cpuUsage"`
	MaxCPUNum                int             `json:"maxCpuNum"`
	CPUHotplugEnabled        bool            `json:"cpuHotplugEnabled"`
	CPUModelType             string          `json:"cpuModelType"`
	CPUModelEnabled          bool            `json:"cpuModelEnabled"`
	RunningTime              float64         `json:"runningTime"`
	RunningTimeInSeconds     int             `json:"runningTimeInSeconds"`
	ShutDownTime             float64         `json:"shutDownTime"`
	StatusChangedReason      string          `json:"statusChangedReason"`
	Boot                     string          `json:"boot"`
	BootMode                 string          `json:"bootMode"`
	NvramFilePath            string          `json:"nvramFilePath,omitempty"`
	PflashFilePath           string          `json:"pflashFilePath,omitempty"`
	BiosSerialNumber         string          `json:"biosSerialNumber"`
	BootDevices              []BootDevice    `json:"bootDevices"`
	SplashTime               int             `json:"splashTime"`
	StoragePriority          int             `json:"storagePriority"`
	Usb                      interface{}     `json:"usb,omitempty"`
	Usbs                     []interface{}   `json:"usbs,omitempty"`
	Cdrom                    Cdrom           `json:"cdrom,omitempty"`
	ScriptLocation           interface{}     `json:"scriptLocation"`
	CloudInit                CloudInit       `json:"cloudInit,omitempty"`
	Floppy                   interface{}     `json:"floppy,omitempty"`
	Disks                    []Disk          `json:"disks"`
	DelVolumes               interface{}     `json:"delVolumes"`
	Nics                     []Nic           `json:"nics"`
	Gpus                     []interface{}   `json:"gpus,omitempty"`
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
	MemReservationInByte     int             `json:"memReservationInByte"`
	LastBackup               interface{}     `json:"lastBackup"`
	VMType                   string          `json:"vmType"`
	SystemVMType             interface{}     `json:"systemVmType"`
	MemBalloonEnabled        bool            `json:"memBalloonEnabled"`
	Completed                bool            `json:"completed"`
	GraphicsCardModel        string          `json:"graphicsCardModel"`
	GraphicsCardMemory       int             `json:"graphicsCardMemory"`
	GraphicsCards            []GraphicsCard  `json:"graphicsCards"`
	VMHostName               string          `json:"vmHostName"`
	DiskTotalSize            float64         `json:"diskTotalSize"`
	DiskTotalSizeInByte      int             `json:"diskTotalSizeInByte"`
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
	CdpInfo                  CdpInfo         `json:"cdpInfo"`
	GuestOSAuthInfo          GuestOSAuthInfo `json:"guestOSAuthInfo"`
	AwareNumaEnabled         bool            `json:"awareNumaEnabled"`
	DrxEnabled               bool            `json:"drxEnabled"`
	Recyled                  bool            `json:"recyled"`
	Hidden                   bool            `json:"hidden"`
	DeleteTime               string          `json:"deleteTime"`
	DestoryedTime            string          `json:"destoryedTime"`
	SecretLevel              string          `json:"secretLevel"`
	SourceFrom               string          `json:"sourceFrom"`
	SourceFromId             string          `json:"sourceFromId"`
	SourceFromResourceName   string          `json:"sourceFromResourceName"`
	VmDataStoreId            string          `json:"vmDataStoreId"`
	EncryptFlag              bool            `json:"encryptFlag"`
	SecretKeyId              string          `json:"secretKeyId"`
	CpuArchType              string          `json:"cpuArchType"`
	CpuVendor                string          `json:"cpuVendor,omitempty"`
	SupportQuiesce           bool            `json:"supportQuiesce"`
	NetHaEnabled             bool            `json:"netHaEnabled"`
	VmDataStoreName          string          `json:"vmDataStoreName"`
	ProtectionGroupId        string          `json:"protectionGroupId,omitempty"`
	ProtectionGroupName      string          `json:"protectionGroupName,omitempty"`
	LocalWebsiteId           string          `json:"localWebsiteId"`
	SysInitResult            interface{}     `json:"sysInitResult"`
	CreateTime               string          `json:"createTime"`
	DiskKbps                 float64         `json:"diskKbps"`
	DiskIops                 float64         `json:"diskIops"`
	UpdateTime               string          `json:"updateTime"`
	AddToGroupAvailable      bool            `json:"addToGroupAvailable"`
	AddToGroupTip            interface{}     `json:"addToGroupTip"`
	AddToGroupTime           interface{}     `json:"addToGroupTime"`
	RecoveryStartInterval    int             `json:"recoveryStartInterval"`
	RecoveryStartSequence    int             `json:"recoveryStartSequence"`
	SerialPortDevices        interface{}     `json:"serialPortDevices"`
	WatchDogs                []interface{}   `json:"watchDogs"`
	UpsCentral               bool            `json:"upsCentral"`
	EnableIntegrityCheck     bool            `json:"enableIntegrityCheck"`
	IdentificationCode       string          `json:"identificationCode"`
	AutoSyncTimeEnabled      bool            `json:"autoSyncTimeEnabled"`
	EnableAntiVirus          bool            `json:"enableAntiVirus"`
	HostAntivirusStatus      string          `json:"hostAntivirusStatus,omitempty"`
	DisplayNodeVersion       string          `json:"displayNodeVersion"`
	VappId                   string          `json:"vappId"`
	VmfpgaDevs               []interface{}   `json:"vmfpgaDevs"`
}

type VMPageResponse struct {
	PageResponse
	Items []VirtualMachine `json:"items"`
}

type VMPageReq struct {
	PageReq
}

type VMPowerState string
