package main

import (
  "fmt"
  "math"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %v", float64(err))
}

func MySqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }
  z, prev_z := x, x+1
  eps := math.Pow(10, -10)
  for math.Abs(z - prev_z) - eps > 0 {
    prev_z = z
    z = z - (math.Pow(z, 2)-x)/(2*z)
  }
  return z, nil
}

func main() {
  var x float64 = 2
  fmt.Println(MySqrt(x))
  fmt.Println(MySqrt(-x))
}
