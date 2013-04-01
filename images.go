package main

import (
  "code.google.com/p/go-tour/pic"
  "image"
  "image/color"
)

type Image struct {
  Width int
  Height int
}

func (i Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
  z := uint8(x * y)
  return color.RGBA{z, z, 255, 255}
}

func main() {
  m := Image{100, 80}
  pic.ShowImage(m)
}
