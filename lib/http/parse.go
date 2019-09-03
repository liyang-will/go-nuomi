package http

import (
  "fmt"
  "regexp"
  "strconv"
  "strings"
)

func ParseCommonHTTPOptions() (*HTTPOptions, error) {
  options := GetHTTPOptions()

  if !strings.HasPrefix(options.URL, "http") {
    // check to see if a port was specified
    re := regexp.MustCompile(`^[^/]+:(\d+)`)
    match := re.FindStringSubmatch(options.URL)

    if len(match) < 2 {
      // no port, default to http on 80
      options.URL = fmt.Sprintf("http://%s", options.URL)
    } else {
      port, err2 := strconv.Atoi(match[1])
      if err2 != nil || (port != 80 && port != 443) {
        return options, fmt.Errorf("url scheme not specified")
      } else if port == 80 {
        options.URL = fmt.Sprintf("http://%s", options.URL)
      } else {  //443 ssl
        options.URL = fmt.Sprintf("https://%s", options.URL)
      }
    }
  }

  // 若不赋值，会怎样
  // options.Cookies, err = cmd.Flags().GetString("cookies")
  // if err != nil {
  //   return options, fmt.Errorf("invalid value for cookies: %v", err)
  // }
  //
  // options.Username, err = cmd.Flags().GetString("username")
  // if err != nil {
  //   return options, fmt.Errorf("invalid value for username: %v", err)
  // }
  //
  // options.Password, err = cmd.Flags().GetString("password")
  // if err != nil {
  //   return options, fmt.Errorf("invalid value for password: %v", err)
  // }
  //
  // options.UserAgent, err = cmd.Flags().GetString("useragent")
  // if err != nil {
  //   return options, fmt.Errorf("invalid value for useragent: %v", err)
  // }
  //
  // options.Proxy, err = cmd.Flags().GetString("proxy")
  // if err != nil {
  //   return options, fmt.Errorf("invalid value for proxy: %v", err)
  // }
  //
  // options.Timeout, err = cmd.Flags().GetDuration("timeout")
  // if err != nil {
  //   return options, fmt.Errorf("invalid value for timeout: %v", err)
  // }
  //
  // options.FollowRedirect, err = cmd.Flags().GetBool("followredirect")
  // if err != nil {
  //   return options, fmt.Errorf("invalid value for followredirect: %v", err)
  // }
  //
  // options.InsecureSSL, err = cmd.Flags().GetBool("insecuressl")
  // if err != nil {
  //   return options, fmt.Errorf("invalid value for insecuressl: %v", err)
  // }
  //
  // headers, err := cmd.Flags().GetStringArray("headers")
  // if err != nil {
  //   return options, fmt.Errorf("invalid value for headers: %v", err)
  // }

  return options, nil
}

