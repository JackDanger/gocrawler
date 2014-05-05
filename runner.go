package main

import (
  "fmt"
  "code.google.com/p/go.net/html"
  "net/http"
)

func Start(uri string) {
  fmt.Println("starting")
  queue := make(chan string)

  go func() {queue <- uri}()

  crawl(queue)
}

func crawl(queue chan string) {
  for uri := range queue {
    fmt.Println("uri:", uri)
    enqueueLinks(uri, queue)
  }
}

func enqueueLinks(uri string, queue chan string) {
  fmt.Println("getting", uri)
  resp, err := http.Get(uri + "/")
  fmt.Println(resp, err)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  p := html.NewTokenizer(resp.Body)
  fmt.Println(p)
  for { 
      // token type
      tokenType := p.Next() 
      fmt.Println(tokenType)
      if tokenType == html.ErrorToken {
          return     
      }       
      token := p.Token()
      fmt.Println(token)
      if tokenType == html.StartTagToken {
        fmt.Println(token.Attr)
      }
  }
}
