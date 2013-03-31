package main

import "code.google.com/p/go-tour/pic"

func Pic(x, y int) [][]uint8 {
  a := make([][]uint8, y)
  for i := 0; i < y; i++ {
    a[i] = make([]uint8, x)
    for l := 0; l < x; l++ {
      a[i][l] = uint8(i*l)
    }
  }
  return a
}

func main() {
  pic.Show(Pic)
}
