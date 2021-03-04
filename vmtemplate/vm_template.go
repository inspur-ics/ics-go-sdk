package vmtemplate

import (
	"context"
	"github.com/inspur-ics/ics-go-sdk/client/methods"
	"github.com/inspur-ics/ics-go-sdk/client/types"
)

func (v *VMTemplateService) GetVMTemplateList(ctx context.Context) ([]types.VMTemplate, error) {
	vmTemplateList, err := methods.GetVMTemplateList(ctx, v.RestAPITripper)
	return vmTemplateList.Items, err
}
