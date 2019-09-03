package lib

import (
  "context"
  "fmt"
  "time"
)

const(
  VERSION = "0.0.1"
)

var libOption *Options

type Options struct {
  Coroutines int
  Wordlist string
  OutputFilename string
  NoStatus bool
  NoProgress bool
  Quiet bool
  WildcardForced bool
  Verbose bool
  Delay time.Duration
}

var mainContext context.Context

func SetLibOptions(o Options){
  libOption = &o
  fmt.Printf("libOption=%#v\n",libOption)
}

func GetLibOptions()*Options{
  return libOption
}