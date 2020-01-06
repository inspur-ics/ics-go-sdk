package types

type ManagedObjectReference struct {
	Type  string
	Value string
}

type ErrorMsg struct {
	Code              string        `json:"code"`
	Message           string        `json:"message"`
	Params            string        `json:"params"`
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