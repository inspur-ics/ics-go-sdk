package types

type Vapp struct {
	ID             string         `json:"id"`
	DataCenterID   string         `json:"dataCenterId"`
	Name           string         `json:"name"`
	CpuCount       string         `json:"cpucount"`
	Memory         string         `json:"memory"`
	MemoryInByte   int            `json:"memoryInByte,omitempty"`
	ActiveVMCount  string         `json:"activevircount"`
	VMCount        string         `json:"vircount"`
	State          string         `json:"state"`
	Health         string         `json:"health,omitempty"`
	Product        string         `json:"product,omitempty"`
	Version        string         `json:"version,omitempty"`
	Supplier       string         `json:"supplier,omitempty"`
	CanPowerOn     bool           `json:"canPowerOn"`
	CanPowerOff    bool           `json:"canPowerOff"`
	CanRestart     bool           `json:"canRestart"`
	DataCenterName string         `json:"dataCenterName"`
	StatusCount    map[string]int `json:"statusCount"`
	description    string         `json:"description,omitempty"`
	VappType       interface{}    `json:"vappType"`
	NodeCount      interface{}    `json:"nodeCount"`
	VmsToAdd       interface{}    `json:"vmsToAdd,omitempty"`
}

type VappListRsp struct {
	Items []Vapp `json:"items"`
}

type VappCreateReq struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	DataCenterID   string `json:"dataCenterId"`
	DataCenterName string `json:"dataCenterNamei,omitempty"`
}
