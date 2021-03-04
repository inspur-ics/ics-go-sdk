package types

type VMTemplate struct {
	ID                       string          `json:"id"`
	CustomVmId               string          `json:"customVmId"`
	Name                     string          `json:"name"`
	State                    string          `json:"state"`
	Status                   string          `json:"status"`
	HostID                   string          `json:"hostId"`
	HostName                 string          `json:"hostName"`
	HostIp                   string          `json:"hostIp"`
	HostStatus               string          `json:"hostStatus"`
	HostMemory               float64         `json:"hostMemory"`
	DataCenterId             string          `json:"dataCenterId"`
	HaEnabled                bool            `json:"haEnabled"`
	RouterFlag               bool            `json:"routerFlag"`
	Migratable               bool            `json:"migratable"`
	HostBinded               bool            `json:"hostBinded"`
	ToolsInstalled           bool            `json:"toolsInstalled"`
	ToolsVersion             interface{}     `json:"toolsVersion"`
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
	DataStoreId              string          `json:"dataStoreId"`
	SdsdomainId              string          `json:"sdsdomainId"`
	ClockModel               string          `json:"clockModel"`
	CPULimit                 int             `json:"cpuLimit"`
	MemShares                int             `json:"memShares"`
	CPUReservation           int             `json:"cpuReservation"`
	MemReservation           int             `json:"memReservation"`
	LastBackup               interface{}     `json:"lastBackup"`
	VMType                   string          `json:"vmType"`
	SystemVmType             interface{}     `json:"systemVmType"`
	MemBalloonEnabled        bool            `json:"memBalloonEnabled"`
	Completed                bool            `json:"completed"`
	GraphicsCardModel        string          `json:"graphicsCardModel"`
	GraphicsCardMemory       int             `json:"graphicsCardMemory"`
	GraphicsCards            interface{}     `json:"graphicsCards"`
	VMHostName               string          `json:"vmHostName"`
	DiskTotalSize            float64         `json:"diskTotalSize"`
	diskUsedSize             float64         `json:"diskUsedSize"`
	diskUsage                float64         `json:"diskUsage"`
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
}

type VMTemplatePageResponse struct {
	PageResponse
	Items []VMTemplate `json:"items"`
}

type VMTemplatePageReq struct {
	PageReq
}
