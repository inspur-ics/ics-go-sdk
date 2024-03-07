package types

type Tag struct {
	ID          string `json:"id"`
	Name        string `json:"tagName"`
	UserName    string `json:"userName"`
	CreateTime  string `json:"createTime"`
	Description string `json:"description"`
	Checked     bool   `json:"checked"`
}

type TagBinding struct {
	Tags          []Tag    `json:"tags"`
	SourceIds     []string `json:"sourceIds"`
	TagSourceType string   `json:"tagSourceType"`
}
