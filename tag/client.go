package tag

import (
    "github.com/inspur-ics/ics-go-sdk/client"
    "github.com/inspur-ics/ics-go-sdk/common"
)

type TagsService struct {
    common.RestAPI
}

// NewTagService returns the session's tag service.
func NewTagsService(c *client.Client) *TagsService {
    tags := TagsService{
        common.RestAPI{
            RestAPITripper: c,
        },
    }

    return &tags
}