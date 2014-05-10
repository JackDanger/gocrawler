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

                                     // Putting a variable outside a function is Go's
var visited = make(map[string]bool)  // version of a global variable.
                                     // This is a map of string -> bool, so
                                     // visited["hi"] works but visited[6] doesn't.
                                     // Setting a value in a map is a simple as:
                                     //   visited["google.com"] = true
                                     // and reading is equally so:
                                     //   visited["google.com"] 
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
    enqueue(uri, queue)
  }
}

func enqueue(uri string, queue chan string) {
  fmt.Println("fetching", uri)
  visited[uri] = true                        // Record that we're going to visit this page
  transport := &http.Transport{
    TLSClientConfig: &tls.Config{
      InsecureSkipVerify: true,
    },
  }
  client := http.Client{Transport: transport}
  resp, err := client.Get(uri)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  links := collectlinks.All(resp.Body) 

  for _, link := range links {
    absolute := fixUrl(link, uri)
    if uri != "" {
      if !visited[absolute] {          // Don't enqueue a page twice!
        go func() { queue <- absolute }()
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
}
