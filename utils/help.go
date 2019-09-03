package utils

import (
  "fmt"
  "go-nuomi/lib"
  "strconv"
  "strings"
)

//增加扩展
func ParseExtensions(extensions string)(lib.StringSet,error){
  if extensions == ""{
    return lib.StringSet{},fmt.Errorf("invalid extension string provided")
  }

  ret := lib.NewStringSet()
  exts := strings.Split(extensions, ",")
  for _, e := range exts{
    e = strings.TrimSpace(e)
    ret.Add(strings.TrimPrefix(e,"."))
  }
  return ret,nil
}

//白名单
func ParseStatusCodes(statuscodes string) (lib.IntSet, error) {
  if statuscodes == "" {
    return lib.IntSet{}, fmt.Errorf("invalid status code string provided")
  }
  ret := lib.NewIntSet()
  for _, c := range strings.Split(statuscodes, ",") {
    c = strings.TrimSpace(c)
    i, err := strconv.Atoi(c)
    if err != nil {
      return lib.IntSet{}, fmt.Errorf("invalid status code given: %s", c)
    }
    ret.Add(i)
  }
  return ret, nil
}

