package mysql

type Domain struct {
  Id string `json:"id"`
  //域名，有可能是真实IP
  DomainName string `json:"domain_name"`
  //子域名
  SubDomains []string `json:"sub_domains"`
  //真实IP
  ReadIP string `json:"read_ip"`
  //有效C段
  ValidC []string `json:"valid_c"`
  //nmap开放的端口信息，暂时是string吧
  nmap string
}

func getDomainNameById(){

}

func AddDomainData(d *Domain){

}

