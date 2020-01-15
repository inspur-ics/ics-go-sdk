package types

type DatastoreInfo struct {
    ID                  string      `json:"id"`
    Name                string      `json:"name"`
}
type Datastore struct {
    DataStoreType       string      `json:"dataStoreType"`
    ID                  string      `json:"id"`
    Name                string      `json:"name"`
    MountPath           string      `json:"mountPath"`
    Capacity            float64     `json:"capacity"`
    UsedCapacity        float64     `json:"usedCapacity"`
    AvailCapacity       float64     `json:"availCapacity"`
    DataCenterID        string      `json:"dataCenterId"`
    HostID              interface{} `json:"hostId"`
    MountStatus         interface{} `json:"mountStatus"`
    HostIP              interface{} `json:"hostIp"`
    UUID                string      `json:"uuid"`
    AbsolutePath        interface{} `json:"absolutePath"`
    DataCenterName      interface{} `json:"dataCenterName"`
    DataCenterOrHostDto struct {
        DataCenterOrHost string      `json:"dataCenterOrHost"`
        DataCenterName   interface{} `json:"dataCenterName"`
        HostName         interface{} `json:"hostName"`
        Status           interface{} `json:"status"`
    } `json:"dataCenterOrHostDto"`
    BlockDeviceDto    interface{} `json:"blockDeviceDto"`
    DataCenterDto     interface{} `json:"dataCenterDto"`
    HostNumbers       int         `json:"hostNumbers"`
    VMNumbers         int         `json:"vmNumbers"`
    VolumesNumbers    int         `json:"volumesNumbers"`
    VMTemplateNumbers int         `json:"vmTemplateNumbers"`
    Tags              interface{} `json:"tags"`
    MaxSlots          int         `json:"maxSlots"`
    Creating          bool        `json:"creating"`
    StorageBackUp     bool        `json:"storageBackUp"`
    ExtensionType     string      `json:"extensionType"`
    CanBeImageStorage bool        `json:"canBeImageStorage"`
    BlockDeviceUUID   interface{} `json:"blockDeviceUuid"`
    OpHostIP          interface{} `json:"opHostIp"`
    IsMount           interface{} `json:"isMount"`
    HostDto           interface{} `json:"hostDto"`
    ScvmOn            bool        `json:"scvmOn"`
}