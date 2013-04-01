package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (rot13 rot13Reader) Read(s []byte) (int, error) {
  read, err := rot13.r.Read(s)

  rotate := func(s byte, lower, higher, amount int) byte {
    chr := int(s)
    if lower <= chr && chr <= higher {
      chr += amount
      if chr > higher { chr -= (higher - lower + 1) }
    }
    return byte(chr)
  }

  for i := 0; i < read; i++ {
    chr := s[i]
    chr = rotate(chr, 65, 90, 13)
    chr = rotate(chr, 97, 122, 13)
    s[i] = chr
  }
  return read, err
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
