package http

import (
  "fmt"
  "strings"
  "testing"
)

func TestString(t *testing.T){
  fmt.Printf("==1 %q\n", strings.SplitN("/home/m_ta/src", "/", 1))

  fmt.Printf("==2 %q\n", strings.SplitN("111/home/m_ta/src", "/", 2))  //["/" "home/" "m_ta/" "src"]
  fmt.Printf("==3 %q\n", strings.SplitN("/home/m_ta/src", "/", -1)) //["" "home" "m_ta" "src"]
  fmt.Printf("==4 %q\n", strings.SplitN("home,m_ta,src", ",", 2))   //["/" "home/" "m_ta/" "src"]

  fmt.Printf("==5 %q\n", strings.SplitN("#home#m_ta#src", "#", -1)) //["/" "home/" "m_ta/" "src"]
}