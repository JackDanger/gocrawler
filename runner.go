package main

import (
  "fmt"
  "code.google.com/p/go.net/html"
  "net/http"
  "crypto/tls"
  "net/url"
)

func Start(uri string) {
  fmt.Println("starting")
  queue := make(chan string)

  go func() {queue <- uri}()

  crawl(queue)
}

func crawl(queue chan string) {
  for uri := range queue {
    if uri != "" {
      enqueueLinks(uri, queue)
    }
  }
}

func enqueueLinks(uri string, queue chan string) {
  visit(uri)
  transport := &http.Transport{ TLSClientConfig: &tls.Config{InsecureSkipVerify: true} }
  client := http.Client{Transport: transport}
  resp, err := client.Get(uri)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  p := html.NewTokenizer(resp.Body)
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
          absolute := fixUrl(attr.Val, uri)
          if !isVisited(absolute) {
            go func() { queue <- absolute }()
          }
        }
      }
    }
  }
}

func fixUrl(href, base string) (string) {
  uri, err := url.Parse(href)
  if err != nil {
    return ""
  }
  baseUrl, err := url.Parse(base)
  if err != nil {
    return ""
  }
  uri = baseUrl.ResolveReference(uri)
  return uri.String()
  return ""
}

var visited = make(map[string]bool)

func visit(uri string) {
  visited[uri] = true
  display(uri)
}
func isVisited(uri string) (bool) {
  return visited[uri]
}

func display(uri string) {
  fmt.Print("\033[A\033[A")
  fmt.Println("visited:", len(visited))
  fmt.Println(uri)
}
