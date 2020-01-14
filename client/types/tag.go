package types

type Tag struct {
    Name        string        `json:"tagName"`
    Description    string        `json:"description"`
}

type TagBinding struct {
    Tags             []Tag        `json:"tags"`
    SourceIds        []string     `json:"sourceIds"`
    TagSourceType    string       `json:"tagSourceType"`
}