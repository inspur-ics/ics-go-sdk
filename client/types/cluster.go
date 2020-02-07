package types
type Cluster struct {
	Name    string    `json:"name"`    //"cluster"
	Id      string    `json:"id"`      //"8a878bda6f7012c7016f87e979120798"
	HostNum int       `json:"hostNum"` //1
	Tags    []Tag `json:"tags"`
}
type ClusterListRsp struct {
	Items []Cluster `json:"items"`
}