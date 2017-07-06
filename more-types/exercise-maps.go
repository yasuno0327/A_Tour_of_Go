package main

import (
  "golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
  return map[string]int{"x": 1}
}

func main () {
  wc.Test(wordCount)
}
