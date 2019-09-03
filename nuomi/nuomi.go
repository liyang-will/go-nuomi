package nuomi

import (
  "go-nuomi/dirbuster"
  "go-nuomi/lib"
  "go-nuomi/lib/http"
)

type NuoMiOption struct {
  LibOption lib.Options
  DirOption dirbuster.DirOptions
  HttpOption http.HTTPOptions
}



