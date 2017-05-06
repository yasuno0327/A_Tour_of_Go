package main

import "fmt"

type Vertex struct {
  lat, long float64
}

var m map[string]Vertex

func main() {
  m = make(map[string]Vertex)
  m["bell labs"] = Vertex{
    40.68433, -74.39677,
  }
  fmt.Println(m["bell labs"])
}
