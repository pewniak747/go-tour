package main

import (
  "fmt"
  "math"
)

func MySqrt(x float64) float64 {
  z, prev_z := x, x+1
  eps := math.Pow(10, -10)
  for math.Abs(z - prev_z) - eps > 0 {
    prev_z = z
    z = z - (math.Pow(z, 2)-x)/(2*z)
  }
  return z
}

func main() {
  var x float64 = 2
  fmt.Println(MySqrt(x))
  fmt.Println(math.Sqrt(x))
}
