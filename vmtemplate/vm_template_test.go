package vmtemplate

import (
	"context"
	"fmt"
	icsgo "github.com/inspur-ics/ics-go-sdk"
	//"github.com/inspur-ics/ics-go-sdk/client/types"
	"testing"
)

func TestGetVMTemplateList(t *testing.T) {
	icsConnection := &icsgo.ICSConnection{
		Username: "admin",
		Password: "Cloud@s1",
		Hostname: "10.48.50.13",
		Port:     "443",
		Insecure: true,
	}
	ctx := context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	vmtService := NewVMTemplateService(icsConnection.Client)
	vmtList, err := vmtService.GetVMTemplateList(ctx)
	if err != nil {
		fmt.Printf("Failed to get VMTemplateList. Error: %v", err)
	} else {
		fmt.Printf("VMTemplateList: %+v", vmtList)
	}
}
