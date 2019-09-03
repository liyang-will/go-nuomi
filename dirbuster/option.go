package dirbuster

import (
  "fmt"
  "go-nuomi/lib"
  "go-nuomi/lib/http"
)

var dirOptions *DirOptions

type DirOptions struct {
  http.HTTPOptions
  Extensions  string
  ExtensionsParsed lib.StringSet

  StatusCodes string
  StatusCodesParsed lib.IntSet
  StatusCodesBlacklist string
  StatusCodesBlacklistParsed lib.IntSet

  UseSlash bool
  WildcardForced bool
  Includelength bool
  //指定呈现完整URL的扩展模式。
  Expanded bool
  //“无状态”模式，禁用结果状态代码的输出。
  NoStatus bool
}

func NewOptionsDir() *DirOptions {
  return &DirOptions{
    StatusCodesBlacklistParsed:lib.NewIntSet(),
    StatusCodesParsed:lib.NewIntSet(),
    ExtensionsParsed:lib.NewStringSet(),
  }
}

func SetDirOptions(o DirOptions){
  dirOptions = &o
  dirOptions.StatusCodesBlacklistParsed = lib.NewIntSet()
  dirOptions.StatusCodesParsed = lib.NewIntSet()
  dirOptions.ExtensionsParsed = lib.NewStringSet()
  fmt.Printf("dirOptions=%#v\n",dirOptions)
}

func GetDirOptions()*DirOptions{
  return dirOptions
}