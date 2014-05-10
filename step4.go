package main

import (
  "crypto/tls"
  "flag"
  "fmt"
  "github.com/jackdanger/collectlinks"
  "net/http"
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

  queue := make(chan string)         // This gives you a new channel that
                                     // receives and delivers strings

  go func() {            // saying 'go someFunction' means "run someFunction asynchronously"
    queue <- args[0]     // This means "put args[0] into the channel".
  }()                    // You don't have to understand this, but 'go' takes
                         // a function call to execute.

                              // 'range' is such an effective iterator keyword
  for uri := range queue {    // that if you ask for the range of a channel it'll
                              // do a permanent blocking read of all the channel contents.
    enqueueLinks(uri, queue)  // we pass each URL we find off to be read & enqueued
  }
}
func enqueueLinks(uri string, queue chan string) {
  fmt.Println("fetching", uri)
  tlsConfig := &tls.Config{
    InsecureSkipVerify: true,
  }
  transport := &http.Transport{
    TLSClientConfig: tlsConfig,
  }
  client := http.Client{Transport: transport}
  resp, err := client.Get(uri)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  for _, link := range collectlinks.All(resp.Body) {  // Here I inline what used to be a 'links' variable
    go func() { queue <- link }() // We asynchronously enqueue what we've found
  }
}
