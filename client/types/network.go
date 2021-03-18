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
	ControllerIP     interface{} `json:"controllerIP"`
	DataCenterDto    Datacenter  `json:"dataCenterDto"`
	HostDtos         interface{} `json:"hostDtos"`
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
	Maxvfs           int         `json:"maxvfs"`
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
}

type Network struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	ResourceID      string        `json:"resourceId"`
	Vlan            int           `json:"vlan"`
	VlanFlag        bool          `json:"vlanFlag"`
	Mtu             interface{}   `json:"mtu"`
	Type            string        `json:"type"`
	VswitchDto      Switch        `json:"vswitchDto"`
	PortDtos        []interface{} `json:"portDtos"`
	VMDtos          interface{}   `json:"vmDtos"`
	VnicDtos        interface{}   `json:"vnicDtos"`
	Vmcount         int           `json:"vmcount"`
	Vniccount       int           `json:"vniccount"`
	ConnectMode     string        `json:"connectMode"`
	Description     interface{}   `json:"description"`
	UplinkRate      int           `json:"uplinkRate"`
	UplinkBurst     int           `json:"uplinkBurst"`
	DownlinkRate    int           `json:"downlinkRate"`
	DownlinkBurst   int           `json:"downlinkBurst"`
	QosEnabled      bool          `json:"qosEnabled"`
	DataServiceType interface{}   `json:"dataServiceType"`
	UserVlan        interface{}   `json:"userVlan"`
	TpidType        interface{}   `json:"tpidType"`
	PermitDel       bool          `json:"permitDel"`
	Cidr            interface{}   `json:"cidr"`
	Gateway         interface{}   `json:"gateway"`
	DhcpEnabled     bool          `json:"dhcpEnabled"`
}
