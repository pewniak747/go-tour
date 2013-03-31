package main

import (
  "code.google.com/p/go-tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  words := strings.Fields(s)
  wordsMap := make(map[string]int)
  for _, v := range words {
    wordsMap[v]++
  }
  return wordsMap
}

func main() {
  wc.Test(WordCount)
}
