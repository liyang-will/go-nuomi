package dirbuster

import (
  "context"
  "fmt"
  "go-nuomi/lib"
  "go-nuomi/lib/http"
  "go-nuomi/utils"
)

type Dirbuster struct {
  options *DirOptions
  libopts *lib.Options
  httpClient    *http.HTTPClient
}

func NewDirbuster(cont context.Context, libOptions *lib.Options,
  opts *DirOptions)(*Dirbuster, error){
    if libOptions == nil{
      return nil, fmt.Errorf("please provide valid global options")
    }
    if opts == nil{
      return nil, fmt.Errorf("please provide valid plugin options")
    }
    d := Dirbuster{
      options:opts,
      libopts:libOptions,
    }

    httpOpts := http.HTTPOptions{
      Proxy:opts.Proxy,
      FollowRedirect: opts.FollowRedirect,
      InsecureSSL:    opts.InsecureSSL,
      IncludeLength:  opts.IncludeLength,
      TimeOut:        opts.TimeOut,
      UserName:       opts.UserName,
      Password:       opts.Password,
      UserAgent:      opts.UserAgent,
      Headers:        opts.Headers,
    }
    h, err := http.NewHTTPClient(cont, &httpOpts)
    if err != nil {
      return nil, err
    }
    d.httpClient = h
    return &d, nil
}

func runDir()error{
  libOptions, pluginOpts, err := parseDirOptions()
  if err != nil{
    return fmt.Errorf("error on parsing arguments: %v", err)
  }

  plugin, err := NewDirbuster()

}


// 获取、解析DirOptions
func parseDirOptions()(*lib.Options, *DirOptions, error){
  libOptions := GetDirOptions()
  if libOptions == nil{
    return nil,nil, fmt.Errorf("invaild libOptions,value=%v",libOptions)
  }

  plugin := GetDirOptions()

  httpOpts, err := http.ParseCommonHTTPOptions()
  if err != nil{
    return nil, nil, fmt.Errorf("invaild HTTPSOptions,error=%v",err)
  }

  //plugin要密码干什么
  plugin.Password = httpOpts.Password
  plugin.URL = httpOpts.URL
  plugin.UserAgent = httpOpts.UserAgent
  plugin.UserName = httpOpts.UserName
  plugin.Proxy = httpOpts.Proxy
  plugin.Cookies = httpOpts.Cookies
  plugin.TimeOut = httpOpts.TimeOut
  plugin.FollowRedirect = httpOpts.FollowRedirect
  plugin.InsecureSSL = httpOpts.InsecureSSL
  plugin.Headers = httpOpts.Headers

  if plugin.Extensions != ""{
    ret, err := utils.ParseExtensions(plugin.Extensions)
    if err != nil{
      return nil, nil, fmt.Errorf("invalid value for extensions: %v", err)
    }
    plugin.ExtensionsParsed = ret
  }

  if plugin.StatusCodesBlacklist != ""{
    ret, err := utils.ParseStatusCodes(plugin.StatusCodesBlacklist)
    if err != nil{
      return nil, nil, fmt.Errorf("invalid value for statuscodesblacklist: %v",err)
    }
    plugin.StatusCodesBlacklistParsed = ret
  }else {
    ret, err := utils.ParseStatusCodes(plugin.StatusCodes)
    if err != nil {
      return nil, nil, fmt.Errorf("invalid value for statuscodes: %v", err)
    }
    plugin.StatusCodesParsed = ret
  }

  return libOptions, plugin, nil
}
