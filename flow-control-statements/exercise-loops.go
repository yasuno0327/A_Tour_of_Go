package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
  res := x - (x * x - x) / 2 * x
  return res
}

func main() {
  fmt.Println(Sqrt(2))
}
