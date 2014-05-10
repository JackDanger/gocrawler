package main

import (
  "crypto/tls"
  "flag"
  "fmt"
  "github.com/jackdanger/collectlinks"
  "net/http"
  "net/url"
  "os"
)

func main() {
  flag.Parse()

  args := flag.Args()
  fmt.Println(args)
  if len(args) < 1 {
    fmt.Println("Please specify start page")
    os.Exit(1)
  }

  queue := make(chan string)

  go func() { queue <- args[0] }()

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

  for _, link := range(collectlinks.All(resp.Body)) {
    absolute := fixUrl(link, uri)
    if !isVisited(absolute) {
      go func() { queue <- absolute }()
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
}

var visited = make(map[string]bool)

func visit(uri string) {
  visited[uri] = true
  display(uri)
}
func isVisited(uri string) (bool) {
  return visited[uri]
}
