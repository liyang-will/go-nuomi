package http

import (
  "context"
  "crypto/tls"
  "fmt"
  "go-nuomi/lib"
  "net/http"
  "net/url"
  "time"
)

var httpOption *HTTPOptions

type HTTPHeader struct {
  Name string
  Value string
}

type HTTPClient struct {
  client *http.Client
  context context.Context
  userAgent string
  defaultUserAgent string
  username string
  password string
  headers []HTTPHeader
  includeLength bool
}

type HTTPOptions struct {
  Password       string
  URL            string
  UserAgent      string
  UserName       string
  Proxy          string
  Cookies        string
  Headers        []HTTPHeader
  TimeOut        time.Duration
  FollowRedirect bool
  InsecureSSL    bool
  IncludeLength bool
}

func NewHTTPClient(c context.Context, opt *HTTPOptions)(*HTTPClient, error){
  var proxyURLFunc func(*http.Request)(*url.URL, error)
  var client HTTPClient
  proxyURLFunc = http.ProxyFromEnvironment

  if opt == nil{
    return nil, fmt.Errorf("options is nil")
  }

  if opt.Proxy != ""{
    proxyURL, err := url.Parse(opt.Proxy)
    if err != nil{
      return nil, fmt.Errorf("proxy URL is invalid (%v)", err)
    }
    proxyURLFunc = http.ProxyURL(proxyURL)
  }

  var redirectFunc func(req *http.Request, via []*http.Request) error
  if !opt.FollowRedirect{
    redirectFunc = func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
    }
  }else {
    redirectFunc = nil
  }

  client.client = &http.Client{
    Timeout: opt.TimeOut,
    CheckRedirect:redirectFunc,
    Transport:&http.Transport{
      Proxy:proxyURLFunc,
      MaxIdleConns:100,
      MaxIdleConnsPerHost:100,
      TLSClientConfig:&tls.Config{
        InsecureSkipVerify:opt.InsecureSSL,
      },
    }}
  client.context = c
  client.username= opt.UserName
  client.password = opt.Password
  client.includeLength = opt.IncludeLength
  client.userAgent = opt.UserAgent
  client.defaultUserAgent = DirDefaultUserAgent()
  client.headers = opt.Headers
  return &client,nil
}

func GetHTTPOptions() *HTTPOptions{
  return httpOption
}

func SetHTTPOptions(h HTTPOptions){
  httpOption = &h
  fmt.Printf("dirOptions=%#v\n",httpOption)
}

// DefaultUserAgent returns the default user agent to use in HTTP requests
func DirDefaultUserAgent() string {
  return fmt.Sprintf("nuomi/%s", lib.VERSION)
}