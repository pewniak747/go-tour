package main

import "fmt"

func fibbonacci() func() int {
  a, b := 0, 1
  return func() int {
    c := a + b
    a = b
    b = c
    return c
  }
}

func main() {
  f := fibbonacci()
  for i:= 0; i < 10; i++ {
    fmt.Println(f())
  }
}
