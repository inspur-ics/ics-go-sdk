package storage

import (
	"context"
	"encoding/json"
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
		Password: "Cloud@s1",
		Hostname: "10.49.34.161",
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
		for i := 0; i < len(storagelist); i++ {
			storageJson, _ := json.MarshalIndent(storagelist[i], "", "\t")
			t.Logf("Storage Info: %s\n", string(storageJson))
		}
	} else {
		t.Errorf("Failed to get storage list. Error: %v\n", err)
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
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}
	storageClient = NewStorageService(icsConnection.Client)
	storageinfo, err := storageClient.GetStorageInfoById(ctx, "8ab1a2968205c11601825d5599e615f7")
	if err != nil {
		fmt.Println(err.Error())
		t.Fatal("Get storage info failed!")
	} else {
		fmt.Println(storageinfo.Name)
		fmt.Println(storageinfo.MountPath)
		fmt.Println(storageinfo.AllocPolicy)
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

func TestGetImageFileList(t *testing.T) {
	fmt.Println("********************TestGetImageFileList**************")
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	imageDatastoreId := "8ab1a2218dca60bc018dca6361be0016"
	storageClient = NewStorageService(icsConnection.Client)
	imageList, err := storageClient.GetImageFileList(ctx, imageDatastoreId)
	if err != nil {
		t.Errorf("Failed to get image file list. Error: %v", err)
	} else {
		for _, imageInfo := range imageList {
			t.Logf("Image Name:%s", imageInfo.Name)
			t.Logf("Image Path:%s", imageInfo.Path)
		}
	}
}

func TestGetImageFileInfoByName(t *testing.T) {
	fmt.Println("********************TestGetImageFileInfoByName**************")
	ctx = context.Background()
	err := icsConnection.Connect(ctx)
	if err != nil {
		t.Fatal("Create ics connection error!")
	}

	imageName := "rh7.6_x86_mini.ova"
	storageClient = NewStorageService(icsConnection.Client)
	imageInfo, err := storageClient.GetImageFileInfoByName(ctx, imageName)
	if err != nil {
		t.Errorf("Failed to get image file info. Error: %v", err)
	} else {
		t.Logf("Image Name:%s", imageInfo.Name)
		t.Logf("Image Path:%s", imageInfo.Path)
	}
}
