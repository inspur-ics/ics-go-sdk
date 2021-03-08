module github.com/inspur-ics/ics-go-sdk

go 1.14

require (
	github.com/go-logr/logr v0.1.0
    github.com/go-resty/resty v1.12.0
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7
	k8s.io/klog v1.0.0
)

replace github.com/go-resty/resty v1.12.0 => gopkg.in/resty.v1 v1.12.0
