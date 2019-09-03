package awvs

//妈的，先放一放吧
import (
  "bytes"
  "crypto/tls"
  "encoding/json"
  "fmt"
  "net/http"
  "net/http/cookiejar"
)

//curl -k --request POST --url
// https://127.0.0.1:13443/api/v1/targets
// --header "X-Auth: 1986ad8c0a5b3df4d7028d5f3c06e936cce6e2a0438ab4a69a83dd6b549d60b62"
// --header "content-type: application/json"
// --data {
// 	"address": "http://www.seu.edu.cn/",
// 	"description": "seu",
// 	"criticality": "10"
// }

type Target struct {
  Address string `json:"address"`
  Description string `json:"description"`
  //Critical [30], High [20], Normal [10], Low [0]
  Criticality string `json:"criticality"`
}


//忽略CA证书的验证
func NewTimeOutHttpReq()(*http.Client){
  cookieJar, _ := cookiejar.New(nil)
  client := &http.Client{
    Transport: &http.Transport{
      TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
    },
    Jar:cookieJar,
  }

  return client
}

func AddTarget(t Target){
//   req := gorequest.New().Timeout(10*time.Second)
//
//   resp, _, err := req.Post(AWVS_TARGET).Set("X-Auth",APIKEY).
//     TLSClientConfig(&tls.Config{InsecureSkipVerify:true}).Send(`
// {
// 	"address": "http://www.seu.edu.cn/",
// 	"description": "seu",
// 	"criticality": "10"
// }`).End()
//   if err != nil{
//     fmt.Println(err)
//     return
//   }
  c := NewTimeOutHttpReq()

  byte, err := json.Marshal(t)
  if err !=nil{
    fmt.Println(err)
    return
  }

  req, _ := http.NewRequest("POST", AWVS_TARGET, bytes.NewBuffer(byte))
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-Auth","1986ad8c0a5b3df4d7028d5f3c06e936cce6e2a0438ab4a69a83dd6b549d60b62")

  resp, err := c.Do(req)
  if err != nil{
    fmt.Println(err)
    return
  }

  if resp.StatusCode == 201{
    fmt.Println("状态返回201")
    return
  }

}



func addScan(){

}
