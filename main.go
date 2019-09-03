package main

import (
  "context"
  "fmt"
  "go-nuomi/dirbuster"
  "go-nuomi/lib"
  "go-nuomi/lib/http"
  "go-nuomi/nuomi"
)

func main(){
  n, r := nuomi.LoadConfig()
  if !r{
    return
  }

  mainContext, cancel := context.WithCancel(context.Background())
  defer cancel()

  fmt.Println(mainContext)

  lib.SetLibOptions(n.LibOption)
  dirbuster.SetDirOptions(n.DirOption)
  http.SetHTTPOptions(n.HttpOption)
}