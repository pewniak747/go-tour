package main

import (
  "net/http"
  "fmt"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, s)
}

type Struct struct {
  Greeting string
  Punct    string
  Who      string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, fmt.Sprintf("%s says%s %s", s.Who, s.Punct, s.Greeting))
}

func main() {
  http.Handle("/string", String("Hello HTTP world!"))
  http.Handle("/struct", Struct{"Hello HTTP world!", ":", "pewniak747"})
  http.ListenAndServe("localhost:4000", nil)
}
