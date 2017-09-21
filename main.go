package main

import (
  "strings"
  "fmt"
  "regexp"
  "io/ioutil"
  "sort"
)

func stopwordsHeuristic(data string){
  sw,err := ioutil.ReadFile("stop-word-list.txt")
  out := []string{}
  if err != nil {
    panic(err)
  }
  lines := strings.Split(string(sw),"\n")
  for _,word := range strings.Split(data," ") {
    fmt.Println(word)
    if (idx := sort.SearchStrings(lines,word);
        idx != len(lines) && lines[idx] == word) {
      out = append(out,word)
    }
  }
  fmt.Printf("%q",out)
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
  stopwordsHeuristic(`Bruce McLaren Motorsport Park (previously known as the Taupo Motorsport Park) is a motorsports circuit located in Broadlands Road, Taupo, New Zealand. It is owned and operated by MIT Development Ltd.`)
}
