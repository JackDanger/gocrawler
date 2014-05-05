package main

import (
  "fmt"
  "code.google.com/p/go.net/html"
  "net/http"
  "crypto/tls"
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
  transport := &http.Transport{ TLSClientConfig: &tls.Config{InsecureSkipVerify: true} }
  client := http.Client{Transport: transport}
  resp, err := client.Get(uri)
  fmt.Println(resp, err)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer resp.Body.Close()

  p := html.NewTokenizer(resp.Body)
  fmt.Println(p)
  for { 
    // token type
    tokenType := p.Next() 
    if tokenType == html.ErrorToken {
        return     
    }       
    token := p.Token()
    if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
      for _, attr := range token.Attr {
        if attr.Key == "href" {
          fmt.Println(attr.Val)
          go func() { queue <- attr.Val }()
        }
      }
    }
  }
}
