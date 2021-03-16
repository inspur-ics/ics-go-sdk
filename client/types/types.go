package types

import "reflect"

type ManagedObjectReference struct {
    Type  string
    Value string
}

type ErrorMsg struct {
    Code              string        `json:"code"`
    Message           string        `json:"message"`
    Params            string        `json:"params"`
}

type Common struct {}

type PageReq struct {
    PageSize            int              `json:"pageSize"`
    CurrentPage         int              `json:"currentPage"`
    SortField           string           `json:"sortField"`
    Sort                string           `json:"sort"`
}

type PageResponse struct {
    TotalPage           int               `json:"totalPage"`
    CurrentPage         int               `json:"currentPage"`
    TotalSize           int               `json:"totalSize"`
    Items               []interface{}     `json:"items"`
}

type TreeItem struct {
    ID         string     `json:"id"`
    Text       string     `json:"text"`
    IconCls    string     `json:"iconCls"`
    Checked    bool       `json:"checked"`
    ViewID     string     `json:"viewId"`
    Children   []TreeItem `json:"children"`
}

type ICSApi struct {
    Api    string        `json:"api"`
    Token  bool          `json:"token"`
}

// ICS LOGIN FORM
type Login struct {
    Username         string        `json:"username"`
    Password         string        `json:"password"`
    Domain           string        `json:"domain"`
    Locale           string        `json:"locale"`
}

// ICS LOGIN RESPONSE BODY
type LoginResponse struct {
    UserId           string        `json:"userId"`
    SessonId         string        `json:"sessonId"`
    Validated        bool          `json:"validated"`
    Message          string        `json:"message"`
    Username         string        `json:"username"`
    Password         string        `json:"password"`
    Captcha          string        `json:"captcha"`
    Locale           string        `json:"locale"`
    Domain           string        `json:"domain"`
    Remains          int           `json:"remains"`
    IP               string        `json:"ip"`
    Operator         string        `json:"operator"`
    LoginTime        string        `json:"loginTime"`
    CreateDate       string        `json:"createDate"`
    RoleType         string        `json:"roleType"`
    Themes           string        `json:"themes"`
}

// ICS LOGIN RESPONSE BODY
type UserSession struct {
    UserId           string        `json:"userId"`
    Username         string        `json:"username"`
    SessonId         string        `json:"sessonId"`
    RoleType         string        `json:"roleType"`
    Locale           string        `json:"locale"`
    IP               string        `json:"ip"`
    Themes           string        `json:"themes"`
    CreateDate       string        `json:"createDate"`
    LoginTime        string        `json:"loginTime"`
}

type ServiceContent struct {
}

type DynamicData struct {
}

func init() {
    t["DynamicData"] = reflect.TypeOf((*DynamicData)(nil)).Elem()
}

type OptionType struct {
    DynamicData

    ValueIsReadonly *bool `json:"valueIsReadonly"`
}

func init() {
    t["OptionType"] = reflect.TypeOf((*OptionType)(nil)).Elem()
}

type OptionValue struct {
    DynamicData

    Key   string  `json:"key"`
    Value AnyType `json:"value"`
}

func init() {
    t["OptionValue"] = reflect.TypeOf((*OptionValue)(nil)).Elem()
}