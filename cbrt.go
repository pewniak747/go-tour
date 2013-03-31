package main

import (
  "fmt"
  "math"
  "math/cmplx"
)

func MyCbrt(x complex128) complex128 {
  z, prev_z := x, x+1
  eps := math.Pow(10, -10)
  for float64(cmplx.Abs(z - prev_z)) - eps > 0 {
    prev_z = z
    z = z - (cmplx.Pow(z, 3) - x)/(3*cmplx.Pow(z, 3))
  }
  return z
}

func main() {
  var x complex128 = 2
  fmt.Println(MyCbrt(x))
  fmt.Println(cmplx.Pow(x, 1.0/3))
}
