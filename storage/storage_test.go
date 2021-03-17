package storage

import (
	"context"
	"fmt"
	icsgo "github.com/inspur-ics/ics-go-sdk"
	"github.com/inspur-ics/ics-go-sdk/client/types"
	"testing"
)

var (
	ctx           context.Context
	storageClient *StorageService
	icsConnection = icsgo.ICSConnection{
		Username: "admin",
		Password: "admin@inspur",
		Hostname: "10.7.11.90",
		Port:     "443",
		Insecure: true,
	}
)

func TestStorageList(t *testing.T) {
	fmt.Println("********************TestStorageList**************")
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	storageClient = NewStorageService(icsConnection.Client)
	storagelist, err := storageClient.GetStoragesList(ctx)
	if storagelist != nil {
		for i := 0; i < len(storagelist); {
			fmt.Println(storagelist[i].Name)
			fmt.Println(storagelist[i].MountPath)
			i++
		}
	} else {
		fmt.Println(err.Error())
	}
}

func TestStorageInfoByName(t *testing.T) {
	fmt.Println("********************TestStorageInfoByName**************")
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	storageClient = NewStorageService(icsConnection.Client)
	storageInfo, err := storageClient.GetStorageInfoByName(ctx, "local")
	if storageInfo != nil {
		fmt.Println(storageInfo.Name)
		fmt.Println(storageInfo.MountPath)
	} else {
		fmt.Println("DataStore not found.")
	}
}

func TestStorageInfoById(t *testing.T) {
	fmt.Println("********************TestStorageInfoById**************")
	storageinfo, err := storageClient.GetStorageInfoById(ctx, "8a878bda6f6f3ca4016f6f458b50003c")
	if storageinfo != nil {
		fmt.Println(storageinfo.Name)
		fmt.Println(storageinfo.MountPath)
	} else {
		fmt.Println(err.Error())
	}
}

func TestStoragePageList(t *testing.T) {
	fmt.Println("********************TestStoragePageList**************")
	pakg := types.PageReq{
		1000,
		1,
		"",
		"desc",
	}
	req := &types.StoragePageReq{
		pakg,
	}
	storagePageList, err := storageClient.GetStoragePageList(req)
	if storagePageList != nil {
		storageList := storagePageList.Items
		if storageList != nil {
			for i := 0; i < len(storageList); {
				fmt.Println(storageList[i].Name)
				fmt.Println(storageList[i].MountPath)
				i++
			}
		} else {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}
