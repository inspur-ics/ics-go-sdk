package types

type NetworkPageResponse struct {
	TotalPage   int       `json:"totalPage"`
	CurrentPage int       `json:"currentPage"`
	TotalSize   int       `json:"totalSize"`
	Items       []Network `json:"items"`
}

type Switch struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	ResourceID       string      `json:"resourceId"`
	ControllerIP     string      `json:"controllerIP"`
	DataCenterDto    Datacenter  `json:"dataCenterDto"`
	HostDtos         Host        `json:"hostDtos"`
	SwitchType       string      `json:"switchType"`
	AppType          string      `json:"appType"`
	Description      string      `json:"description"`
	NetworkDtos      interface{} `json:"networkDtos"`
	SdnNetworkDtos   interface{} `json:"sdnNetworkDtos"`
	VMDtos           interface{} `json:"vmDtos"`
	HostNum          int         `json:"hostNum"`
	PnicNum          int         `json:"pnicNum"`
	NetworkNum       int         `json:"networkNum"`
	VMNum            int         `json:"vmNum"`
	ProdNum          int         `json:"prodNum"`
	MaxVfs           int         `json:"maxvfs"`
	ThirdPartySDN    bool        `json:"thirdPartySDN"`
	Hierarchy        bool        `json:"hierarchy"`
	ConnectStorage   bool        `json:"connectStorage"`
	ConnectManage    bool        `json:"connectManage"`
	ConnectSwitches  interface{} `json:"connectSwitches"`
	DhcpProtection   bool        `json:"dhcpProtection"`
	NeutronName      interface{} `json:"neutronName"`
	NeutronPassword  interface{} `json:"neutronPassword"`
	ConnectScvm      bool        `json:"connectScvm"`
	SwitchUplinkType string      `json:"switchUplinkType"`
	ComputerNetNum   int         `json:"computerNetNum"`
	DataNetNum       int         `json:"dataNetNum"`
	MigrateNetNum    int         `json:"migrateNetNum"`
	VmMigBandWidth   string      `json:"vmMigBandWidth"`
	EnableDpdk       bool        `json:"enableDpdk"`
	SflowStatus      bool        `json:"sflowStatus"`
	NetflowStatus    bool        `json:"netflowStatus"`
	MulticastStatus  bool        `json:"multicastStatus"`
	MirrorStatus     bool        `json:"mirrorStatus"`
	BrLimitStatus    bool        `json:"brLimitStatus"`
	Hidden           bool        `json:"hidden"`
	NetworkTopoly    bool        `json:"networkTopoly"`
	ArbitrativeIp    string      `json:"arbitrativeIp"`
	HbAutoCreate     bool        `json:"hbAutoCreate"`
	EnableFcoe       bool        `json:"enableFcoe"`
	EnableSc         bool        `json:"enableSc"`
	EnableTrust      bool        `json:"enableTrust"`
	MTU              int         `json:"mtu"`
	ConnectInCloudOs bool        `json:"connectInCloudOs"`
	IamServerDto     interface{} `json:"iamServerDto"`
}

type Network struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	ResourceID      string        `json:"resourceId"`
	Vlan            int           `json:"vlan"`
	VlanFlag        bool          `json:"vlanFlag"`
	Mtu             int           `json:"mtu"`
	Type            string        `json:"type"`
	VswitchDto      Switch        `json:"vswitchDto"`
	PnicDto         interface{}   `json:"pnicDto"`
	PortDtos        []interface{} `json:"portDtos"`
	VMDtos          interface{}   `json:"vmDtos"`
	VnicDtos        interface{}   `json:"vnicDtos"`
	VmCount         int           `json:"vmcount"`
	VnicCount       int           `json:"vniccount"`
	ConnectMode     string        `json:"connectMode"`
	Description     string        `json:"description"`
	UplinkRate      int           `json:"uplinkRate"`
	UplinkBurst     int           `json:"uplinkBurst"`
	DownlinkRate    int           `json:"downlinkRate"`
	DownlinkBurst   int           `json:"downlinkBurst"`
	QosEnabled      bool          `json:"qosEnabled"`
	DataServiceType interface{}   `json:"dataServiceType"`
	UserVlan        interface{}   `json:"userVlan"`
	TpidType        interface{}   `json:"tpidType"`
	PermitDel       bool          `json:"permitDel"`
	Cidr            string        `json:"cidr"`
	CidrType        int           `json:"cidrType"`
	Gateway         string        `json:"gateway"`
	DhcpEnabled     bool          `json:"dhcpEnabled"`
	GatewayEnabled  bool          `json:"gatewayEnabled"`
	DNS             string        `json:"dns"`
	DataCenterDto   Datacenter    `json:"dataCenterDto"`
	NetworkTopoly   bool          `json:"networkTopoly"`
	UseTypes        string        `json:"useTypes"`
	StartIp         string        `json:"startIp"`
	EndIp           string        `json:"endIp"`
	Pools           []interface{} `json:"pools"`
	UsedByHbLink    bool          `json:"usedByHbLink"`
}
