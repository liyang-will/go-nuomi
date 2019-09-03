package nuomi

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

//读取测试数据的文件内容
func readTestDataFile() []byte{
  path := "./nuomi/config.json"
  fi, err := os.Open(path)
  if err != nil {
    fmt.Println(err)
    panic(err)
  }
  defer fi.Close()
  fd, err := ioutil.ReadAll(fi)
  return fd
}

func LoadConfig()(*NuoMiOption,bool){
  NuoMiOption := &NuoMiOption{}
  data := readTestDataFile()
  err := json.Unmarshal(data, NuoMiOption)
  if err != nil {
    fmt.Errorf("LibOptionLoad error:%v",err)
    return NuoMiOption,false
  }
  return NuoMiOption,true
}