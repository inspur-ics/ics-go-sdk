package restful

import (
	"encoding/json"

	"context"
	"crypto/tls"
	"github.com/inspur-ics/ics-go-sdk/client/types"
	"net/url"
	"regexp"
	"sync"
	"time"

	"github.com/go-resty/resty"
)

const (
	SessionCookieName = "JSESSOINID"
)

type Error struct {
	/* variables */
}

type RestAPITripper interface {
	GetTrip(ctx context.Context, api types.ICSApi, req, res interface{}) (*Response, error)
	PostTrip(ctx context.Context, api types.ICSApi, req, res interface{}) (*Response, error)
	PutTrip(ctx context.Context, api types.ICSApi, req, res interface{}) (*Response, error)
	DeleteTrip(ctx context.Context, api types.ICSApi, req, res interface{}) (*Response, error)
}

type Client struct {
	HttpClient *resty.Client

	u *url.URL
	k bool // Named after curl's -k flag

	hostsMu sync.Mutex
	hosts   map[string]string

	Namespace string // ics internal
	Version   string //  ics api version

	Authorization    string
}

var schemeMatch = regexp.MustCompile(`^\w+://`)

// ParseURL is wrapper around url.Parse, where Scheme defaults to "https" and Path defaults to "/"
func ParseURL(s string) (*url.URL, error) {
	var err error
	var u *url.URL

	if s != "" {
		// Default the scheme to https
		if !schemeMatch.MatchString(s) {
			s = "https://" + s
		}

		u, err = url.Parse(s)
		if err != nil {
			return nil, err
		}

		// Default the path to /sdk
		if u.Path == "" {
			u.Path = "/"
		}

		if u.User == nil {
			u.User = url.UserPassword("", "")
		}
	}

	return u, nil
}

func ParseURI(url *url.URL, uri string) string {
	apiPath := url.String()
	apiPath += uri
	return apiPath
}

func NewClient(u *url.URL, insecure bool) *Client {
	c := Client{
		u: u,
		k: insecure,
	}

	c.HttpClient = resty.New()

	c.HttpClient.SetHeader("Content-Type", "application/json; charset=utf-8").
		SetHeader("Accept", "application/json; charset=utf-8").
		SetHeader("version", "5.6")

	c.HttpClient.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: insecure })

	c.HttpClient.SetTimeout(1 * time.Minute)

	c.hosts = make(map[string]string)

	// Remove user information from a copy of the URL
	c.u = c.URL()
	c.u.User = nil

	return &c
}

func (c *Client) URL() *url.URL {
	urlCopy := *c.u
	return &urlCopy
}

func (c *Client) GetToken() string {
	return c.Authorization
}

func (c *Client) SetToken(token string) {
	c.Authorization = token
}

func (c *Client) generateRequest(api types.ICSApi) *resty.Request {
	client := c.HttpClient
	request := client.R()
	if api.Token {
		if len(c.GetToken()) == 0 {
			request.SetHeader("Authorization", "anonymous")
		} else {
			request.SetHeader("Authorization", c.GetToken())
		}
	}
	return request
}

func (c *Client) GetTrip(ctx context.Context, api types.ICSApi, req, res interface{}) (*Response, error) {
	var errorParam    Response

	apiPath := ParseURI(c.URL(), api.Api)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return &errorParam, err
	}

	var getReq map[string]string
	err = json.Unmarshal([]byte(reqBody), &getReq)
	if err != nil {
		return &errorParam, err
	}

	resp, err := c.generateRequest(api).
		SetQueryParams(getReq).
		Get(apiPath)

	response := Response{
		resp.RawResponse,
		resp.Body(),
	}
	return &response, err
}

func (c *Client) PostTrip(ctx context.Context, api types.ICSApi, req, res interface{}) (*Response, error) {
	var errorParam    Response

	apiPath := ParseURI(c.URL(), api.Api)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return &errorParam, err
	}

	// POST JSON string
	resp, err := c.generateRequest(api).
		SetBody(reqBody).
		Post(apiPath)

	response := Response{
		resp.RawResponse,
		resp.Body(),
	}
	return &response, err
}

func (c *Client) PutTrip(ctx context.Context, api types.ICSApi, req, res interface{}) (*Response, error) {
	var errorParam    Response

	apiPath := ParseURI(c.URL(), api.Api)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return &errorParam, err
	}

	// Just one sample of PUT, refer POST for more combination
	resp, err := c.generateRequest(api).
		SetBody(reqBody).
		SetError(&Error{}).
		Put(apiPath)

	response := Response{
		resp.RawResponse,
		resp.Body(),
	}
	return &response, err
}

func (c *Client) DeleteTrip(ctx context.Context, api types.ICSApi, req, res interface{}) (*Response, error) {
	var errorParam    Response

	apiPath := ParseURI(c.URL(), api.Api)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return &errorParam, err
	}
	// DELETE a articles with payload/body as a JSON string
	resp, err := c.generateRequest(api).
		SetError(&Error{}).
		SetBody(reqBody).
		Delete(apiPath)

	response := Response{
		resp.RawResponse,
		resp.Body(),
	}
	return &response, err
}
