package main

import (
  "fmt"
  "time"
)

type ErrNegativeSqrt struct {
  when time.Time
  what string
}

func (e ErrNegativeSqrt) Error() string

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return
  }
  return 0, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
