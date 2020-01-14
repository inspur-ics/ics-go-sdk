package tag

import (
    "context"

    "github.com/inspur-ics/ics-go-sdk/client/types"
)

//FIXME TODO.WANGYONGCHAO
// ListAttachedTags fetches the array of tag IDs attached to the given object.
func (t *TagsService) ListAttachedTags(ctx context.Context, targetType string, ref string) ([]string, error) {
    var res []string
    return res, nil
}

//FIXME TODO.WANGYONGCHAO
func (c *TagsService) GetTag(ctx context.Context, id string) (*types.Tag, error) {
    var res types.Tag
    return &res, nil
}