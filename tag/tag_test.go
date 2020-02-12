package tag

import (
    "context"
    "fmt"
    icsgo "github.com/inspur-ics/ics-go-sdk"
    "testing"
)

func TestGetTag(t *testing.T) {
   icsConnection := &icsgo.ICSConnection{
       Username: "admin",
       Password: "admin@inspur",
       Hostname: "10.7.11.90",
       Port:     "443",
       Insecure: true,
   }
   ctx := context.Background()
   err := icsConnection.Connect(ctx)
   if err != nil {
       t.Fatal("Create ics connection error!")
   }

   tagClient := NewTagsService(icsConnection.Client)
   tag, err := tagClient.GetTag(ctx,"242c3f4c-fdd0-4ad0-b250-90b51581d5a9")
   if tag != nil {
       fmt.Printf("Tag json:%+v.", tag)
   }
}

func TestListAttachedTags(t *testing.T) {
    icsConnection := &icsgo.ICSConnection{
        Username: "admin",
        Password: "admin@inspur",
        Hostname: "10.7.11.90",
        Port:     "443",
        Insecure: true,
    }
    ctx := context.Background()
    err := icsConnection.Connect(ctx)
    if err != nil {
        t.Fatal("Create ics connection error!")
    }

    tagClient := NewTagsService(icsConnection.Client)
    tags, err := tagClient.ListAttachedTags(ctx,"HOST", "792b1e3f-8be6-43de-b8c7-27a0327dcd97")
    if tags != nil {
        fmt.Printf("tags:%+v", tags)
    }

}
