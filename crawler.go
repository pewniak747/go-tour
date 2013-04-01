package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "regexp"
  "strings"
  "os"
)

type Fetcher interface {
  Fetch(url string) (response Response, err error)
}

type Response interface {
  Url() string
  Body() string
  URLs() []string
}

type HTTPFetcher struct{}

type HTTPResponse struct {
  _Url string
  _Body string
}

func (r HTTPResponse) URLs() []string {
  rx, _ := regexp.Compile("<a.*href=[\"'](\\S+)[\"']")
  match := rx.FindAllStringSubmatch(r.Body(), -1)
  urls  := make([]string, 0)
  for _, v := range(match) {
    if(strings.HasPrefix(v[1], "http")) {
      urls = append(urls, v[1])
    }
  }
  return urls
}

func (r HTTPResponse) Url() string {
  return r._Url
}

func (r HTTPResponse) Body() string {
  return r._Body
}

func (f HTTPFetcher) Fetch(url string) (response Response, err error) {
  resp, err := http.Get(url)
  if(err != nil) {
    return HTTPResponse{}, err
  }
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  return HTTPResponse{url, string(body)}, nil
}

func CrawlHTTP(url string, depth int, fetcher Fetcher, fetched map[string]bool, ch chan Response) {
  if _, ok := fetched[url]; depth <= 0 || ok {
    return
  }
  fetched[url] = true
  response, err := fetcher.Fetch(url)
  if err != nil {
    fmt.Println(err)
    return
  }
  ch <- response
  for _, u := range response.URLs() {
    go CrawlHTTP(u, depth-1, fetcher, fetched, ch)
  }
}

func Crawl(url string, depth int, fetcher Fetcher) {
  var fetched = make(map[string]bool)
  ch := make(chan Response)
  go CrawlHTTP(url, depth - 1, fetcher, fetched, ch)
  for response := range(ch) {
    fmt.Printf("> %s, body size: %d, links: %d\n", response.Url(), len(response.Body()), len(response.URLs()))
  }
}

func main() {
  fetcher := HTTPFetcher{}
  Crawl(os.Args[1], 5, fetcher)
}
