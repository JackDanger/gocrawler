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
                                     // receives and delivers strings. There's nothing
                                     // more you need to do to set it up – it's
                                     // all ready to have data fed into it.

  go func() {            // saying "go someFunction()" means
                         // "run someFunction() asynchronously"
    queue <- args[0]     // This means "put args[0] into the channel".
  }()

  for uri := range queue {     // 'range' is such an effective iterator keyword
                               // that if you ask for the range of a channel it'll
                               // do an efficient, continuous blocking read of
                               // all the channel contents.

    enqueue(uri, queue)  // we pass each URL we find off to be read & enqueued
  }
}

func enqueue(uri string, queue chan string) {
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

  links := collectlinks.All(resp.Body)

  for _, link := range links {
    go func() { queue <- link }() // We asynchronously enqueue what we've found
  }
}
