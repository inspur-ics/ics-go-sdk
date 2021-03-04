package vmtemplate

import (
	"github.com/inspur-ics/ics-go-sdk/client"
	"github.com/inspur-ics/ics-go-sdk/common"
)

type VMTemplateService struct {
	common.RestAPI
}

// NewVMTemplateService returns the session's virtual machine template service.
func NewVMTemplateService(c *client.Client) *VMTemplateService {
	vmt := VMTemplateService{
		common.RestAPI{
			RestAPITripper: c,
		},
	}

	return &vmt
}
