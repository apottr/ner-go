package main

import (
  "strings"
  "fmt"
  "regexp"
)

func stopwordsHeuristic(data string){
  
}

func capHeuristic(data string){
  lst := strings.Split(data," ")
  out := []string{}
  re := regexp.MustCompile("^[A-Z]")
  for i,v := range lst {
    fmt.Printf("%v,%v\n",i,v)
    if re.Match([]byte(v)){
      out = append(out,v)
    }
  }
  fmt.Printf("%q",out)
}

func main(){
  capHeuristic(`Bruce McLaren Motorsport Park (previously known as the Taupo Motorsport Park) is a motorsports circuit located in Broadlands Road, Taupo, New Zealand. It is owned and operated by MIT Development Ltd.`)
}
